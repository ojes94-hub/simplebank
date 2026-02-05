package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	p := Person{Name: "Alice Odido", Age: 30, Phone: "07069567577"}
	fmt.Println("Name:", p.GetName())
	fmt.Println("Age:", p.GetAge())
	fmt.Println("Phone:", p.GetPhone())
	fmt.Println("Updated Age:", p.GetAge())
	fmt.Println("Updated Person Info:", p.GetPerson())
	fmt.Printf("Full Info: %+v \n", p.FullInfo())

	reader := bufio.NewReader(os.Stdin)
	// newPhone, _ := reader.ReadString('\n')
	// newAge, _ := reader.ReadString('\n')
	fmt.Print("Enter 'a' to update Name, 'b' to update Age, 'c' to update Phone: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "a":
		fmt.Println("Enter User Name:")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)
		p = p.SetPerson(Person{Name: newName})
	case "b":
		fmt.Println("Enter User Age:")
		newAge, _ := reader.ReadString('\n')
		newAge = (strings.TrimSpace(newAge))
		newAgeInt, _ := strconv.Atoi(newAge)
		p = p.SetPerson(Person{Age: newAgeInt})
	case "c":
		fmt.Println("Enter User Phone:")
		newPhone, _ := reader.ReadString('\n')
		newPhone = strings.TrimSpace(newPhone)
		p = p.SetPerson(Person{Phone: newPhone})
	default:
		fmt.Println("Invalid option")
	}
	fmt.Println("Updated Person Info:", p.FullInfo())

}

type Person struct {
	Name  string
	Age   int
	Phone string
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) GetPhone() string {
	return p.Phone
}

func UpdatePerson(p *Person, name string, age int, phone string) {
	p.Name = name
	p.Age = age
	p.Phone = phone
}

func (p Person) GetPerson() string {
	return fmt.Sprintf("Name: %s, Age: %d, Phone: %s", p.Name, p.Age, p.Phone)
}

func (p Person) FullInfo() Person {
	return Person{
		Name:  p.Name,
		Age:   p.Age,
		Phone: p.Phone,
	}
}

func (p *Person) SetPerson(person Person) Person {
	if person.Name != "" {
		p.Name = person.Name
	}
	if person.Age != 0 {
		p.Age = person.Age
	}
	if person.Phone != "" {
		p.Phone = person.Phone
	}
	return *p
}
