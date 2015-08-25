package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Hello Delve")
}


