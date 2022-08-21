package locks

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func BenchMarkMutex() {
	producer := func(wg *sync.WaitGroup, lock sync.Locker) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			lock.Lock()
			lock.Unlock()
			time.Sleep(1)
		}
	}

	consumer := func(wg *sync.WaitGroup, lock sync.Locker) {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		startTime := time.Now()
		var wg sync.WaitGroup
		wg.Add(count + 1)
		go producer(&wg, mutex)
		for i := 0; i < count; i++ {
			go consumer(&wg, rwMutex)
		}
		wg.Wait()

		return time.Since(startTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var mutex sync.RWMutex
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%v\t%v\t%v\n", count, test(count, &mutex, &mutex), test(count, &mutex, mutex.RLocker()))
	}

}
