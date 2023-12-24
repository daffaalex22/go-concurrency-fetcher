package main

type Job struct {
	number int
}

type AnimeAttribute struct {
	Synopsis       string `json:"synopsis"`
	Description    string `json:"description"`
	Rating         string `json:"averageRating"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EpisodeCount   int    `json:"episodeCount"`
	EpisodeLength  int    `json:"episodeLength"`
	ShowType       string `json:"showType"`
	RatingRank     int    `json:"ratingRank"`
	PopularityRank int    `json:"popularityRank"`
}

type Anime struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes AnimeAttribute `json:"attributes"`
}

type Body struct {
	Data Anime `json:"data"`
}
