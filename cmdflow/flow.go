package main

import "time"

type Flow struct {
	in ConsoleParser
	// inpurt buf
	buf *Buffer

	// history 执行历史

}

func (f *Flow) Run() {
	bufCh := make(chan []byte, 128)
	stopReadBufCh := make(chan struct{})
	go f.readBuffer(bufCh, stopReadBufCh)
}

func (f *Flow) readBuffer(bufCh chan []byte, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			debug.Log("stop reading buffer")
			return
		default:
			if b, err := p.in.Read(); err == nil && !(len(b) == 1 && b[0] == 0) {
				bufCh <- b
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}
