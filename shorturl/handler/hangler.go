package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}

	err := yaml.Unmarshal([]byte(yml), &pathToUrls)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}

	err := json.Unmarshal([]byte(yml), &pathToUrls)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil

}
