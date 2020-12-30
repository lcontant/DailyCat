package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const API_ENDPOINT = "https://discord.com/api/v6"
const GRANT_TYPE = "authorization_code"
const SCOPE = "bot"

type DiscordHandler struct {
	configuration Configuration
}

type DiscordCodeExchangeRequestBody struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	Permission   int    `json:"permission"`
}

type DiscordCreateMessageRequestBody struct {
	Content string `json:"content"`
	Tts bool `json:"tts"`
	File []byte `json:"file"`
}

type DiscordCodeExchangeResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}



func (DiscordHandler) exchange_code(code string) DiscordCodeExchangeResponse {
	request_data := url.Values{}
	request_data.Set("client_id", discordHandler.configuration.values["CLIENT_ID"])
	request_data.Set("client_secret", discordHandler.configuration.getStringValue("CLIENT_SECRET"))
	request_data.Set("grant_type", GRANT_TYPE)
	request_data.Set("code", code)
	request_data.Set("redirect_uri", discordHandler.configuration.getStringValue("REDIRECT_URI"))
	request_data.Set("permission", discordHandler.configuration.getStringValue("PERMISSION"))
	request_url := API_ENDPOINT + "/oauth2/token"
	resp, _ := http.Post(request_url, "application/x-www-form-urlencoded", strings.NewReader(request_data.Encode()))
	defer resp.Body.Close()
	defer resp.Request.Body.Close()
	var bytes []byte
	resp.Request.Body.Read(bytes)
	raw_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	fmt.Println(string(raw_body))
	parsed_response := DiscordCodeExchangeResponse{}
	if resp.StatusCode == 200 {
		json.Unmarshal(raw_body, &parsed_response)
		discordHandler.configuration.values["access_token"] = parsed_response.AccessToken
		discordHandler.configuration.saveConfig()
	}
	return parsed_response
}

func (DiscordHandler) sendMessage(msg string) {

}

