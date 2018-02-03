package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/nlopes/slack"
)

type PostMsg struct {
	TeamID string `param:"team_id"`
	UserID string `param:"user_id"`
	Text   string `param:"text"`
}

func main() {
	api := slack.New(os.Getenv("FIN_SLACK_KEY"))

	data, err := ioutil.ReadAll(os.Stdin)
	check(err)

	m, _ := url.ParseQuery(string(data))

	/* here's the map you get
	token=gIkuvaNzQIHg97ATvDxqgjtO
	team_id=T0001
	team_domain=example
	enterprise_id=E0001
	enterprise_name=Globular%20Construct%20Inc
	channel_id=C2147483705
	channel_name=test
	user_id=U2147483697
	user_name=Steve
	command=/weather
	text=94070
	response_url=https://hooks.slack.com/commands/1234/5678
	trigger_id=13345224609.738474920.8088930838d88f008e0
	*/

	getURL := "http://api.fnservice.io/r/" + m["text"][0]
	res, _ := http.Get(getURL)
	defer res.Body.Close()

	var b bytes.Buffer
	b.WriteString("Hello Chad you triggered " + m["text"][0] + " func")

	params := slack.PostMessageParameters{
		AsUser: true,
	}
	_, _, err = api.PostMessage("demostream", b.String(), params)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
