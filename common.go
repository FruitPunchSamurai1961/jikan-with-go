//these are all the structs that are common among some or most of the other requests made
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

//The main endpoint where we'll be adding on to in order to get data for the user
const endPoint = "https://api.jikan.moe/v3"

var myClient = &http.Client{Timeout: 10 * time.Second}

//The function that will be used to access the API since all of the methods are gets and implementing different structs
func getData(url string, target interface{}) error {
	r, err := myClient.Get(endPoint + url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//shared by anime and manga struct
type Prop struct {
	From struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}
	To struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}
}

//shared by anime and manga struct
type Dates struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Prop   Prop   `json:"prop"`
	String string `json:"string"`
}

//Shared by Manga and Anime structs
type Related struct {
	Adaptation []Item `json:"Adaptation"`
	SideStory  []Item `json:"Side story"`
	Summary    []Item `json:"Summary"`
}

//Shared by countless ... too many to name
type Item struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}
