package main

import (
	"context"
	"fmt"
)

type message struct {
	responseChan chan<- int
	parameter    string
	ctx          context.Context
}

func ProcessMessages(work <-chan message) {
	for job := range work {
		select {
		case <-job.ctx.Done():
			continue
		default:
		}
		hardToCalculate := len(job.parameter)
		select {
		case <-job.ctx.Done():
		case job.responseChan <- hardToCalculate:
		}
	}
}

func newRequest(ctx context.Context, input string, q chan<- message) {
	r := make(chan int)
	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q could see")
		return
	case q <- message{
		responseChan: r,
		parameter:    input,
		ctx:          ctx,
	}:
		{
			select {
			case out := <-r:
				fmt.Printf("The len of %s is %d\n", input, out)
			case <-ctx.Done():
				fmt.Println("Context ended before q could process message")
			}
		}
	}
}

func main() {
	q := make(chan message)
	go ProcessMessages(q)
	ctx := context.Background()
	newRequest(ctx, "hi", q)
	newRequest(ctx, "hello", q)
	close(q)
}
