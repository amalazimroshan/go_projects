package main

import "fmt"
// func  main()  {
// 	mux := defaultMux()

// 	pathsToUrls := map[string]string{
// 		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
// 	}
// 	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
// }

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p:=person{name:name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Alice",20})
}