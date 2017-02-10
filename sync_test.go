package syncbench

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkNoContention(b *testing.B) {
	var counter int64

	for i := 0; i < b.N; i++ {
		counter++
	}
}

func BenchmarkMutexNoContention(b *testing.B) {
	var mutex sync.Mutex
	var counter int64

	for i := 0; i < b.N; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func BenchmarkMutexContention(b *testing.B) {
	var mutex sync.Mutex
	var counter int64
	var done bool

	go func() {
		for {
			mutex.Lock()
			if done {
				mutex.Unlock()
				return
			}
			mutex.Unlock()
		}
	}()

	for i := 0; i < b.N; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}

	mutex.Lock()
	done = true
	mutex.Unlock()
}

func BenchmarkAtomicNoContention(b *testing.B) {
	var counter int64

	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func BenchmarkAtomicContention(b *testing.B) {
	var counter int64
	var done int32

	go func() {
		for {
			if atomic.LoadInt32(&done) > 0 {
				return
			}
			atomic.AddInt64(&counter, 1)
		}
	}()

	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}

	atomic.StoreInt32(&done, 1)
}

func BenchmarkChannelSelect(b *testing.B) {
	deltaChan := make(chan int64)
	resultChan := make(chan int64)

	go func() {
		var counter int64
		for {
			select {
			case delta := <-deltaChan:
				counter += delta
			case resultChan <- counter:
				return
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		deltaChan <- 1
	}

	<-resultChan
}

func BenchmarkChannelRange(b *testing.B) {
	deltaChan := make(chan int64)
	resultChan := make(chan int64)

	go func() {
		var counter int64
		for delta := range deltaChan {
			counter += delta
		}
		resultChan <- counter
	}()

	for i := 0; i < b.N; i++ {
		deltaChan <- 1
	}

	close(deltaChan)
	<-resultChan
}
