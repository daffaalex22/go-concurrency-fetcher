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
	EpisodeCount   int32  `json:"episodeCount"`
	EpisodeLength  int32  `json:"episodeLength"`
	ShowType       string `json:"showType"`
	RatingRank     int32  `json:"ratingRank"`
	PopularityRank int32  `json:"popularityRank"`
}

type Anime struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes AnimeAttribute `json:"attributes"`
}

type Body struct {
	Data Anime `json:"data"`
}
