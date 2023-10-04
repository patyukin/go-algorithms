package bufferpool

import (
	"bytes"
	"sync"
)

type Pool struct {
	pool sync.Pool
}

func NewPool() *Pool {
	return &Pool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		}}
}

func (b *Pool) GetBuffer() *bytes.Buffer {
	return b.pool.Get().(*bytes.Buffer)
}

func (b *Pool) PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	b.pool.Put(buf)
}
