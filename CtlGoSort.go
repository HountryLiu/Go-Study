package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	dogCh := make(chan struct{})
	fishCh := make(chan struct{})
	catCh := make(chan struct{})

	wg := sync.WaitGroup{}

	dog(dogCh, catCh, &wg)
	fish(fishCh, dogCh, &wg)
	cat(catCh, fishCh, &wg)

	wg.Wait()
}

func dog(dogCh, catCh chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			fmt.Println("dog")
			dogCh <- struct{}{}
			<-catCh
		}
		wg.Done()
	}()
}
func fish(fishCh, dogCh chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-dogCh
			fmt.Println("fish")
			fishCh <- struct{}{}
		}
		wg.Done()
	}()
}
func cat(catCh, fishCh chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-fishCh
			fmt.Println("cat")
			time.Sleep(time.Second)
			catCh <- struct{}{}
		}
		wg.Done()
	}()
}
