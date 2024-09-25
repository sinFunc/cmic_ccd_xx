package main

import (
	"cmic_ccd_xx/internal/app"
	"fmt"
	"time"
)

func testChannel() {
	const size = 32
	c := make(chan int, size)
	go func() {
		for {
			select {
			case r := <-c:
				fmt.Println("rcv data=", r)
				if r > 1000 {
					return
				}
			default:
				fmt.Println("lost data this time.")

			}
		}
	}()

	for i := 0; i < 1002; i++ {
		c <- i
	}

	time.Sleep(time.Second)

}

func main() {
	//testChannel()
	app.Start()

}
