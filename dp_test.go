package diningPhilosopher

import (
	"container/list"
	DP "github.com/Sreevani871/diningPhilosopher"
	"testing"
	"time"
)

func TestDiningPhilosopher(t *testing.T) {
	var size = 6
	D := DP.Initialize(size)
	l := list.New()

	for i := 1; i <= size; i++ { //Adding philosophers to the doubly linked list.
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		go D.Think(v)
	}
	time.Sleep(time.Second * 40)
	for _, v := range D.Count {
		if v != 5 {
			t.Error("Expected", 5, "Got", v)
		}
	}
}
