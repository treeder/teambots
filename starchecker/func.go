package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/nlopes/slack"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fnRes, _, err := client.Repositories.Get(ctx, "fnproject", "fn")
	if err != nil {
		panic(nil)
	}
	openfaasRes, _, err := client.Repositories.Get(ctx, "openfaas", "faas")
	if err != nil {
		panic(nil)
	}
	fnStars := *fnRes.StargazersCount
	openFaasStars := *openfaasRes.StargazersCount

	api := slack.New(os.Getenv("FIN_SLACK_KEY"))

	params := slack.PostMessageParameters{
		AsUser: true,
	}

	behind := openFaasStars - fnStars

	var b bytes.Buffer
	b.WriteString("Fn Stars: " + strconv.Itoa(fnStars) + "\nOpenFaas Stars: " + strconv.Itoa(openFaasStars))
	b.WriteString("\nWe're " + strconv.Itoa(behind) + " stars behind. Get to work!")

	_, _, err = api.PostMessage("demostream", b.String(), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
