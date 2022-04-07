package main

import (
	"fmt"
)

const (
	BEIJING = iota
	SHANGHAI
	NANJING
)

func main() {
	fmt.Println("BEIJING=", BEIJING, "SHANGHAI=", SHANGHAI, "NANJING=", NANJING)
}
