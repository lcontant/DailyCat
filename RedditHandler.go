package main

import (
	"encoding/json"
)

type RedditJsonChildren struct {
	Kind string `json:"kind"`
	Data map[string]string `json:"data"`
}
func getTopImageFromSubreddit(subreddit string) string {
	requestUrl := "https://www.reddit.com/r/" + subreddit + "/.json?Count=20"
	bytes := Get(requestUrl)
	var parsedResponse interface{}
	json.Unmarshal(bytes, &parsedResponse)
	data := parsedResponse.(map[string]interface{})
	children := data["data"].(map[string]interface{})["children"].([]RedditJsonChildren)
	return children[0].Data["url"]
}
