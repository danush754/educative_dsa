package educativedsa

import (
	"fmt"
	"sync"
)

func SendData(ch *chan interface{}) {
	*ch <- "Dhanush"
	*ch <- 123
	*ch <- 33.987
	*ch <- nil

	close(*ch)
}

func GetData(ch *chan interface{}, wg *sync.WaitGroup) {

	// fmt.Println("chaah", ch)
	defer wg.Done()
	for input := range *ch {

		fmt.Println("inputt", input)
	}

	// close(ch)
}

func Semaphore() {

	maxGoRuntines := 5

	semaphore := make(chan struct{}, maxGoRuntines)

	fmt.Println("mamaam", semaphore)
}

func Sum(x, y int, c chan int) {

	for i := 0; i < y; i++ {
		c <- i * x
	}

}
