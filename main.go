package main

type Recipent struct {
	Name  string
	Email string
}

func main() {

	recipentChannel := make(chan Recipent)
	loadRecipent("./email.csv", recipentChannel)
}
