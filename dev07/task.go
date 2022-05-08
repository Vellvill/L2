package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, ch2, ch3, ch4 := make(chan interface{}), make(chan interface{}), make(chan interface{}), make(chan interface{})
	channels := []chan interface{}{ch1, ch2, ch3, ch4}
	chdata := [][]interface{}{{1, 1, 1, 1}, {"2", 2.2, "2", "2"}, {"3, 3, 3", 3, 3, 3}, {-4, "4"}}
	for i, _ := range chdata {
		go workerWriter(channels[i], chdata[i])
	}
	for v := range Or(ch1, ch2, ch3, ch4) {
		fmt.Println(v)
	}

}

func Or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan struct{})
	for i, ch := range channels {
		go workerListen(ch, done, i)
	}
	for {
		select {
		case <-done:
			return merge(channels)
		}
	}
}

func workerWriter(ch chan<- interface{}, data []interface{}) {
	defer close(ch)
	for _, v := range data {
		ch <- v
	}
}

func workerListen(ch <-chan interface{}, done chan struct{}, i int) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				done <- struct{}{}
				fmt.Printf("closing %d ch\n", i)
				return
			}
		case <-done:
			return
		}
	}
}

func merge(chs []<-chan interface{}) <-chan interface{} {
	merged := make(chan interface{})
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan interface{}) {
				defer wg.Done()
				for data := range ch {
					merged <- data
				}
			}(ch)
		}
		wg.Wait()
		close(merged)
	}()
	return merged
}
