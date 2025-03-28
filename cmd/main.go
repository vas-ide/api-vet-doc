package main

import "fmt"


type Emploue struct {
	Name string
	Age int
}


func main() {
	vas := &Emploue{"VAS",38}
	fmt.Println(vas)
	fmt.Printf("Hello %v.\n", vas.Name)
	fmt.Printf("Can emploue %s drink bear ? - %b.\t It's age = %d", vas.Name, canDrink(&vas.Age), *getAge(*vas))
}


func canDrink(age *int) bool {
	return *age >=18
}

func getAge(e Emploue) *int{
	return &e.Age
}

