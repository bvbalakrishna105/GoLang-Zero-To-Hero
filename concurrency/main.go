package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//exampleChannel()
	//exampleFanInOut()
	//exampleSelect()
	//exampleBufferedChannel()
	//exampleMutex()
	exampleRWMutex()

}

var (
	value   int
	rwMutex sync.RWMutex
)

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Printf("Read : %v\n", value)
}

func write(newValue int) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Printf("Write : %v \n", newValue)
	value = newValue
}

func exampleRWMutex() {

	for i := 0; i < 3; i++ {
		go read()
	}

	for i := 0; i < 2; i++ {
		go write(i)
	}
	time.Sleep(2 * time.Second)
}

var (
	counter int
	mutex   sync.Mutex
)

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	fmt.Printf("Counter: %v \n", counter)
}
func exampleMutex() {
	for i := 1; i < 10; i++ {
		go increment()
	}
	time.Sleep(time.Second)
	fmt.Printf("Final counter %v: \n", counter)
}

func exampleBufferedChannel() {
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("Sending: %v \n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("Receiving : %v \n", num)
	}
}

func printNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		//fmt.Printf("Value : %v \n", i)
		ch <- i
		time.Sleep(time.Microsecond * 500)
	}
	close(ch)
}

func exampleChannel() {

	ch := make(chan int)
	go printNumbers(ch)

	for num := range ch {
		fmt.Println(num)
	}
	//time.Sleep(time.Second * 3)
	fmt.Println("Main function is done")
}

func exampleFanInOut() {
	numberOfJobs := 5
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)
	numberOfWorkers := 3
	for i := 1; i <= numberOfWorkers; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg := sync.WaitGroup{}
	wg.Add(numberOfJobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Result %d, \n", result)
		wg.Done()
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Woker %d started job %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished %d \n", id, job)
		results <- job * 2
	}
}

func exampleSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "Hello from Channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!!!")
	}
}
