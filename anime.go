package jikan

import (
	"errors"
	"fmt"
	"strconv"
)

//struct for the Endpoint + "Anime/{ID}" url
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
	Score         float64  `json:"score"`
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
	OpeningThemes []string `json:"opening_themes"`
	EndingThemes  []string `json:"ending_themes"`
}

//struct for Endpoint + "/anime/{ID}/characters_staff
type AnimeCharactersStaff struct {
	Characters []struct {
		MalID       int    `json:"mal_id"`
		URL         string `json:"url"`
		ImageURL    string `json:"image_url"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		VoiceActors []struct {
			MalID    int    `json:"mal_id"`
			Name     string `json:"name"`
			URL      string `json:"url"`
			ImageURL string `json:"image_url"`
			Language string `json:"language"`
		} `json:"voice_actors"`
	} `json:"characters"`
	Staff []struct {
		MalID     int      `json:"mal_id"`
		URL       string   `json:"url"`
		Name      string   `json:"name"`
		ImageURL  string   `json:"image_url"`
		Positions []string `json:"positions"`
	} `json:"staff"`
}

//Struct for the episodes url: Endpoint + "/{ID}/episodes
type EpisodesList struct {
	EpisodesLastPage int             `json:"episodes_last_page"`
	Episodes         []EpisodeDetail `json:"episodes"`
}

type EpisodeDetail struct {
	EpisodeID     int    `json:"episode_id"`
	Title         string `json:"title"`
	TitleJapanese string `json:"title_japanese"`
	TitleRomanji  string `json:"title_romanji"`
	Aired         string `json:"aired"`
	Filler        bool   `json:"filler"`
	Recap         bool   `json:"recap"`
	VideoUrl      string `json:"video_url"`
	ForumUrl      string `json:"forum_url"`
}

type EpisodeRange struct {
	Start, End int
}

//News structs for the "/news" endpoint
type News struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Url        string `json:"url"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	AuthorName string `json:"author_name"`
	AuthorUrl  string `json:"author_url"`
	ForumUrl   string `json:"forum_url"`
	ImageUrl   string `json:"image_url"`
	Comments   int    `json:"comments"`
	Intro      string `json:"intro"`
}

//This struct is to handle errors and help the developer debug
type Error struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

//gets the anime details using Anime struct
func GetAnimeById(id int) (Anime, Error) {
	animeData := Anime{}
	request := "/anime/" + strconv.Itoa(id)
	err := getData(request, &animeData)
	if err != nil {
		badStatus := Error{}
		_ = getData(request, &badStatus)
		return Anime{}, badStatus
	}
	return animeData, Error{}
}

//gets the characters and the staff that worked on specified anime
func GetAnimeCharactersStaff(id int) (AnimeCharactersStaff, error) {
	animeCharactersStaffData := AnimeCharactersStaff{}
	url := fmt.Sprintf("/anime/%v/characters_staff", id)
	err := getData(url, &animeCharactersStaffData)
	if err != nil {
		return AnimeCharactersStaff{}, err
	}
	if len(animeCharactersStaffData.Characters) == 0 && len(animeCharactersStaffData.Staff) == 0 {
		return AnimeCharactersStaff{}, errors.New("The reason you got {[] []} as the data is one of the two following reasons:\n" +
			"1. The API just isn't cooperating, so it returns nothing. Nothing we can do here.\n" +
			"2. The ID just doesn't exist. Try the GetAnimeByID method to see if it gives you a better error code! :)")
	}
	return animeCharactersStaffData, nil
}

func GetEpisodeList(id int) (EpisodesList, error) {
	episodeList := EpisodesList{}
	url := fmt.Sprintf("/anime/%v/episodes", id)
	err := getData(url, &episodeList)
	if err != nil {
		return EpisodesList{}, err
	}
	i := 2
	for i <= episodeList.EpisodesLastPage {
		tmp := EpisodesList{}
		url := fmt.Sprintf("/anime/%v/episodes/%v", id, i)
		err := getData(url, &tmp)
		if err != nil {
			return EpisodesList{}, err
		}
		episodeList.Episodes = append(episodeList.Episodes, tmp.Episodes...)
		i++
	}
	return episodeList, nil
}

// function to get only the specific episodes from the given anime
func GetEpisodesRange(id int, episodeRange EpisodeRange) (EpisodesList, error) {
	i := 0
	if episodeRange.Start%100 == 0 {
		i = episodeRange.Start / 100
	} else {
		i = (episodeRange.Start / 100) + 1
	}
	res := EpisodesList{}
	if episodeRange.Start > episodeRange.End && episodeRange.End != 0 {
		return EpisodesList{}, errors.New("the End int must be greater than the Start int")
	}
	if episodeRange.End == 0 {
		url := fmt.Sprintf("/anime/%v/episodes/%v", id, i)
		err := getData(url, &res)
		if err != nil {
			return EpisodesList{}, err
		}
		for i <= res.EpisodesLastPage {
			i++
			tmp := EpisodesList{}
			url := fmt.Sprintf("/anime/%v/episodes/%v", id, i)
			err := getData(url, &tmp)
			if err != nil {
				return EpisodesList{}, err
			}
			res.Episodes = append(res.Episodes, tmp.Episodes...)
		}
	} else {
		end := episodeRange.End / 100
		for i <= end+1 {
			tmp := EpisodesList{}
			url := fmt.Sprintf("/anime/%v/episodes/%v", id, i)
			err := getData(url, &tmp)
			if err != nil {
				return EpisodesList{}, err
			}
			res.Episodes = append(res.Episodes, tmp.Episodes...)
			i++
		}
	}
	diff := episodeRange.Start - res.Episodes[0].EpisodeID
	res.Episodes = res.Episodes[diff:]
	if len(res.Episodes) == 0 {
		return EpisodesList{}, errors.New("please check your episode range as the API is giving an error of no episodes in that range for the specified anime")
	}
	return res, nil
}

//function to get the news related to the anime
func getNews(id int) (News, error) {
	res := News{}
	url := fmt.Sprintf("/anime/%v/news", id)
	err := getData(url, &res)
	if err != nil {
		return News{}, err
	}
	return res, nil
}
