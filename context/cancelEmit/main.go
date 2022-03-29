package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operationX(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("fail on purpose")
}

func operationY(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operationY")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := operationX(ctx)
		if err != nil {
			cancel()
		}
	}()
	operationY(ctx)
}
