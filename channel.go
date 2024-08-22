package main

type Channel struct {
	name      string
	observers []Observer
}

func (ch *Channel) addObserver(observer Observer) {
	ch.observers = append(ch.observers, observer)
}

func (ch *Channel) removeObserver(observer Observer) {
	for i, obs := range ch.observers {
		if obs == observer {
			ch.observers = append(ch.observers[:i], ch.observers[i+1:]...)
			return
		}
	}
}

func (ch *Channel) notifyObservers() {
	for _, obs := range ch.observers {
		obs.update(ch)
	}
}
