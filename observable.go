package main

type Observable interface {
	addObserver(observer  Observer)
	removeObserver(observer  Observer) 
	notifyObservers()
}
