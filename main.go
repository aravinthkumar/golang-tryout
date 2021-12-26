package main

import "fmt"

type user struct {
	ID        int
	FirstName string
}

func main() {

	user1 := user{1, "Aravinth"}
	user2 := user{2, "Rahul"}

	fmt.Println(user1, user2)
	if user1 == user2 {
		println("Same person")
	} else if user1.FirstName == user2.FirstName {
		println("Same name")
	} else {
		println("Different person")
	}

}
