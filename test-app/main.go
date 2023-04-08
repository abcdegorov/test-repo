package main

import (
	"fmt"
	"time"
)

func main() {
	curTime := time.Now()
	fmt.Printf("Hello World!\nThe current time is %v.\n", curTime.Format("15:04:05"))
}
