package shikimorigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// GetAnimes ...
func GetAnimes(query AnimeParams) ([]Anime, error) {
	v := url.Values{}
	// TODO: Add muti-params support
	v.Set("Page", strconv.Itoa(query.Page))
	v.Set("Limit", strconv.Itoa(query.Limit))
	v.Set("Order", query.Order)
	v.Set("Kind ", query.Kind)
	v.Set("Status", query.Status)
	v.Set("Season", query.Season)
	v.Set("Score", fmt.Sprintf("%f", query.Score))
	v.Set("Duration", query.Duration)
	v.Set("Rating", query.Rating)
	v.Set("Genre", strconv.Itoa(query.Genre))
	v.Set("Studio", strconv.Itoa(query.Studio))
	v.Set("Franchise", query.Franchise)
	v.Set("Censoured", fmt.Sprintf("%t", query.Censoured))
	v.Set("IDs", strconv.Itoa(query.IDs))
	v.Set("ExcludeIDs", strconv.Itoa(query.ExcludeIDs))
	v.Set("Search", query.Search)

	resp, err := http.Get("https://shikimori.one/api/animes" + v.Encode())
	if err != nil {
		return []Anime{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Anime{}, err
	}
	anime := []Anime{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return nil, err
	}
	return anime, nil
}

// GetAnimesID ...
func GetAnimesID(id int) (Anime, error) {
	resp, err := http.Get("https://shikimori.one/api/animes/" + strconv.Itoa(id))
	if err != nil {
		return Anime{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Anime{}, err
	}
	anime := Anime{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return Anime{}, err
	}
	return anime, nil
}
