package main

import (
	"fmt"
	"sync"
)

func odd(ch chan int, wg *sync.WaitGroup, flag chan struct{}) {
	defer wg.Done()

	for {

		select {
		case no, ok := <-ch:

			if !ok {
				flag <- struct{}{}
				return

			}
			if no%2 != 0 {
				fmt.Println("odd :: ", no)

			} else {
				ch <- no
			}

		}

	}

}

func even(ch chan int, wg *sync.WaitGroup, flag chan struct{}) {
	defer wg.Done()
	for {
		select {
		case no, ok := <-ch:

			if !ok {
				flag <- struct{}{}
				return
			}
			if no%2 == 0 {
				fmt.Println("Even :: ", no)
			} else {
				ch <- no
			}

		}
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	ch := make(chan int)

	flag := make(chan struct{})
	go odd(ch, &wg, flag)
	go even(ch, &wg, flag)
	// go Number(ch, &wg)

	for i := 1; i <= 10; i++ {
		ch <- i
		// time.Sleep(time.Second * 1)
	}

	// go func() {
	close(ch)
	wg.Wait()
	<-flag
	<-flag

	// }()

	// 1 2
	// 2 1

	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// a1 := arr[2:5] // 3 ,4 , 2,8
	// a1 = append(a1, 11, 12, 13, 14, 15)

	// fmt.Println(a1, len(a1), cap(a1))

	// fmt.Println(a1, len(a1), cap(a1))

}
