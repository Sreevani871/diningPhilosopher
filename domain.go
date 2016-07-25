
package main

import (
	"container/list"
	"fmt"
	DP "github.com/Sreevani871/diningPhilosopher"
	"time"
)

func main() {
	var size int
	fmt.Println("Enter the No.Of Philosophers")
	fmt.Scanf("%d", &size)
	D := DP.Initialize(size)
	l := list.New()

	for i := 1; i <= size; i++ { //Adding philosophers to the doubly linked list.
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		go D.Think(v) //And calling Think method for each philosopher.

	}

	time.Sleep(time.Second * 100) //Wait until all goroutines execution completes.

}
