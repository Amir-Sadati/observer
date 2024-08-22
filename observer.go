package main

type Observer interface {
	update(observable Observable)
}
