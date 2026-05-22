package parallel

import "sync"

func Execute(
	workers int,
	fn func(),
) {

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
		}()
	}

	wg.Wait()
}
