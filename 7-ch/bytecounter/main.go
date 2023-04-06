package main

import (
	"bytecounter/bytec"
	"bytecounter/scanwords"
	"fmt"
)

func main() {
	var c bytec.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var c2 scanwords.Counter
	c2.Write([]byte("hello\nmy friend"))
	fmt.Println(c2)
}
