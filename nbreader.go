package main

import (
	"bufio"
	"os"
	"time"
)

type NonBlockingReader struct {
	err      chan error
	data     chan string
	ctrl     chan bool
	prompt   string
	sentinal byte
	errFunc  func(error)
}

func (r *NonBlockingReader) New() {
	r.err = make(chan error)
	r.ctrl = make(chan bool)
	r.data = make(chan string)
	r.sentinal = '\n'
	r.errFunc = func(err error) {}

	go r.readLine()
}

func (r *NonBlockingReader) Close() {
	// This will cause a deadlock - there is no way to close the routine
	// given that the readline blocks.
	//r.ctrl <- true
}

func (r *NonBlockingReader) BlockingRead() (string, error) {

	select {
	case cmd := <-r.data:
		return cmd, nil
	case err := <-r.err:
		return "", err
	}
}

func (r *NonBlockingReader) NonBlockingRead() (string, error) {
	select {
	case cmd := <-r.data:
		return cmd, nil // cmd must be non-empty
	case err := <-r.err:
		return "", err // maybe EOF or something else
	case <-time.After(50 * time.Millisecond):
		return "", ErrNoData
	}
}

func (r *NonBlockingReader) readLine() {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case ctrl := <-r.ctrl:
			if ctrl {
				return
			}

		default:
			s, err := reader.ReadString(r.sentinal)
			if err != nil {
				go r.errFunc(err)
				r.err <- err
			} else {
				r.data <- s
			}
		}
	}
}
