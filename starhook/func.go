package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/nlopes/slack"
)

type payloadIn struct {
	Action string `json:"action"`
	Sender sender `json:"sender"`
	Repo   repo   `json:"repository"`
}

type sender struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type repo struct {
	Name            string `json:"name"`
	StarGazersCount int    `json:"stargazers_count"`
}

func main() {
	p := new(payloadIn)
	json.NewDecoder(os.Stdin).Decode(p)

	api := slack.New(os.Getenv("FIN_SLACK_KEY"))
	params := slack.PostMessageParameters{
		AsUser: true,
	}

	var b bytes.Buffer
	b.WriteString(":star: *" + p.Sender.Login + "* starred the *" + p.Repo.Name + "* repo\n")
	b.WriteString("        Total stars now *" + strconv.Itoa(p.Repo.StarGazersCount) + "*")

	room := os.Getenv("ROOM_TO_POST")
	_, _, err := api.PostMessage(room, b.String(), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
