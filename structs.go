package shikimorigo

import (
	"net/http"
)

/*
	! Shikimori API v1
	TODO: Support v2
*/

//* Animes

// Anime ...
type Anime struct {
	ID                 int         `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Russian            interface{} `json:"russian,omitempty"`
	Image              Image       `json:"image,omitempty"`
	URL                string      `json:"url,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	Score              string      `json:"score,omitempty"`
	Status             string      `json:"status,omitempty"`
	Episodes           int         `json:"episodes,omitempty"`
	EpisodesAired      int         `json:"episodes_aired,omitempty"`
	AiredOn            string      `json:"aired_on,omitempty"`
	ReleasedOn         interface{} `json:"released_on,omitempty"`
	Rating             string      `json:"rating,omitempty"`
	English            []string    `json:"english,omitempty"`
	Japanese           []string    `json:"japanese,omitempty"`
	Synonyms           []string    `json:"synonyms,omitempty"`
	LicenseNameRu      string      `json:"license_name_ru,omitempty"`
	Duration           int         `json:"duration,omitempty"`
	Description        string      `json:"description,omitempty"`
	DescriptionHTML    string      `json:"description_html,omitempty"`
	DescriptionSource  string      `json:"description_source,omitempty"`
	Franchise          string      `json:"franchise,omitempty"`
	Favoured           bool        `json:"favoured,omitempty"`
	Anons              bool        `json:"anons,omitempty"`
	Ongoing            bool        `json:"ongoing,omitempty"`
	ThreadID           int         `json:"thread_id,omitempty"`
	TopicID            int         `json:"topic_id,omitempty"`
	MyanimelistID      int         `json:"myanimelist_id,omitempty"`
	RatesScoresStats   []Rate      `json:"rates_scores_stats,omitempty"`
	RatesStatusesStats []Status    `json:"rates_statuses_stats,omitempty"`
	UpdatedAt          string      `json:"updated_at,omitempty"`
	NextEpisodeAt      string      `json:"next_episode_at,omitempty"`
	Genres             []Genre     `json:"genres,omitempty"`
	Studios            []Studio    `json:"studios,omitempty"`
	Screenshots        []Image     `json:"screenshots,omitempty"`
}

// Image ...
type Image struct {
	Original string `json:"original,omitempty"`
	Preview  string `json:"preview,omitempty"`
	X96      string `json:"x96,omitempty"`
	X48      string `json:"x48,omitempty"`
}

// Rate ...
type Rate struct {
	Name  int `json:"name,omitempty"`
	Value int `json:"value,omitempty"`
}

// Status ...
type Status struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

// Genre ...
type Genre struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Russian string `json:"russian,omitempty"`
	Kind    string `json:"kind,omitempty"`
}

// Studio ...
type Studio struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	FilteredName string `json:"filtered_name,omitempty"`
	Real         bool   `json:"real,omitempty"`
	Image        Image  `json:"image,omitempty"`
}

// SearchParams ...
type SearchParams struct {
	Page       int
	Limit      int
	Order      string
	Kind       string
	Status     string
	Season     string
	Score      float32
	Duration   string
	Rating     string
	Genre      int
	Studio     int
	Franchise  string
	Censoured  bool
	IDs        int
	ExcludeIDs int
	Search     string
}

// RelatedObject ...
type RelatedObject struct {
	Relation        string `json:"relation,omitempty"`
	RelationRussian string `json:"relation_russian,omitempty"`
	Anime           Anime  `json:"anime,omitempty"`
	Manga           Manga  `json:"manga,omitempty"`
}

// * Mangas

// Manga ...
type Manga struct {
	ID                 int         `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Russian            string      `json:"russian,omitempty"`
	Image              Image       `json:"image,omitempty"`
	URL                string      `json:"url,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	Score              string      `json:"score,omitempty"`
	Status             string      `json:"status,omitempty"`
	Volumes            int         `json:"volumes,omitempty"`
	Chapters           int         `json:"chapters,omitempty"`
	AiredOn            string      `json:"aired_on,omitempty"`
	ReleasedOn         string      `json:"released_on,omitempty"`
	English            []string    `json:"english,omitempty"`
	Japanese           []string    `json:"japanese,omitempty"`
	Synonyms           []string    `json:"synonyms,omitempty"`
	LicenseNameRu      string      `json:"license_name_ru,omitempty"`
	Description        string      `json:"description,omitempty"`
	DescriptionHTML    string      `json:"description_html,omitempty"`
	DescriptionSource  string      `json:"description_source,omitempty"`
	Franchise          string      `json:"franchise,omitempty"`
	Favoured           bool        `json:"favoured,omitempty"`
	Anons              bool        `json:"anons,omitempty"`
	Ongoing            bool        `json:"ongoing,omitempty"`
	ThreadID           int         `json:"thread_id,omitempty"`
	TopicID            int         `json:"topic_id,omitempty"`
	MyanimelistID      int         `json:"myanimelist_id,omitempty"`
	RatesScoresStats   []Rate      `json:"rates_scores_stats,omitempty"`
	RatesStatusesStats []Status    `json:"rates_statuses_stats,omitempty"`
	Genres             []Genre     `json:"genres,omitempty"`
	Publishers         []Publisher `json:"publishers,omitempty"`
	UserRate           string      `json:"user_rate,omitempty"`
}

// Publisher ...
type Publisher struct {
	ID   int
	Name string
}

// * Ranobes

// Ranobe ...
type Ranobe struct {
	ID                 int         `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Russian            string      `json:"russian,omitempty"`
	Image              Image       `json:"image,omitempty"`
	URL                string      `json:"url,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	Score              string      `json:"score,omitempty"`
	Status             string      `json:"status,omitempty"`
	Volumes            int         `json:"volumes,omitempty"`
	Chapters           int         `json:"chapters,omitempty"`
	AiredOn            string      `json:"aired_on,omitempty"`
	ReleasedOn         string      `json:"released_on,omitempty"`
	English            []string    `json:"english,omitempty"`
	Japanese           []string    `json:"japanese,omitempty"`
	Synonyms           []string    `json:"synonyms,omitempty"`
	LicenseNameRu      string      `json:"license_name_ru,omitempty"`
	Description        string      `json:"description,omitempty"`
	DescriptionHTML    string      `json:"description_html,omitempty"`
	DescriptionSource  string      `json:"description_source,omitempty"`
	Franchise          string      `json:"franchise,omitempty"`
	Favoured           bool        `json:"favoured,omitempty"`
	Anons              bool        `json:"anons,omitempty"`
	Ongoing            bool        `json:"ongoing,omitempty"`
	ThreadID           int         `json:"thread_id,omitempty"`
	TopicID            int         `json:"topic_id,omitempty"`
	MyanimelistID      int         `json:"myanimelist_id,omitempty"`
	RatesScoresStats   []Rate      `json:"rates_scores_stats,omitempty"`
	RatesStatusesStats []Status    `json:"rates_statuses_stats,omitempty"`
	Genres             []Genre     `json:"genres,omitempty"`
	Publishers         []Publisher `json:"publishers,omitempty"`
	UserRate           string      `json:"user_rate,omitempty"`
}

// User ...
type User struct {
	ID           int    `json:"id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Image        Image  `json:"image,omitempty"`
	LastOnlineAt string `json:"last_online_at,omitempty"`
	Name         string `json:"name,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Website      string `json:"website,omitempty"`
	BirthOn      string `json:"birth_on,omitempty"`
	Locale       string `json:"locale,omitempty"`
	Client       *http.Client
}

// AvatarImage ...
type AvatarImage struct {
	X160 string `json:"x_160,omitempty"`
	X148 string `json:"x_148,omitempty"`
	X80  string `json:"x_80,omitempty"`
	X64  string `json:"x_64,omitempty"`
	X48  string `json:"x_48,omitempty"`
	X32  string `json:"x_32,omitempty"`
	X16  string `json:"x_16,omitempty"`
}

// UserRate ...
type UserRate struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	TargetID   int    `json:"target_id"`
	TargetType string `json:"target_type"`
	Score      int    `json:"score"`
	Status     string `json:"status"`
	Rewatches  int    `json:"rewatches"`
	Episodes   int    `json:"episodes"`
	Volumes    int    `json:"volumes"`
	Chapters   int    `json:"chapters"`
	Text       string `json:"text,omitempty"`
	TextHTML   string `json:"text_html,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
