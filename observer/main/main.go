package main

import (
	"SDP/observer"
)

func main() {
	Samsung_S23 := observer.NewItem("Samsung S23")

	c1 := observer.NewCustomer("Maga", "maga@gmail.com")
	c2 := observer.NewCustomer("Beka", "beka@gmail.com")
	c3 := observer.NewCustomer("Era", "era@gmail.com")

	Samsung_S23.AddObserver(c1)
	Samsung_S23.AddObserver(c2)
	Samsung_S23.AddObserver(c3)

	Samsung_S23.PrintAllObserver()

}
