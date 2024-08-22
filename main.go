package main 

func main(){
	youtubeCh:= Channel{name: "youtube"}
	newsCh:= Channel{name: "news"}
	user1:=&User{name: "amir"}
	user2:=&User{name: "ali"}
	user3:=&User{name: "erfan"}
	youtubeCh.addObserver(user1)
	youtubeCh.addObserver(user2)
	newsCh.addObserver(user1)
	newsCh.addObserver(user2)
	newsCh.addObserver(user3)
	newsCh.removeObserver(user2)
      youtubeCh.notifyObservers()
	newsCh.notifyObservers()
}