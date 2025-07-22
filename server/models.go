package server

type Link struct {
	ID      int    `json:"id"`
	UrlOrig string `json:"url"`
	ShortID string `json:"ShortUrl"`
	Time    string `json:"Time"`
}

// ID here for testing, will be removed on release
var links = []Link{
	{ID: 0,
		UrlOrig: "https://testlink.com/database/index012345",
		ShortID: "https://shortlink.com/0"},
	{ID: 100000,
		UrlOrig: "https://testlink.com/database/index012345",
		ShortID: "FMA3SCnI"},
}
