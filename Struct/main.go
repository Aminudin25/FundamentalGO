package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

//embeded struct
type Group struct {
	Name     string
	Admin    User
	Users    []User
	IsActive bool
}

func main() {
	user := User{}

	user.ID = 1
	user.FirstName = "Aminudin"
	user.LastName = "Nurdinillah"
	user.Email = "aminudinnurdinillah0@gmail.com"
	user.IsActive = true

	fmt.Println(user)

	user2 := User{
		ID:        2,
		FirstName: "Dede",
		LastName:  "Nurdini",
		Email:     "dedenurdini@gmail.com",
		IsActive:  true,
	}

	fmt.Println(user2)

	users := []User{user, user2}

	group := Group{
		Name:     "Gamer",
		Admin:    user,
		Users:    users,
		IsActive: true,
	}

	fmt.Println(group)
	displayGroup(group)

	//panggil method
	result := user.display()
	fmt.Println(result)
}

func displayGroup(group Group) {
	fmt.Printf("Group : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count :%d", len(group.Users))
	fmt.Println("")

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}
