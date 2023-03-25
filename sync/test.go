package sync

import "fmt"

func test() {

	ch1 := make(chan int, 3)

	//mutex:=&sync.Mutex{}
	f := func(i int) {
		//mutex.Lock()
		fmt.Printf("the %v th goroutine", i)
		ch1 <- i
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		go f(i)
	}

	for i := range ch1 {
		fmt.Printf("the %v th goroutine", i)
	}

}
