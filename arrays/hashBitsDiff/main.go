package main

import (
	"crypto/sha256"
	"log"
)

func countDiff(a *[32]byte, b *[32]byte) (count int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < 8; j++ {
			if (int(a[i]>>(j)) & 1) != (int(b[i]>>(j)) & 1) {
				count++
			}
		}
		log.Println(count)
		log.Printf("a[i]: %08b\n", a[i])
		log.Printf("b[i]: %08b\n\n", b[i])

	}
	return count
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	log.Println(countDiff(&c1, &c2))
}
