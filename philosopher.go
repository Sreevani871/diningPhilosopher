//Dining Philosopher problem is philosophers sitting around the cirular table,and their works are thinking and eating.
//There is only limited no of spoons available to eat by tha all philosophers.
//Philosophers must have two spoon to eat.For this need to synchronise the process of eating and ensures the mutual exclusion for the shared resouce.
//And also need to ensure no deadlock results.

package diningPhilosopher

import (
	"fmt"
	"sync"
	"time"
)

//Philosophers struct consists of array or Philosoper struct which represents number of philosophers sittirn arround circular table , slice of spoons represent the shared resource and spoonChan slice is for mutual exclusion and synchronization.
type Philosophers struct {
	philosopherNo []*Philosopher
	spoon         []int
	SpoonChan     []chan string
	Count         []int
	sync.Mutex
}

//Philosopher struct represent structure of individual philosopher state.
type Philosopher struct {
	state int
}

const (
	Thinking = 0
	Hungry   = 1
	Eating   = 2
)

//Initilaize function initializes the struct variables with the specified number of philosophers sitting around a table and returns a struct.
func Initialize(n int) *Philosophers {

	d := new(Philosophers)
	d.SpoonChan = make([]chan string, n)
	for i := 0; i < n; i++ {
		ph := new(Philosopher)
		d.spoon = append(d.spoon, i)
		d.SpoonChan[i] = make(chan string, 1)
		ph.state = Thinking
		d.philosopherNo = append(d.philosopherNo, ph)
		d.Count = append(d.Count, 0)
	}
	go func() {
		for i := 0; i < n; i++ {
			d.SpoonChan[i] <- "free"
		}

	}()

	return d
}

//Think method made an individual phlosopher to think for a while.After thinking he proceed to eat.
func (d *Philosophers) Think(philosopher int) {

	fmt.Println("Philosopher", philosopher, "is thinking")
	time.Sleep(1 * time.Second) //Thinking for 1 second

	d.philosopherNo[philosopher-1].state = Hungry

	d.Test(philosopher) //Calling Eat method

}

//Eat method checks for the availability of shared resource(spoons).
//If the spoons are free then allowed the philosopher to eat.After performing eating put down the spoons on the table.
//Otherwise wait them unitl they are available.
func (d *Philosophers) Test(philosopher int) {

	//fmt.Println("I am into eat method")
	//fmt.Println("phichan", msg1)
	size := len(d.philosopherNo)
	leftspoon := philosopher - 1
	rightspoon := (philosopher + size) % size
	var availleftspoon, availrightspoon bool
	if d.philosopherNo[philosopher-1].state == Hungry {
		d.Lock()
		//fmt.Println("left", availleftspoon, philosopher)
		availleftspoon = d.CheckAvailability(leftspoon)
		//fmt.Println("right", availrightspoon, philosopher)
		availrightspoon = d.CheckAvailability(rightspoon)
		//fmt.Println("left,right", availleftspoon, availrightspoon)
	}
	d.Eat(availleftspoon, availrightspoon, philosopher, leftspoon, rightspoon)

}

func (d *Philosophers) Eat(availleftspoon, availrightspoon bool, philosopher int, leftspoon, rightspoon int) {
	//fmt.Println("left,right", availleftspoon, availrightspoon, philosopher)
	if availleftspoon == true && availrightspoon == true {
		d.philosopherNo[philosopher-1].state = Eating
		fmt.Println("Philosopher", philosopher, "is eating")
		time.Sleep(1 * time.Second)
		fmt.Println("***********Done with eating Philosopher :", philosopher, "************* for the", d.Count[philosopher-1]+1, "time")
		d.Count[philosopher-1]++

		d.Unlock()
		d.philosopherNo[philosopher-1].state = Thinking

		d.SpoonChan[leftspoon] <- "free" // After done wtih eating send the message "free" to channel that spoon is now availble.
		d.SpoonChan[rightspoon] <- "free"
		if d.Count[philosopher-1] != 5 {
			d.Think(philosopher)
		}
	}

}

//CheckAvailability method checks for the availability of spoon whether it is free or in use by some other phiolosopher.
func (d *Philosophers) CheckAvailability(spoon int) bool {
	//fmt.Println("I am in checking", spoon)
	msg := <-d.SpoonChan[spoon] //msg: recieves whether it is free or not.
	if "free" == msg {          //If spoon i free returns true otherwise false.
		//fmt.Println("free")
		return true
	} else {
		//fmt.Println("not free")
		return false
	}
}
