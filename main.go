package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type discordEmbed struct {
	Desc  string     `json:"description"`
	URL   string     `json:"url"`
	Color int        `json:"color"`
	Image discordImg `json:"image"`
}

type discordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []discordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}

func getTime() (time.Time, time.Time) {
	t := time.Now()
	y := time.Now().AddDate(0, 0, -1)
	return t, y
}

func sendMessage(url string, dw *discordWebhook) {
	j, err := json.Marshal(dw)
	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
}

func main() {
	today, yesterday := getTime()

	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var q struct {
		User struct {
			ContrbutionsCollection struct {
				TotalCommitContributions githubv4.Int
			} `graphql:"contributionsCollection(from: $Yesterday,to: $Today)"`
		} `graphql:"user(login: \"DuGlaser\")"`
	}

	variable := map[string]interface{}{
		"Today":     githubv4.DateTime{today},
		"Yesterday": githubv4.DateTime{yesterday},
	}

	err := client.Query(context.Background(), &q, variable)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(q.User.ContrbutionsCollection.TotalCommitContributions)
}
