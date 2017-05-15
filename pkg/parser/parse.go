package parser

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nukr/street_name/pkg/types"
)

// LoadAndParse loads xml encoded with following format
// <?xml version="1.0" encoding="UTF-8"?>
// <dataroot xmlns:od="urn:schemas-microsoft-com:officedata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  xsi:noNamespaceSchemaLocation="Xml_10510.xsd" generated="2016-12-13T15:49:13">
// <Xml_10510>
// <欄位1>10058</欄位1>
// <欄位4>臺北市中正區</欄位4>
// <欄位2>八德路１段</欄位2>
// <欄位3>全</欄位3>
// </Xml_10510>
// </dataroot>
func LoadAndParse(path string) *types.Address {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bsXML, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	addr := parse(bsXML)
	patchCountry(addr)
	return addr
}

func parse(bsXML []byte) *types.Address {
	type XMLAddress struct {
		Zip            string `xml:"欄位1"`
		Street         string `xml:"欄位2"`
		Range          string `xml:"欄位3"`
		CountyDistrict string `xml:"欄位4"`
	}
	type Result struct {
		XMLName      xml.Name     `xml:"dataroot"`
		XMLAddresses []XMLAddress `xml:"Xml_10510"`
	}
	result := Result{}
	xml.Unmarshal(bsXML, &result)
	address := types.Address{
		Country:  make(map[string]*types.StringSet),
		County:   make(map[string]*types.StringSet),
		District: make(map[string]*types.StringSet),
		Street:   make(map[string]*types.StringSet),
		Zip:      make(map[string]*types.StringSet),
	}
	for _, xmladdr := range result.XMLAddresses {
		// TODO: avoid hardcode here
		country := "台灣"
		county := string([]rune(xmladdr.CountyDistrict)[0:3])
		district := string([]rune(xmladdr.CountyDistrict)[3:])
		if v, found := address.County[country]; found {
			v.Add(county)
		} else {
			address.County[country] = &types.StringSet{}
			address.County[country].Add(county)
		}
		if v, found := address.District[country+county]; found {
			v.Add(district)
		} else {
			address.District[country+county] = &types.StringSet{}
			address.District[country+county].Add(district)
		}
		if v, found := address.Street[country+xmladdr.CountyDistrict]; found {
			v.Add(xmladdr.Street)
		} else {
			address.Street[country+xmladdr.CountyDistrict] = &types.StringSet{}
			address.Street[country+xmladdr.CountyDistrict].Add(xmladdr.Street)
		}
		if v, found := address.Zip[country+xmladdr.CountyDistrict+xmladdr.Street+xmladdr.Range]; found {
			v.Add(xmladdr.Zip)
		} else {
			address.Zip[country+xmladdr.CountyDistrict+xmladdr.Street+xmladdr.Range] = &types.StringSet{}
			address.Zip[country+xmladdr.CountyDistrict+xmladdr.Street+xmladdr.Range].Add(xmladdr.Zip)
		}
	}
	return &address
}

func patchCountry(addr *types.Address) {
	filelist, _ := ioutil.ReadDir("country_list")
	for _, fileItem := range filelist {
		file, err := os.Open("country_list/" + fileItem.Name())
		if err != nil {
			log.Fatal(err)
		}
		bsCountryJSON, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		var countryJSON []string
		json.Unmarshal(bsCountryJSON, &countryJSON)
		lang := strings.Split(fileItem.Name(), ".")[0]
		for _, country := range countryJSON {
			if v, ok := addr.Country[lang]; ok {
				v.Add(country)
			} else {
				addr.Country[lang] = &types.StringSet{}
				addr.Country[lang].Add(country)
			}
		}
	}
}
