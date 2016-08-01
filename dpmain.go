package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

const (
	Thinking = 1
	Eating   = 2
)

var wg sync.WaitGroup

type Philosopher struct {
	spoon     int
	SpoonChan chan bool
	Count     int
	States    int
}

type DPhilosophers struct {
	Philosophers []*Philosopher
	Size         int
}

var D *DPhilosophers

func Initialize(n int) *DPhilosophers {
	DP := new(DPhilosophers)
	DP.Size = n

	for i := 0; i < n; i++ {
		d := new(Philosopher)
		d.spoon = i
		d.SpoonChan = make(chan bool, 1)
		d.States = Thinking
		d.Count = 0
		d.SpoonChan <- true
		DP.Philosophers = append(DP.Philosophers, d)

	}

	D = DP
	return DP
}

func (DP *DPhilosophers) Run(PhiNo int) {
	wg.Add(1)
	DP.Philosophers[PhiNo-1].Think(PhiNo)
	ls, rs := DP.Philosophers[PhiNo-1].Spoons(PhiNo)

	if ls == true && rs == true {
		DP.Philosophers[PhiNo-1].Eat(PhiNo)
	} else {
		DP.Philosophers[PhiNo-1].Check(PhiNo, ls, rs)
	}

	if DP.Philosophers[PhiNo-1].Count != 5 {
		DP.Run(PhiNo)
	}
	wg.Done()

}

func (DP *Philosopher) Think(PhiNo int) {
	time.Sleep(1 * time.Second)

}

func (DP *Philosopher) Spoons(PhiNo int) (bool, bool) {

	rsp := D.Philosophers[PhiNo-1].spoon
	lsp := D.Philosophers[(PhiNo+D.Size)%D.Size].spoon
	var ls, rs bool

	select {
	case <-D.Philosophers[lsp].SpoonChan:
		ls = true
	default:
		ls = false
	}

	select {
	case <-D.Philosophers[rsp].SpoonChan:
		rs = true
	default:
		rs = false

	}

	return ls, rs
}

func (DP *Philosopher) Eat(PhiNo int) {

	DP.States = Eating
	fmt.Println("**********Eating PhilosopherNO:", PhiNo, "for the", DP.Count+1, "time")
	time.Sleep(time.Second * 1)
	DP.Count++
	D.Philosophers[PhiNo-1].SpoonChan <- true
	D.Philosophers[(PhiNo+D.Size)%D.Size].SpoonChan <- true

}

func (DP *Philosopher) Check(PhiNo int, ls, rs bool) {

	rsp := D.Philosophers[PhiNo-1].spoon
	lsp := D.Philosophers[(PhiNo+D.Size)%D.Size].spoon

	if ls == false && rs == true {
		D.Philosophers[rsp].SpoonChan <- true
	}

	if rs == false && ls == true {
		D.Philosophers[lsp].SpoonChan <- true

	}
	DP.Think(PhiNo)

}

func main() {

	var no int
	fmt.Println("Enter the number of philosophers")
	fmt.Scanf("%d", &no)

	l := list.New()

	for i := 0; i < no; i++ {
		l.PushBack(i + 1)
	}

	DP := Initialize(no)
	for e := l.Front(); e != nil; e = e.Next() {
		value := e.Value.(int)
		go DP.Run(value)

	}
	wg.Wait()
}
