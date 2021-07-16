package main

import "fmt"

func PublishTask(t int) bool {
	fmt.Printf("task {%d} published\n", t)
	return t%3 == 0
}

func PublishStatus(b bool) bool {
	fmt.Printf("status {%v} published\n", b)
	return true
}

func Publish(statuses chan bool, tasks, quit chan int) {
	for {
		select {
		case <-quit:
			return
		case t := <-tasks:
			statuses <- PublishTask(t)
		case s := <-statuses:
			PublishStatus(s)
		}
	}
}

func main() {
	ctasks := make(chan int)
	statuses := make(chan bool)
	done := make(chan int)
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8}

	go func() {
		for _, t := range tasks {
			ctasks <- t
		}
		done <- 0
	}()
	Publish(statuses, ctasks, done)
}
