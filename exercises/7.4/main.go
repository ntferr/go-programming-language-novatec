package main

type Reader struct {
	s        string
	i        int64
	prevRune int
}

func (r *Reader) Read([]byte) (n int, err error) {

	return 0, nil
}

func NewReader(s string) *Reader { return &Reader{s, 0, 1} }

func main() {
	html.Parse()
}
