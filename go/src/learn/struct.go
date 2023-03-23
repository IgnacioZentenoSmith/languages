package learn

import "fmt"

/* To declare a structure in Go, use the type and struct keywords:
	type struct_name struct {
		member1 datatype;
		member2 datatype;
		member3 datatype;
		...
	}
*/

// minuscule -> private, Mayuscule -> public
type Person struct {
  name string
  age int
  job string
  salary int
}

// Access Struct Members
func Struct_ex() {
	var pers1 Person
  var pers2 Person

	// Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

	// Print Pers1 info by calling a function
  printPerson(pers1)
  // Print Pers2 info by calling a function
  printPerson(pers2)

	// Create Pers3 info by calling a function, infered type by Go
	var pers3 = createPerson("Ignacio", 25, "Engineer", 8000)
	// Print Pers3 info by calling a function
	printPerson(pers3)
}

func createPerson(name string, age int, job string, salary int) (p Person) {
	var per Person
	// Pers specification
  per.name = name
  per.age = age
  per.job = job
  per.salary = salary
	return per
}

// Pass Struct as Function Arguments
func printPerson(pers Person) {
	//Access and print Pers info
  fmt.Println("Name: ", pers.name)
  fmt.Println("Age: ", pers.age)
  fmt.Println("Job: ", pers.job)
  fmt.Println("Salary: ", pers.salary)
}