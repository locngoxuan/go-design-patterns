package main

import (
	"fmt"
	"log"
)

type Element interface {
	Info() string
}

type Iterator interface {
	GetNext() Element
	HasMore() bool
}

type IteratorCollection interface {
	CreateIterator() Iterator
}

type User struct {
	Id int
}

func (u User) Info() string {
	return fmt.Sprintf("user id = %v", u.Id)
}

type UserCollection struct {
	users []User
}

func (u UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

type UserIterator struct {
	index int
	users []User
}

func (u *UserIterator) GetNext() Element {
	if u.HasMore() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

func (u *UserIterator) HasMore() bool {
	return u.index < len(u.users)
}

func main() {
	users := make([]User, 0)
	for i := 0; i < 10; i++ {
		users = append(users, User{Id: i + 1})
	}
	collection := &UserCollection{
		users: users,
	}
	itr := collection.CreateIterator()
	for itr.HasMore() {
		log.Println(itr.GetNext().Info())
	}
}
