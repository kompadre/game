package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(99))
	done := make(chan bool)
	var i = 0
	for ; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("%v\n", i)
				select {
				case <-done:
					fmt.Printf("One worker is finishing their job\n")
					return
				default:
				}
				time.Sleep(time.Second / 2)
			}
		}(i)
	}
	for ; i < 200; i++ {
		time.Sleep(time.Second / 2)
		if r.Intn(10) > 5 {
			close(done)
			time.Sleep(time.Second * 30)
			break
		}
	}
}
