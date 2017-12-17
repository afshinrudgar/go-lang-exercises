package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader    io.Reader
	readCount int64
	nops      int
	mutex     sync.Mutex
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.reader.Read(p)
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	rc.readCount += int64(n)
	rc.nops++
	return
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	n = rc.readCount
	nops = rc.nops
	return
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

type writeCounter struct {
	writer     io.Writer
	writeCount int64
	nops       int
	mutex      sync.Mutex
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.writer.Write(p)
	wc.mutex.Lock()
	defer wc.mutex.Unlock()
	wc.writeCount += int64(n)
	wc.nops++
	return
}
func (wc *writeCounter) WriteCount() (n int64, nops int) {
	n = wc.writeCount
	nops = wc.nops
	return
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

type readWriteCounter struct {
	*readCounter
	*writeCounter
}

func NewReadWriteCounter(readWriter io.ReadWriter) *readWriteCounter {
	return &readWriteCounter{
		readCounter:  &readCounter{reader: readWriter},
		writeCounter: &writeCounter{writer: readWriter},
	}
}
