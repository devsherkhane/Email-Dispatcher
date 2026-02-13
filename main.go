package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipent struct {
	Name  string
	Email string
}

func main() {

	recipentChannel := make(chan Recipent)

	go func() {
		loadRecipent("./email.csv", recipentChannel)

	}()
	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipentChannel, &wg)

	}

	wg.Wait()

}

func executeTemplate(r Recipent) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
