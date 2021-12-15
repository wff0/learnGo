package testDemo

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch1 := make(chan int, 1)
	fmt.Println(len(ch1))
	ch1 <- 1
	fmt.Println(len(ch1))
	<-ch1
	fmt.Println(len(ch1))
	close(ch1)
}

func TestChannel(t *testing.T) {
	//c := make(chan int)
	c := make(chan int, 10)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ready ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

func recv(c <-chan int) {
	for n := range c {
		fmt.Println("received ", n)
	}
}
