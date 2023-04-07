package main

import (
	"fmt"

	"7.1/counter"
)

func main() {
	fmt.Println("bytecounter")
	var cb counter.Bytecounter
	cb.Write([]byte("hello"))
	fmt.Println(cb)

	cb = 0
	var name = "Dolly"
	fmt.Fprintf(&cb, "hello, %s", name)
	fmt.Println(cb)

	fmt.Println("\nlinecounter")
	var lb counter.Linecounter
	lb.Write([]byte("hello\ndev"))
	fmt.Println(lb)

	lb = 0
	var place = "Florida"
	fmt.Fprintf(&lb, "stranger: hello, where you live?\nyou:%s\nstranger:that's cool!", place)
	fmt.Println(lb)
}
