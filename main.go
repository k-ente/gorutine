package main

import "fmt"

func rutine1(ch chan<- int) {
	for i := range 15 {
		fmt.Println("i:", i)
		i++
		if i%3 == 0 {
			ch <- i
		}
	}
}

func rutine2(ch <-chan int) {
	for i := range 15 {
		i++
		k := <-ch
		fmt.Println("k:", k)
	}

}

func main() {
	fmt.Print("start")
	t := make(chan int)
	go rutine1(t)
	rutine2(t)
	fmt.Print("end")

}
