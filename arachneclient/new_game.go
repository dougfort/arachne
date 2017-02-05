package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const method = "http"
const host = "localhost"
const port = 8000
const gamesPath = "arachne/games/"

func newGame() (gameData, error) {
	uri := fmt.Sprintf("%s://%s:%d/%s", method, host, port, gamesPath)
	resp, err := http.Post(uri, "application/octet-stream", nil)
	if err != nil {
		return gameData{}, fmt.Errorf("http POST failed: %s", err)
	}
	x, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return gameData{}, fmt.Errorf("read body failed: %s", err)
	}
	log.Printf("debug: newGame: body = %s", string(x))
	return gameData{}, fmt.Errorf("not implemented")
}
