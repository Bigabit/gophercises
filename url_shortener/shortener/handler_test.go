package urlshortener

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//	Default Mux for testing MapHandler
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}

func testDataJSON() string {
	return `
	[
		{
			"path": "/test",
			"url":  "https://blog.questionable.services/article/testing-http-handlers-go/"
		}
	]`
}

func testDataYAML() string {
	return `- path: /test
  url: https://blog.questionable.services/article/testing-http-handlers-go/
`
}

func testMapHandlerInit() http.HandlerFunc {
	//	Test data
	urls := map[string]string{
		"/test": "https://blog.questionable.services/article/testing-http-handlers-go/",
	}

	return MapHandler(urls, defaultMux())
}

func TestJsonHandled(t *testing.T) {
	//	Create request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	//	Create recorder
	recorder := httptest.NewRecorder()

	handler, err := FileHandler([]byte(testDataJSON()), "json", defaultMux())
	if err != nil {
		t.Error(err.Error())
	}
	handler.ServeHTTP(recorder, req)

	//	Assert status code 301
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusMovedPermanently,
		)
	}
}

func TestMapHandlerRedirects(t *testing.T) {

	//	Create request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	//	Create recorder
	recorder := httptest.NewRecorder()

	//	Initialize MapHandler
	handler := testMapHandlerInit()
	handler.ServeHTTP(recorder, req)

	//	Assert status code 301
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusMovedPermanently,
		)
	}
}

func TestParseJson(t *testing.T) {
	j := testDataJSON()
	data, err := parseJSON([]byte(j))
	if err != nil {
		t.Error(err.Error())
	}

	if _, ok := data["/test"]; !ok {
		t.Error(err.Error())
	}
}

func TestParseYAML(t *testing.T) {
	y := testDataYAML()
	result, err := parseYAML([]byte(y))
	if err != nil {
		t.Error(err.Error())
	}

	if _, ok := result["/test"]; !ok {
		t.Error(err.Error())
	}
}

func TestReadFile(t *testing.T) {
	expected := testDataYAML()
	got := Read("test.yml")

	if strings.EqualFold(expected, string(got)) {
		t.Errorf("Expected %s\n got %s", expected, got)
	}
}

func TestYAMLHandler(t *testing.T) {
	//	Create request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	//	Create recorder
	recorder := httptest.NewRecorder()

	handler, err := FileHandler([]byte(testDataYAML()), "yml", defaultMux())
	if err != nil {
		t.Error(err.Error())
	}
	handler.ServeHTTP(recorder, req)

	//	Assert status code 301
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusMovedPermanently,
		)
	}
}
