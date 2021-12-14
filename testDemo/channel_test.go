package testDemo

import (
	"fmt"
	"testing"
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
