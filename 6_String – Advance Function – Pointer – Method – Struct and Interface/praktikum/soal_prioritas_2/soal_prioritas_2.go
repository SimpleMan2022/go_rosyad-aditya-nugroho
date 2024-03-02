package main

import (
	"fmt"
)

type Vehicle struct {
	Name     string
	Merk     string
	TypeName string
	Price    int
}

type Customer struct {
	Name    string
	Address string
	Vehicle []Vehicle
}

func (c *Customer) DisplayCustomerData() {
	fmt.Println("Name: ", c.Name)
	fmt.Println("Address: ", c.Address)

}

var rentedVehicles = make(map[string]bool)

func (c *Customer) RentalVehicle(vehicles []Vehicle) {
	for _, vehicle := range vehicles {
		if rentedVehicles[vehicle.Name] {
			fmt.Printf("Vehicle: %s is already rented\n", vehicle.Name)
		} else {
			fmt.Printf("Thank you for rent the vehicle: %s\n", vehicle.Name)
			rentedVehicles[vehicle.Name] = true
			c.Vehicle = append(c.Vehicle, vehicle)
		}
	}
}

func (c *Customer) ReturnVehicle(vehicles []Vehicle) {
	var found bool = false
	for _, vehicle := range vehicles {
		rentedVehicles[vehicle.Name] = false
		for i, v := range c.Vehicle {
			if v.Name == vehicle.Name {
				found = true
				fmt.Println("Thank you for returning the vehicle: ", vehicle.Name)
				c.Vehicle = append(c.Vehicle[:i], c.Vehicle[i+1:]...)
			}
		}
		if !found {
			fmt.Printf("Vehicle: %s not found\n", vehicle.Name)
		}
	}
}

func (c *Customer) showRentedVehicles() {
	for _, vehicle := range c.Vehicle {
		fmt.Printf("list vehicle : %s\n", vehicle.Name)
	}
}

func main() {
	vehicle1 := Vehicle{
		Name:     "Toyota avanza ",
		Merk:     "Toyota",
		TypeName: "Sedan",
		Price:    1000000,
	}

	vehicle2 := Vehicle{
		Name:     "Honda Brio ",
		Merk:     "Honda",
		TypeName: "Sedan",
		Price:    1500000,
	}

	vehicle3 := Vehicle{
		Name:     "Honda Civic ",
		Merk:     "Honda",
		TypeName: "Sedan",
		Price:    2500000,
	}

	customer1 := Customer{
		Name:    "Rosyad Aditya Nugroho",
		Address: "Jalan Binjai 1",
		Vehicle: []Vehicle{},
	}

	customer2 := Customer{
		Name:    "Jamaluddin Malik",
		Address: "Jalan Jambi no 5",
		Vehicle: []Vehicle{},
	}

	fmt.Println("Customer Data : ")
	customer1.DisplayCustomerData()
	fmt.Println("")
	fmt.Println("Rental Vehicle : ")
	customer1.RentalVehicle([]Vehicle{vehicle1, vehicle2})
	fmt.Println("")
	fmt.Println("List of Rented Vehicles : ")
	customer1.showRentedVehicles()
	fmt.Println("")
	fmt.Println("Try a vehicle that has been rented: ")
	customer2.RentalVehicle([]Vehicle{vehicle1})
	fmt.Println(customer2.Vehicle)
	fmt.Println("")
	fmt.Println("Return a vehicle that has been rented: ")
	customer1.ReturnVehicle([]Vehicle{vehicle1})
	fmt.Println(customer1.Vehicle)
	fmt.Println("")
	fmt.Println("Try to return a vehicle that has not been rented: ")
	customer1.ReturnVehicle([]Vehicle{vehicle3})
	fmt.Println(customer1.Vehicle)

}
