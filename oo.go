package main

import "fmt"

type Person struct {
	Name 	string
  Age		int
}

type Staff struct {
	Position string
	Person Person
}

func main(){
	fmt.Println("Hello go")
  person := Person{Name: "Hans", Age: 20} 
	fmt.Println(person.Name)
	person.Print("This is my message.")
  staff := Staff{Position: "CEO", Person: person}
	staff.Print()
}

func (person Person) Print(message string){
  fmt.Printf("Hello %s, this is your message: %s\n", person.Name, message)
}

func (staff Staff) Print(){
	staff.Person.Print("Staff printing")
	fmt.Printf("Position: %s\n", staff.Position)
}