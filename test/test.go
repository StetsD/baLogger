package main

import (
	"baLogger"
	"time"
)

func something(a, b int) int {
	time.Sleep(2 * time.Second)
	return a + b
}

func main() {
	baLogger.Log(func() { something(2, 2) }).Flag("Mazafaka").Bench().Call()
}
