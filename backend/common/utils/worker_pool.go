package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

var (
	ErrInvalidMaxRoutinesCount = errors.New("Passed argument max goroutines count value is invalid.")
)

type workItem = func()

type WorkerPool struct {
	maxRoutinesCount  int
	stopWaiters		  sync.WaitGroup

	stopChannel		  chan struct{}
	workItems         chan workItem
}

func NewWorkerPool(maxRoutinesCount int) (*WorkerPool, error) {

	if maxRoutinesCount == 0 {
		return nil, ErrInvalidMaxRoutinesCount
	}

	wp := &WorkerPool{
		maxRoutinesCount: maxRoutinesCount,
		stopChannel: make(chan struct{}),
		workItems: make(chan workItem),
	}

	wp.stopWaiters.Add(maxRoutinesCount)

	for num := 0; num < maxRoutinesCount; num++ {
		go wp.workFunc()
	}

	return wp, nil
}

func (wp *WorkerPool) QueueWorkItem(item workItem) {

	wp.workItems <- item
}

func (wp *WorkerPool) WaitForCompletion() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic was called in WaitForCompletion func, ", r)
		}
	}()

	wp.stopChannel <- struct{}{}
	close(wp.stopChannel)

	wp.stopWaiters.Wait()
}

func (wp *WorkerPool) workFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic was called in worker func", r)
		}
	}()

	for {
		select {
		case item := <-wp.workItems:
			wp.doWork(item)
			break
		case <-wp.stopChannel:
			wp.stopWaiters.Done()
			return
		}
	}
}

func (wp *WorkerPool) doWork(item workItem) {
	item()
}


