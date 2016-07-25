//Dining Philosopher problem is philosophers sitting around the cirular table,and their works are thinking and eating.
//There is only limited no of spoons available to eat by tha all philosophers.
//Philosophers must have two spoon to eat.For this need to synchronise the process of eating and ensures the mutual exclusion for the shared resouce.
//And also need to ensure no deadlock results.

package diningPhilosopher

import (
	//"fmt"
	"time"
)

//Philosophers struct consists of array or Philosoper struct which represents number of philosophers sittirn arround circular table , slice of spoons represent the shared resource and spoonChan slice is for mutual exclusion and synchronization.
type Philosophers struct {
	philosopherNo []*Philosopher
	spoon         []int
	spoonChan     []chan string
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
	d.spoonChan = make([]chan string, n)

	for i := 0; i < n; i++ {
		ph := new(Philosopher)
		d.spoon = append(d.spoon, i)
		d.spoonChan[i] = make(chan string)
		ph.state = Thinking
		d.philosopherNo = append(d.philosopherNo, ph)
	}
	go func() {
		for i := 0; i < n; i++ {
			d.spoonChan[i] <- "free"
		}

	}()
	return d
}

//Think method made an individual phlosopher to think for a while.After thinking he proceed to eat.
func (d *Philosophers) Think(philosopher int) {
	//fmt.Println("Philosopher", philosopher, "is thinking")
	time.Sleep(1 * time.Second) //Thinking for 1 second
	d.philosopherNo[philosopher-1].state = Hungry
	d.Eat(philosopher) //Calling Eat method
}

//Eat method checks for the availability of shared resource(spoons).
//If the spoons are free then allowed the philosopher to eat.After performing eating put down the spoons on the table.
//Otherwise wait them unitl they are available.
func (d *Philosophers) Eat(philosopher int) {
	size := len(d.philosopherNo)
	leftspoon := philosopher - 1
	rightspoon := (philosopher + size) % size
	var availleftspoon, availrightspoon bool

	if d.philosopherNo[philosopher-1].state == Hungry {
		availleftspoon = d.CheckAvailability(leftspoon)
		availrightspoon = d.CheckAvailability(rightspoon)
	}

	if availleftspoon == true && availrightspoon == true {
		d.philosopherNo[philosopher-1].state = Eating
		//fmt.Println("Philosopher", philosopher, "is thinking")
		time.Sleep(1 * time.Second)
		//fmt.Println("***********Done with eating Philosopher :", philosopher, "*************")
		d.philosopherNo[philosopher-1].state = Thinking
		d.spoonChan[leftspoon] <- "free"  // After done wtih eating send the message "free" to channel that spoon is now availble.
		d.spoonChan[rightspoon] <- "free" // After done wtih eating send the message "free" to channel that spoon is now availble.
	}

}

//CheckAvailability method checks for the availability of spoon whether it is free or in use by some other phiolosopher.
func (d *Philosophers) CheckAvailability(spoon int) bool {
	msg := <-d.spoonChan[spoon] //msg: recieves whether it is free or not.
	if "free" == msg {          //If spoon i free returns true otherwise false.
		return true
	} else {
		return false
	}
}
