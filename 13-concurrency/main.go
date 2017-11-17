package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	//goroutines1()
	//goroutines2()
	//channels1()
	//doSomething()
	//channels2()
	locks1()
}

func goroutines1() {
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("ping")
		}
	}()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("pong")
	}
}

func goroutines2() {
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("clear cache")
		}
	}()
	for i := 0; i < 20; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("incoming request")
	}
}

func doSomething() {
	c := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		c <- true
	}()

	fmt.Println(<-c)
}

func channels1() {
	c := make(chan string)
	go func() {
		for s := range c {
			fmt.Println("received", s)
		}
	}()
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		fmt.Println("sending", s)
		c <- s
	}
}

func channels2() {
	c := make(chan string, 5)
	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		fmt.Println("sending", s)
		c <- s
	}
	go func() {
		for s := range c {
			fmt.Println("received", s)
		}
	}()
	time.Sleep(time.Second)
}

func locks1() {
	count := 0
	for g := 0; g < 10; g++ {
		go func() {go func() {
			var mu sync.Mutex
			for i := 0; i < 10; i++ {
				mu.Lock()
				newCount := count
				time.Sleep(time.Millisecond)
				newCount++
				count = newCount
				mu.Unlock()
			}
		}()}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
