package counter

import (
	"bufio"
	"bytes"
)

type Bytecounter int

func (c *Bytecounter) Write(p []byte) (int, error) {
	*c += Bytecounter(len(p))
	return len(p), nil
}

type Linecounter int

func (c *Linecounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	var count int
	for s.Scan() {
		count++
	}
	*c = Linecounter(count)
	return count, nil
}
