package main


import (
	"fmt"
)

func main() {
	test()
}

func test () {
	var user User
	user.Name="user"
	var user1 User
	user1.Name="user1"
	var user2 User
	user2.Name="user2"
	var user3 User
	user3.Name="user3"
	var user4 User
	user4.Name="user4"

	var q Queue
	q.AddUser(user)
	q.AddUser(user1)
	q.AddUser(user2)
	q.AddUser(user3)
	q.AddUser(user4)

	fmt.Println(q)
	q.RemoveUser(user2)
	fmt.Println(q)
	q.SwapUsers(user1,user4)
	fmt.Println(q)


	WriteJsonFile(q,"/home/olling/jsontest.json")

	var queue Queue
	ReadJsonFile("/home/olling/jsontest.json",&queue)
	fmt.Println(queue)
}
