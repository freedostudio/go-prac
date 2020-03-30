package dog

import "fmt"

// Dog : here you tell us what Salutation is
type Dog struct {
	Kind     string
	Name     string
	Location string
}

func (d *Dog) MakeSound() {
	fmt.Println("wang wang")
}

func (d *Dog) Run() {
	fmt.Println("dog run!")
}
