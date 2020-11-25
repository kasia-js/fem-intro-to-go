package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// YOUR CODE HERE
var u = User {
	ID: 1,
	FirstName: "Kasia",
	LastName: "Kolny",
	Email: "kolnykatarzyna@gmail.com",
}
func updateEmail(u *User) { //passing address
	u.Email = "kasia@gmail.com" // no need for * since struct
	fmt.Println("updated Email:", u.Email)
}
func main() {
		fmt.Println("Pointers!")
		updateEmail(&u) //invoking passing address
}

