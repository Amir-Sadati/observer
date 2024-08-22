package main

import (
	"fmt"
)

type User struct {
	name string
}

func (user *User) update(observable Observable) {
	ch, ok := observable.(*Channel)
	if ok {
		fmt.Printf("user:%v received message from channel %v \n", user.name, ch.name)
	} else {
		fmt.Printf("user:%v Received update from unknown observable \n", user.name)
	}
}
