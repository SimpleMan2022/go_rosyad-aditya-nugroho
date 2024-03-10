package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Name struct {
	FirstName string
	LastName  string
}

type User struct {
	ID       int
	Username string
	Email    string
	Name     Name
}

func main() {

	done := make(chan any)
	go func() {
		url := "https://fakestoreapi.com/users"
		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var users []User
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&users); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Users:")
		for _, user := range users {
			fmt.Println("ID:", user.ID)
			fmt.Println("Username:", user.Username)
			fmt.Println("Email:", user.Email)
			fmt.Println("First name:", user.Name.FirstName)
			fmt.Println("Last name:", user.Name.LastName)
			fmt.Println("====")
			fmt.Println("====")
		}

		done <- true
	}()
	<-done

}
