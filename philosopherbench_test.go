package diningPhilosopher

import (
	"container/list"
	DP "github.com/Sreevani871/diningPhilosopher"
	"testing"
)

func BenchmarkInitialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DP.Initialize(5)
	}

}

func BenchmarkThink(b *testing.B) {
	D := DP.Initialize(5)
	l := list.New()
	for i := 1; i <= 5; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		go D.Think(1)
	}
}

func BenchmarkCheckAvailability(b *testing.B) {
	D := DP.Initialize(5)
	l := list.New()
	for i := 1; i <= 5; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		D.CheckAvailability(2)
	}

}

func BenchmarkEat(b *testing.B) {
	D := DP.Initialize(5)
	l := list.New()
	for i := 1; i <= 5; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		D.Eat(2)
	}
}
