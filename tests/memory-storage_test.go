package main

import (
	"fmt"

	"github.com/badico-cloud-hub/battery-go/storages"
)

func main() {
	valueA := "teste"

	valueB := 123

	var valueC []interface{}
	valueC = append(valueC, valueB)
	valueC = append(valueC, valueA)
	fmt.Println("A: ", valueA)
	fmt.Println("B: ", valueB)
	fmt.Println("C: ", valueC)

	myStorage := storages.New()

	myStorage.Set("A", valueA)
	fmt.Println("storage", myStorage)

	myStorage.Set("B", valueB)
	fmt.Println("storage", myStorage)

	myStorage.Set("C", valueC)
	fmt.Println("storage", myStorage)

	returnedA, errA := myStorage.Get("A")
	if errA != nil {
		fmt.Println(errA)
	} else {
		fmt.Println("item A found")
		fmt.Println(returnedA)
	}

	returnedB, errB := myStorage.Get("B")
	if errB != nil {
		fmt.Println(errB)
	} else {
		fmt.Println("item B found")
		fmt.Println(returnedB)
	}

	returnedC, errC := myStorage.Get("C")
	if errC != nil {
		fmt.Println(errC)
	} else {
		fmt.Println("item C found")
		fmt.Println(returnedC)
	}

	returnedD, errD := myStorage.Get("D")
	if errD != nil {
		fmt.Println(errD)
	} else {
		fmt.Println("item D found")
		fmt.Println(returnedD)
	}
	fmt.Println("finish!")
}
