package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkCity(b *testing.B) {
	b.StopTimer()
	parseFiles()
	server := httptest.NewServer(createRouter())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city")
		res.Body.Close()
	}
}

func BenchmarkCityArea(b *testing.B) {
	b.StopTimer()
	parseFiles()
	server := httptest.NewServer(createRouter())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/city_area/花蓮縣")
		res.Body.Close()
	}
}

func BenchmarkStreetName(b *testing.B) {
	b.StopTimer()
	parseFiles()
	server := httptest.NewServer(createRouter())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(server.URL + "/street_name/花蓮縣/吉安鄉")
		res.Body.Close()
	}
}
