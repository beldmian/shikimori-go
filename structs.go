package shikimorigo

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
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Russian    string `json:"russian,omitempty"`
	Image      Image  `json:"image,omitempty"`
	URL        string `json:"url,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Score      string `json:"score,omitempty"`
	Status     string `json:"status,omitempty"`
	Volumes    string `json:"volumes,omitempty"`
	Chapters   string `json:"chapters,omitempty"`
	AiredOn    string `json:"aired_on,omitempty"`
	ReleasedOn string `json:"released_on,omitempty"`
}
