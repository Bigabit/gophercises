package urlshortener

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type url struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func makeMap(urls []url) map[string]string {
	urlsMap := make(map[string]string)

	for _, u := range urls {
		urlsMap[u.Path] = u.Url
	}

	return urlsMap
}

func parseJSON(data []byte) (map[string]string, error) {
	var urls []url
	err := json.Unmarshal(data, &urls)

	return makeMap(urls), err
}

func parseYAML(data []byte) (map[string]string, error) {
	var urls []url
	err := yaml.Unmarshal(data, &urls)

	return makeMap(urls), err
}

func Read(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
