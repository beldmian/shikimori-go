package shikimorigo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

// * Anime

// GetAnimes ...
func GetAnimes(query SearchParams) ([]Anime, error) {
	v := url.Values{}
	// TODO: Add muti-params support
	v.Set("page", strconv.Itoa(query.Page))
	v.Set("limit", strconv.Itoa(query.Limit))
	v.Set("order", query.Order)
	v.Set("kind ", query.Kind)
	v.Set("status", query.Status)
	v.Set("season", query.Season)
	v.Set("score", fmt.Sprintf("%f", query.Score))
	v.Set("duration", query.Duration)
	v.Set("rating", query.Rating)
	v.Set("genre", strconv.Itoa(query.Genre))
	v.Set("studio", strconv.Itoa(query.Studio))
	v.Set("franchise", query.Franchise)
	v.Set("censoured", fmt.Sprintf("%t", query.Censoured))
	v.Set("ids", strconv.Itoa(query.IDs))
	v.Set("exclude_ids", strconv.Itoa(query.ExcludeIDs))
	v.Set("search", query.Search)

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

// GetAnimeID ...
func GetAnimeID(id int) (Anime, error) {
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

// GetSimilarAnime ...
func GetSimilarAnime(id int) ([]Anime, error) {
	resp, err := http.Get("https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/similar")
	if err != nil {
		return []Anime{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Anime{}, err
	}
	anime := []Anime{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return []Anime{}, err
	}
	return anime, nil
}

// GetAnimeRelatedObjects ...
func GetAnimeRelatedObjects(id int) ([]RelatedObject, error) {
	resp, err := http.Get("https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/related")
	if err != nil {
		return []RelatedObject{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []RelatedObject{}, err
	}
	anime := []RelatedObject{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return []RelatedObject{}, err
	}
	return anime, nil
}

// GetAnimeScreenshots ...
func GetAnimeScreenshots(id int) ([]Image, error) {
	resp, err := http.Get("https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/screenshots")
	if err != nil {
		return []Image{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Image{}, err
	}
	anime := []Image{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return []Image{}, err
	}
	return anime, nil
}

// * Mangas

// GetMangas ...
func GetMangas(query SearchParams) ([]Manga, error) {
	v := url.Values{}
	// TODO: Add muti-params support
	v.Set("page", strconv.Itoa(query.Page))
	v.Set("limit", strconv.Itoa(query.Limit))
	v.Set("order", query.Order)
	v.Set("kind ", query.Kind)
	v.Set("status", query.Status)
	v.Set("season", query.Season)
	v.Set("score", fmt.Sprintf("%f", query.Score))
	v.Set("duration", query.Duration)
	v.Set("rating", query.Rating)
	v.Set("genre", strconv.Itoa(query.Genre))
	v.Set("studio", strconv.Itoa(query.Studio))
	v.Set("franchise", query.Franchise)
	v.Set("censoured", fmt.Sprintf("%t", query.Censoured))
	v.Set("ids", strconv.Itoa(query.IDs))
	v.Set("exclude_ids", strconv.Itoa(query.ExcludeIDs))
	v.Set("search", query.Search)

	resp, err := http.Get("https://shikimori.one/api/mangas" + v.Encode())
	if err != nil {
		return []Manga{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Manga{}, err
	}
	manga := []Manga{}
	if err := json.Unmarshal(body, &manga); err != nil {
		return nil, err
	}
	return manga, nil
}

// GetMangaID ...
func GetMangaID(id int) (Manga, error) {
	resp, err := http.Get("https://shikimori.one/api/mangas/" + strconv.Itoa(id))
	if err != nil {
		return Manga{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Manga{}, err
	}
	manga := Manga{}
	if err := json.Unmarshal(body, &manga); err != nil {
		return Manga{}, err
	}
	return manga, nil
}

// GetSimilarManga ...
func GetSimilarManga(id int) ([]Manga, error) {
	resp, err := http.Get("https://shikimori.one/api/mangas/" + strconv.Itoa(id) + "/similar")
	if err != nil {
		return []Manga{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Manga{}, err
	}
	manga := []Manga{}
	if err := json.Unmarshal(body, &manga); err != nil {
		return []Manga{}, err
	}
	return manga, nil
}

// GetMangaRelatedObjects ...
func GetMangaRelatedObjects(id int) ([]RelatedObject, error) {
	resp, err := http.Get("https://shikimori.one/api/mangas/" + strconv.Itoa(id) + "/related")
	if err != nil {
		return []RelatedObject{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []RelatedObject{}, err
	}
	relatedobject := []RelatedObject{}
	if err := json.Unmarshal(body, &relatedobject); err != nil {
		return []RelatedObject{}, err
	}
	return relatedobject, nil
}

// * Ranobe

// GetRanobes ...
func GetRanobes(query SearchParams) ([]Ranobe, error) {
	v := url.Values{}
	// TODO: Add muti-params support
	v.Set("page", strconv.Itoa(query.Page))
	v.Set("limit", strconv.Itoa(query.Limit))
	v.Set("order", query.Order)
	v.Set("kind ", query.Kind)
	v.Set("status", query.Status)
	v.Set("season", query.Season)
	v.Set("score", fmt.Sprintf("%f", query.Score))
	v.Set("duration", query.Duration)
	v.Set("rating", query.Rating)
	v.Set("genre", strconv.Itoa(query.Genre))
	v.Set("studio", strconv.Itoa(query.Studio))
	v.Set("franchise", query.Franchise)
	v.Set("censoured", fmt.Sprintf("%t", query.Censoured))
	v.Set("ids", strconv.Itoa(query.IDs))
	v.Set("exclude_ids", strconv.Itoa(query.ExcludeIDs))
	v.Set("search", query.Search)

	resp, err := http.Get("https://shikimori.one/api/ranobe" + v.Encode())
	if err != nil {
		return []Ranobe{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Ranobe{}, err
	}
	ranobe := []Ranobe{}
	if err := json.Unmarshal(body, &ranobe); err != nil {
		return nil, err
	}
	return ranobe, nil
}

// GetRanobeID ...
func GetRanobeID(id int) (Ranobe, error) {
	resp, err := http.Get("https://shikimori.one/api/ranobe/" + strconv.Itoa(id))
	if err != nil {
		return Ranobe{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Ranobe{}, err
	}
	ranobe := Ranobe{}
	if err := json.Unmarshal(body, &ranobe); err != nil {
		return Ranobe{}, err
	}
	return ranobe, nil
}

// GetSimilarRanobe ...
func GetSimilarRanobe(id int) ([]Ranobe, error) {
	resp, err := http.Get("https://shikimori.one/api/ranobe/" + strconv.Itoa(id) + "/similar")
	if err != nil {
		return []Ranobe{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Ranobe{}, err
	}
	ranobe := []Ranobe{}
	if err := json.Unmarshal(body, &ranobe); err != nil {
		return []Ranobe{}, err
	}
	return ranobe, nil
}

// GetRanobeRelatedObjects ...
func GetRanobeRelatedObjects(id int) ([]RelatedObject, error) {
	resp, err := http.Get("https://shikimori.one/api/ranobe/" + strconv.Itoa(id) + "/related")
	if err != nil {
		return []RelatedObject{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []RelatedObject{}, err
	}
	relatedobject := []RelatedObject{}
	if err := json.Unmarshal(body, &relatedobject); err != nil {
		return []RelatedObject{}, err
	}
	return relatedobject, nil
}

// Auth ...
func Auth(ClientID, ClientSecret string) (User, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"user_rates"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://shikimori.one/oauth/authorize",
			TokenURL: "https://shikimori.one/oauth/token",
		},
	}
	param := oauth2.SetAuthURLParam("grant_type", "authorization_code")
	redir := oauth2.SetAuthURLParam("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	url := conf.AuthCodeURL("state", param, redir)
	fmt.Printf("Visit the URL for the auth dialog: %v \n", url)

	if err := open.Run(url); err != nil {
		return User{}, err
	}

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return User{}, err
	}
	tok, err := conf.Exchange(ctx, code, redir)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("GET", "https://shikimori.one/api/users/whoami", nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Set("User-Agent", "ShikiCLI")
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)

	clientTemp := &http.Client{}
	respTemp, err := clientTemp.Do(req)
	if err != nil {
		return User{}, err
	}
	bodyTemp, err := ioutil.ReadAll(respTemp.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	if err := json.Unmarshal(bodyTemp, &user); err != nil {
		return User{}, err
	}

	client := conf.Client(ctx, tok)
	user.Client = client
	return user, nil
}

// GetRateByID ...
func (u User) GetRateByID(id int) (UserRate, error) {
	resp, err := u.Client.Get("https://shikimori.one/api/v2/user_rates" + strconv.Itoa(id))
	if err != nil {
		return UserRate{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UserRate{}, err
	}
	var rate UserRate
	if err := json.Unmarshal(body, &rate); err != nil {
		return UserRate{}, err
	}
	return rate, nil
}

// GetRates ...
func (u User) GetRates() ([]UserRate, error) {
	resp, err := u.Client.Get("https://shikimori.one/api/v2/user_rates")
	if err != nil {
		return []UserRate{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []UserRate{}, err
	}
	var rate []UserRate
	if err := json.Unmarshal(body, &rate); err != nil {
		return []UserRate{}, err
	}
	return rate, nil
}
