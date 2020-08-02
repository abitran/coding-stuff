package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

var students []Student

func (s *students) get_info() {
	for _, p := range s {
		fmt.Printf("%v\n", p.FirstName)
		fmt.Printf("%v\n", p.LastName)
		fmt.Printf("%v\n", p.Address)
		fmt.Printf("\n")
	}
}

func main() {
	student := &Students{Lstudents: {FirstName: "Arthur", LastName: "Powell", Age: 34, Address: "Los Leones 80 depto 36B"}}

}
