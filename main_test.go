package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCity(t *testing.T) {
	cityMap := make(CityMap)
	parseFiles(cityMap)
	server := httptest.NewServer(createRouter(cityMap))
	defer server.Close()
	res, _ := http.Get(server.URL + "/city")
	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if strings.Index(string(b), "花蓮縣") == -1 {
		t.Error("result data is not expected")
	}
}

func BenchmarkCity(b *testing.B) {
	b.StopTimer()
	cityMap := make(CityMap)
	parseFiles(cityMap)
	server := httptest.NewServer(createRouter(cityMap))
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city")
		res.Body.Close()
	}
}

func BenchmarkCityArea(b *testing.B) {
	b.StopTimer()
	cityMap := make(CityMap)
	parseFiles(cityMap)
	server := httptest.NewServer(createRouter(cityMap))
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city_area/花蓮縣")
		res.Body.Close()
	}
}

func BenchmarkStreetName(b *testing.B) {
	b.StopTimer()
	cityMap := make(CityMap)
	parseFiles(cityMap)
	server := httptest.NewServer(createRouter(cityMap))
	defer server.Close()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/street_name/花蓮縣/吉安鄉")
		res.Body.Close()
	}
}
