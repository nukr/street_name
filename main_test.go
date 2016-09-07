package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCity(t *testing.T) {
	server := httptest.NewServer(createRouter())
	defer server.Close()
	res, _ := http.Get(server.URL + "/city")
	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if strings.Index(string(b), "花蓮縣") == -1 {
		t.Error("result data is not expected")
	}
}

func TestCountry(t *testing.T) {
	server := httptest.NewServer(createRouter())
	defer server.Close()

	f, err := os.Open("./country_list/zh-tw")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	if b[len(b)-1] == byte(10) {
		b = b[:len(b)-1]
	}
	s := strings.Split(string(b), "\n")
	expectedJSONBytes, _ := json.Marshal(s)
	expected := string(expectedJSONBytes)
	res, _ := http.Get(server.URL + "/list")
	apiReturnedBytes, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	apiReturnJSONString := string(apiReturnedBytes)
	if expected != apiReturnJSONString {
		t.Errorf("ERROR: expected %s, got %s", expected, apiReturnJSONString)
	}
}

func BenchmarkCity(b *testing.B) {
	b.StopTimer()
	server := httptest.NewServer(createRouter())
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city")
		res.Body.Close()
	}
}

func BenchmarkCityArea(b *testing.B) {
	b.StopTimer()
	server := httptest.NewServer(createRouter())
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city_area/花蓮縣")
		res.Body.Close()
	}
}

func BenchmarkStreetName(b *testing.B) {
	b.StopTimer()
	server := httptest.NewServer(createRouter())
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/street_name/花蓮縣/吉安鄉")
		res.Body.Close()
	}
}
