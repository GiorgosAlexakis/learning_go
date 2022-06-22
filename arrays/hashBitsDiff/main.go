package main

import (
	"crypto/sha256"
	"log"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDiff(a *[32]byte, b *[32]byte) (count int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < 8; j++ {
			if int(pc[a[i]>>(j*8)]) != int(pc[b[i]>>(j*8)]) {
				count++
			}
		}
	}
	return count
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	log.Println(countDiff(&c1, &c2))
}
