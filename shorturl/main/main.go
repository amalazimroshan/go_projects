package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"shorturl/handler"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func main() {
	yamlFile := flag.String("yaml", "", "path to yaml file")
	jsonFile := flag.String("json", "", "path to json file")
	flag.Parse()

	mux := defaultMux()
	var handlerFunction http.Handler

	if *jsonFile != "" {
		json, err := os.ReadFile(*jsonFile)
		if err != nil {
			log.Fatalf("file error %v", err)
		}
		jsonHandler, err := handler.JSONHandler([]byte(json), mux)
		if err != nil {
			panic(err)
		}
		handlerFunction = jsonHandler
	} else if *yamlFile != "" {
		yaml, err := os.ReadFile(*yamlFile)
		if err != nil {
			panic("file error!")
		}
		yamlHandler, err := handler.YAMLHandler([]byte(yaml), mux)
		if err != nil {
			panic(err)
		}
		handlerFunction = yamlHandler
	} else {
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}
		handlerFunction = handler.MapHandler(pathsToUrls, mux)
	}

	http.ListenAndServe(":8090", handlerFunction)
}
