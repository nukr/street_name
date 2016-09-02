package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// CityArea ...
type CityArea struct {
	Name       string   `json:"name"`
	Zip        int      `json:"zip"`
	StreetName []string `json:"street_name"`
}

var cityMap = make(map[string]map[string]CityArea)

func main() {
	parseFiles()
	router := createRouter()
	http.ListenAndServe(":3456", router)
}

func createRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/city", getCity)
	router.GET("/city_area/:city", getCityArea)
	router.GET("/street_name/:city/:city_area", getStreetName)
	return router
}

func getCity(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h := w.Header()
	h.Add("Content-Type", "application/json")
	var s []string
	for k := range cityMap {
		s = append(s, k)
	}
	j, _ := json.Marshal(s)
	fmt.Fprintf(w, string(j))
}

func getCityArea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := w.Header()
	h.Add("Content-Type", "application/json")
	city := p.ByName("city")
	cityArea := cityMap[city]
	var s []string
	for k := range cityArea {
		s = append(s, k)
	}
	j, _ := json.Marshal(s)
	fmt.Fprintf(w, string(j))
}

func getStreetName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := w.Header()
	h.Add("Content-Type", "application/json")
	cityName := p.ByName("city")
	cityAreaName := p.ByName("city_area")
	streetName := cityMap[cityName][cityAreaName]
	j, _ := json.Marshal(streetName)
	fmt.Fprintf(w, string(j))
}

func parseFiles() {
	file, _ := os.Open("./streetName")
	finfo, _ := file.Readdir(0)
	for _, f := range finfo {
		streetFile, _ := os.Open("./streetName/" + f.Name())
		streetName, _ := ioutil.ReadAll(streetFile)
		arr := strings.Split(f.Name(), "_")
		zip, _ := strconv.Atoi(arr[0])
		city := arr[1]
		cityArea := strings.Replace(arr[2], ".txt", "", 1)
		if cityMap[city] == nil {
			cityMap[city] = make(map[string]CityArea)
		}
		cityMap[city][cityArea] = CityArea{
			Name:       cityArea,
			Zip:        zip,
			StreetName: strings.Split(string(streetName), ","),
		}
	}
}
