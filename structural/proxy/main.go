package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type User struct {
	id   int64
	name string
}

type Database struct {
	users []User
}

func (d *Database) AddUser(u User) error {
	d.users = append(d.users, u)
	return nil
}

func (d *Database) GetUser(id int64) (User, error) {
	//simulate network latency and database processing time
	r := rand.Int31n(1000) + 100
	<-time.After(time.Duration(r) * time.Millisecond)
	for _, u := range d.users {
		if u.id == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

type DatabaseProxy struct {
	Database
	cache map[int64]User
}

func (d *DatabaseProxy) AddUser(u User) error {
	return d.Database.AddUser(u)
}

func (d *DatabaseProxy) GetUser(id int64) (User, error) {
	u, ok := d.cache[id]
	if ok {
		return u, nil
	}
	u, err := d.Database.GetUser(id)
	if err != nil {
		return u, err
	}
	d.cache[id] = u
	return u, nil
}

func main() {
	users := []User{
		{
			id:   1,
			name: "David De Gea",
		},
		{
			id:   20,
			name: "Diago Dalot",
		},
		{
			id:   19,
			name: "Raphael Varane",
		},
		{
			id:   6,
			name: "Lisandro Martinez",
		},
		{
			id:   12,
			name: "Tyrell Malacia",
		},
		{
			id:   39,
			name: "Scott Mctominay",
		},
		{
			id:   14,
			name: "Christian Eriksen",
		},
		{
			id:   21,
			name: "Antony Santos",
		},
		{
			id:   8,
			name: "Bruno Fernandes",
		},
		{
			id:   25,
			name: "Jadon Sancho",
		},
		{
			id:   10,
			name: "Marcus Rashford",
		},
	}

	database := Database{}
	for _, user := range users {
		_ = database.AddUser(user)
	}
	databaseProx := DatabaseProxy{
		Database: database,
		cache:    map[int64]User{},
	}

	start := time.Now()
	for _, user := range users {
		_, err := database.GetUser(user.id)
		if err != nil {
			log.Fatalln(err)
		}
		_, _ = database.GetUser(user.id)
	}
	log.Printf("looking %v users (2 times) from database take %v ms", len(users), time.Since(start).Milliseconds())

	start = time.Now()
	for _, user := range users {
		_, err := databaseProx.GetUser(user.id)
		if err != nil {
			log.Fatalln(err)
		}
		_, _ = databaseProx.GetUser(user.id)
	}
	log.Printf("looking %v users (2 times) from database-proxy take %v ms", len(users), time.Since(start).Milliseconds())
}
