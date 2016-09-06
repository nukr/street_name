package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// CityMap ...
type CityMap map[string]map[string]CityArea

// CountryMap ...
type CountryMap map[string]CityMap

func main() {
	router := createRouter()
	http.ListenAndServe(":3456", router)
}

func createRouter() *httprouter.Router {
	countryMap := make(CountryMap)
	countryList := make(map[string][]string)
	parseFiles(countryMap, countryList)
	router := httprouter.New()
	router.GET("/healthz", healthCheck())
	router.GET("/list", getCountry(countryList))
	router.GET("/list/:country", getCity(countryMap))
	router.GET("/list/:country/:city", getCityArea(countryMap))
	router.GET("/list/:country/:city/:city_area", getStreetName(countryMap))
	return router
}

func healthCheck() httprouter.Handle {
	return func(
		w http.ResponseWriter,
		req *http.Request,
		_ httprouter.Params,
	) {
		h := w.Header()
		h.Add("Cache-Control", "no-cache, no-store, must-revalidate")
		h.Add("Pragma", "no-cache")
		h.Add("Expires", "0")
		fmt.Fprintf(w, "OK")
	}
}

func getCountry(countryList map[string][]string) httprouter.Handle {
	return func(
		w http.ResponseWriter,
		req *http.Request,
		_ httprouter.Params,
	) {
		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		acceptLang := strings.Split(strings.ToLower(req.Header.Get("Accept-Language")), ",")[0]
		if acceptLang == "" {
			acceptLang = "zh-tw"
		}
		j, _ := json.Marshal(countryList[acceptLang])
		fmt.Fprintf(w, string(j))
	}
}

func getCity(countryMap CountryMap) httprouter.Handle {
	return func(
		w http.ResponseWriter,
		req *http.Request,
		p httprouter.Params,
	) {
		country := p.ByName("country")
		if countryMap[country] == nil {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		} else {
			h := w.Header()
			h.Add("Content-Type", "application/json; charset=utf-8")
			var s []string
			for k := range countryMap[country] {
				s = append(s, k)
			}
			j, _ := json.Marshal(s)
			fmt.Fprintf(w, string(j))
		}
	}
}

func getCityArea(countryMap CountryMap) httprouter.Handle {
	return func(
		w http.ResponseWriter,
		req *http.Request,
		p httprouter.Params,
	) {
		country := p.ByName("country")
		city := p.ByName("city")

		cityMap := countryMap[country]
		cityArea := cityMap[city]

		if cityMap == nil || cityArea == nil {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		} else {
			h := w.Header()
			h.Add("Content-Type", "application/json; charset=utf-8")
			var s []string
			for k := range cityArea {
				s = append(s, k)
			}
			j, _ := json.Marshal(s)
			fmt.Fprintf(w, string(j))
		}
	}
}

func getStreetName(countryMap CountryMap) httprouter.Handle {
	return func(
		w http.ResponseWriter,
		req *http.Request,
		p httprouter.Params,
	) {
		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		countryName := p.ByName("country")
		cityName := p.ByName("city")
		cityAreaName := p.ByName("city_area")

		cityMap := countryMap[countryName]
		cityAreaMap := cityMap[cityName]

		if cityMap == nil || cityAreaMap == nil {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		} else {
			streetName := cityAreaMap[cityAreaName]
			j, _ := json.Marshal(streetName)
			fmt.Fprintf(w, string(j))
		}
	}
}

func parseCountryList(countryList map[string][]string) {
	countryListDir, err := os.Open("./country_list")
	if err != nil {
		log.Fatal(err)
	}
	countryListFilesInfo, err := countryListDir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range countryListFilesInfo {
		countryFile, err := os.Open("./country_list/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(countryFile)
		if err != nil {
			log.Fatal(err)
		}
		if b[len(b)-1] == byte(10) {
			b = b[:len(b)-1]
		}
		countryList[f.Name()] = strings.Split(string(b), "\n")
	}
}

func parseCountryMap(countryMap CountryMap) {
	countryMap["台灣"] = make(CityMap)
	cityMap := countryMap["台灣"]
	file, err := os.Open("./streetName")
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range finfo {
		streetFile, err := os.Open("./streetName/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		streetName, err := ioutil.ReadAll(streetFile)
		if err != nil {
			log.Fatal(err)
		}
		arr := strings.Split(f.Name(), "_")
		zip, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
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

func parseFiles(countryMap CountryMap, countryList map[string][]string) {
	parseCountryList(countryList)
	parseCountryMap(countryMap)
}
