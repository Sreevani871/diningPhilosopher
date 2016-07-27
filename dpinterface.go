
package diningPhilosopher

type DiningPhilosopher interface{
	Think(int)
	Test(int)
	CheckAvailability(int)
	Eat(int)

}
