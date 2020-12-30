package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const API_ENDPOINT = "https://discord.com/api/v8"
const GRANT_TYPE = "authorization_code"
const SCOPE = "bot identify email connections webhook.incoming messages.read"

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
	Content string              `json:"content"`
	Tts     bool                `json:"tts"`
	Embeds 	[]map[string]map[string]string `json:"embeds"`
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
	raw_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(raw_body))
	parsed_response := DiscordCodeExchangeResponse{}
	json.Unmarshal(raw_body, &parsed_response)
	if resp.StatusCode == 200 {
		discordHandler.configuration.values["access_token"] = parsed_response.AccessToken
		discordHandler.configuration.saveConfig()
	}
	return parsed_response
}

func (DiscordHandler) sendMessage(msg string) {
	client := http.Client{}
	requestUrl := discordHandler.configuration.values["webhook"]
	requestBodyData := DiscordCreateMessageRequestBody{
		Content: "Here is your daily cat",
		Tts:     false,
		Embeds: []map[string]map[string]string {
			{
				"image" : {
					"url" : "https://icatcare.org/app/uploads/2018/07/Thinking-of-getting-a-cat.png",
				},
			},
		},
	}
	requestBody, _ := json.Marshal(requestBodyData)
	request, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	request.Header.Add("Authorization", "Bearer "+discordHandler.configuration.values["access_token"])
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("User-Agent", "DiscordBot (louiscontant.com, 1)")
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	defer resp.Request.Body.Close()
	var bytes []byte
	resp.Request.Body.Read(bytes)
	raw_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(raw_body))
}
