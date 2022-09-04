package orchannel

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

func SimulateOrChannel() {
	sleep := func(n time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(n)
		}()
		return c
	}

	fmt.Println("started at", time.Now())
	<-or(
		sleep(15*time.Second),
		sleep(10*time.Second),
		sleep(50*time.Second),
		sleep(25*time.Second),
		sleep(100*time.Second),
	)
	fmt.Println("finished at", time.Now())
}
