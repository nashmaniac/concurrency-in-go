package selectstatement

import (
	"fmt"
	"sync"
	"time"
)

func BasicChannelTest() {

	type element struct {
		channel int
		data    int
	}
	var wg sync.WaitGroup
	var c1, c2 chan element
	c1 = make(chan element)
	c2 = make(chan element)

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(c1)
		for i := 0; i < 10; i++ {
			c1 <- element{channel: 1, data: i + 1}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c2)
		for i := 0; i < 15; i++ {
			c2 <- element{channel: 2, data: i + 1}
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case c1data, ok := <-c1:
			{
				if ok {
					fmt.Printf("%+v %s\n", c1data, time.Now())
				}

			}
		case c2data, ok := <-c2:
			{
				if ok {
					fmt.Printf("%+v %s\n", c2data, time.Now())
				}

			}
		case <-time.After(5 * time.Second):
			{
				fmt.Println("5 seconds passed")
			}
		}
	}

	wg.Wait()

}
