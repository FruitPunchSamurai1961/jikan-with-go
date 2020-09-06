package jikan

import (
	"strconv"
)

type Anime struct {
	MalID         int      `json:"mal_id"`
	Url           string   `json:"url"`
	ImageUrl      string   `json:"image_url"`
	Title         string   `json:"title"`
	TitleEnglish  string   `json:"title_english"`
	TitleSynonyms []string `json:"title_synonyms"`
	Type          string   `json:"type"`
	Source        string   `json:"source"`
	Episodes      int      `json:"episodes"`
	Status        string   `json:"status"`
	Airing        bool     `json:"airing"`
	Aired         Dates    `json:"aired"`
	Duration      string   `json:"duration"`
	Rating        string   `json:"rating"`
	Score         int      `json:"score"`
	ScoredBy      int      `json:"scored_by"`
	Rank          int      `json:"rank"`
	Popularity    int      `json:"popularity"`
	Members       int      `json:"members"`
	Favorites     int      `json:"favorites"`
	Synopsis      string   `json:"synopsis"`
	Background    string   `json:"background"`
	Premiered     string   `json:"premiered"`
	Broadcast     string   `json:"broadcast"`
	Related       Related  `json:"related"`
	Producers     []Item   `json:"producers"`
	Licensors     []Item   `json:"licensors"`
	Studios       []Item   `json:"studios"`
	Genres        []Item   `json:"genres"`
	OpeningThemes string   `json:"opening_themes"`
	EndingThemes  string   `json:"ending_themes"`
}

func GetAnimeById(id int) (Anime, error) {
	animeData := Anime{}
	request := "/anime" + strconv.Itoa(id)
	err := getData(request, &animeData)
	if err != nil {
		return Anime{}, err
	}
	return animeData, nil
}