package main

import (
	"go-spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.Default, 5, spinner.NoDuration)

	q := make(chan struct{})
	go func() {
		// work
		time.Sleep(time.Second * 5)
		// end
		q <- struct{}{}
	}()
	s.Start(q)
}
