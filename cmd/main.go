//Напишите программу, которая делает следующее:
//одна горутина по порядку отсылает числа от 1 до 100 в канал,
//вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).
package main

import (
	"fmt"
	"sync"
)

// Функция sends по порядку отсылает числа от 1 до 100 в канал.
func sends(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}

	wg.Done()
}

// функция accPrints принимает значения в правильном порядке
// и печатает на экран (в консоль).
func accPrints(p chan int) {
	for v := range p {
		fmt.Print(v, " ")
	}
}

func main() {
	// Инициализация канал ch1
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go sends(ch, &wg)
	go accPrints(ch)

	wg.Wait()
}
