package types

import (
	"encoding/json"
	"log"
)

// StringSet is guarantee order
type StringSet struct {
	data []string
	set  map[string]bool
}

// Add ...
func (set *StringSet) Add(s string) bool {
	if set.set == nil {
		set.set = make(map[string]bool)
	}
	_, found := set.set[s]
	if !found {
		set.data = append(set.data, s)
		set.set[s] = true
	}
	return !found
}

// ToJSON prints data in JSON format
func (set *StringSet) ToJSON() string {
	bsJSONData, err := json.Marshal(set.data)
	if err != nil {
		log.Println(err)
	}
	return string(bsJSONData)
}

// Address ...
type Address struct {
	Country  map[string]*StringSet `json:"country"`
	County   map[string]*StringSet `json:"county"`
	District map[string]*StringSet `json:"district"`
	Street   map[string]*StringSet `json:"street"`
	Zip      map[string]*StringSet `json:"zip"`
}
