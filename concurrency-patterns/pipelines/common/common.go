package common

func RepeatStage(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	repeatedValues := make(chan interface{})
	go func() {
		defer close(repeatedValues)

		for {
			for _, i := range values {
				select {
				case <-done:
					return
				case repeatedValues <- i:
				}
			}
		}
	}()
	return repeatedValues
}

func TakeStage(done <-chan interface{}, valueStream <-chan interface{}, n int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func RepeatFunctionStage(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	functionValueStream := make(chan interface{})
	go func() {
		defer close(functionValueStream)
		for {
			select {
			case <-done:
				return
			case functionValueStream <- fn():
			}

		}
	}()
	return functionValueStream
}
