package main

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) []byte{
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}