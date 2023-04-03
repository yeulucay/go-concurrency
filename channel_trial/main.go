package main

import (
	"fmt"
	"time"
)

type Worker struct {
	c1   chan string
	quit chan bool
}

func (w *Worker) finish() {
	w.quit <- true
}

func (w *Worker) listen() {
	i := 0

	for {
		fmt.Printf("Loop %d\n", i)

		select {
		case abc := <-w.c1:
			fmt.Printf("ABC: %s\n", abc)
		case q := <-w.quit:
			if q {
				close(w.c1)
				close(w.quit)
				return
			}
		}

		i++
	}
}

func main() {

	w := Worker{
		c1:   make(chan string),
		quit: make(chan bool),
	}

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			w.c1 <- fmt.Sprintf("abc %d", i)
		}
		w.finish()
	}()

	w.listen()

	fmt.Println("process finished")
}
