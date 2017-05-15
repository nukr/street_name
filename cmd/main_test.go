package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestList(t *testing.T) {
	server := httptest.NewServer(createRouter())
	defer server.Close()

	file, err := os.Open("./country_list/zh-tw.json")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	expected := string(b)
	res, _ := http.Get(server.URL + "/list")
	bsBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	actual := string(bsBody)
	if expected != actual {
		t.Errorf("%s expected %s", actual, expected)
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
