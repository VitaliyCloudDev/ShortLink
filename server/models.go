package server

type Link struct {
	ID       int    `json:"id"`
	UrlOrig  string `json:"url"`
	ShortUrl string `json:"ShortUrl"`
	Time     string `json:"Time"`
}

// ID 0 for testing
var links = []Link{
	{ID: 0,
		UrlOrig:  "https://testlink.com/database/index012345",
		ShortUrl: "https://shortlink.com/0"},
	{ID: 100000,
		UrlOrig:  "https://testlink.com/database/index012345",
		ShortUrl: "FMA3SCnI"},
}
