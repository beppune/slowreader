package slowreader

import (
	"io"
	"time"
)

type SlowReader struct {
	r   io.Reader
	bps int
}

func NewSlowReader(bytesPerSecond int, r io.Reader) SlowReader {
	slow := SlowReader{r, bytesPerSecond}
	return slow
}

func (r *SlowReader) Read(b []byte) (int, error) {
	buffer := b[0:r.bps]
	time.Sleep(1 * time.Second)
	s, e := r.r.Read(buffer)
	return s, e
}
