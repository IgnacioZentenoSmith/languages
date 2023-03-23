package learn

import "fmt"

type Dog struct {
  name string
  age int
  alive bool
}

// adding funcs to struct

// method
func (thisDog Dog) String() string {
	return fmt.Sprintf("My name is: %s and my age is: %d", thisDog.name, thisDog.age)
}

func (thisDog Dog) getAge() int {
	return thisDog.age
}
func (thisDog Dog) getName() string {
	return thisDog.name
}
func (thisDog Dog) isAlive() bool {
	return thisDog.alive
}

// function
func newDog (name string) Dog {
	newborn := Dog{name: name, age: 0, alive: true}
	return newborn
}

func (thisDog *Dog) Birthday() {
	thisDog.age = thisDog.age + 1
}

func (thisDog *Dog) Death() {
	thisDog.alive = false
}
