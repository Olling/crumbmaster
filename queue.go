package main

import (
	"github.com/olling/slog"
	"errors"
)

type Queue struct {
	Users []User
}


func (queue *Queue) AddUser (user User) {
	queue.Users = append(queue.Users,user)
}


func (queue *Queue) GetUser (name string) (int, User) {
	slog.PrintTrace("GetUser:",name)
	for i,u := range queue.Users {
		if u.Name == name {
			return i,u
		}
	}
	var user User
	return -1,user
}


func (queue *Queue) RemoveUser (user User) {
	slog.PrintTrace("RemoveUser:",user)
	for i,u := range queue.Users {
		if u.Name == user.Name {
			queue.Users = append(queue.Users[:i], queue.Users[i+1:]...)
		}
	}
}


func (queue *Queue) MoveToBack (user User) {
	slog.PrintTrace("MoveToBack:",user)
	queue.RemoveUser(user)
	queue.AddUser(user)
}


func (queue *Queue) MoveFirstToBack () {
	slog.PrintTrace("MoveFirstToBack:")
	if len(queue.Users) > 1 {
		queue.MoveToBack(queue.Users[0])
	}
}

func (queue *Queue) GetResponsible() (user User) {
	if len(queue.Users) > 0 {
		return queue.Users[0]
	}
	return user
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


func (queue *Queue) WriteToDisk (path string) {
	WriteJsonFile(&queue, path)
}

func GetQueueFromDisk(path string) (queue Queue) {
	ReadJsonFile(path,&queue)
	return queue
}

func GetCurrentQueue() (queue Queue) {
	return GetQueueFromDisk(CurrentConfiguration.PathJsonConfiguration)
}

func (queue *Queue) Write () {
	queue.WriteToDisk(CurrentConfiguration.PathJsonConfiguration)
}
