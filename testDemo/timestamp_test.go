package testDemo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(T *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}
