package main

type Payload struct {
	SearchText string   `json:"search_text"`
	Tags       []string `json:"tags"`
	TagsMode   string   `json:"tags_mode"`
	Brands     []string `json:"brands"`
	Blacklist  []string `json:"blacklist"`
	OrderBy    string   `json:"order_by"`
	Ordering   string   `json:"ordering"`
	Page       int      `json:"page"`
}

type Response struct {
	Page        int    `json:"page"`
	NbPages     int    `json:"nbPages"`
	NbHits      int    `json:"nbHits"`
	HitsPerPage int    `json:"hitsPerPage"`
	Hits        string `json:"hits"`
}

type Hit struct {
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Titles       []string    `json:"titles"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Views        int         `json:"views"`
	Interests    int         `json:"interests"`
	PosterUrl    string      `json:"poster_url"`
	CoverUrl     string      `json:"cover_url"`
	Brand        string      `json:"brand"`
	BrandId      int         `json:"brand_id"`
	DurationInMs int         `json:"duration_in_ms"`
	IsCensored   bool        `json:"is_censored"`
	Rating       interface{} `json:"rating"`
	Likes        int         `json:"likes"`
	Dislikes     int         `json:"dislikes"`
	Downloads    int         `json:"downloads"`
	MonthlyRank  int         `json:"monthly_rank"`
	Tags         []string    `json:"tags"`
	CreatedAt    int         `json:"created_at"`
	ReleasedAt   int         `json:"released_at"`
}
