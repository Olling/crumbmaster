package main

import (
	"errors"
)


type User struct {
	Name string
}

type Queue struct {
	Users []User
}


func (queue *Queue) AddUser (user User) {
	queue.Users = append(queue.Users,user)
}


func (queue *Queue) GetUser (name string) (int, User) {
	for i,u := range queue.Users {
		if u.Name == name {
			return i,u
		}
	}
	var user User
	return -1,user
}


func (queue *Queue) RemoveUser (user User) {
	for i,u := range queue.Users {
		if u.Name == user.Name {
			queue.Users = append(queue.Users[:i], queue.Users[i+1:]...)
		}
	}
}


func (queue *Queue) SwapUsers (user1 User,user2 User) (err error) {
	ui1,_ :=  queue.GetUser(user1.Name)
	ui2,_ :=  queue.GetUser(user2.Name)

	if ui1 != -1 && ui2 != -1 {
		queue.Users[ui1], queue.Users[ui2] = queue.Users[ui2], queue.Users[ui1]
	} else {
		err = errors.New("The user(s) was not found")
	}
	return err
}


