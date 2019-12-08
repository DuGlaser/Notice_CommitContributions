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

type DiscordImage struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}

type DiscordEmbed struct {
	Desc  string       `json:"description"`
	Image DiscordImage `json:"image"`
	Color int          `json:"color"`
	Title string       `json:"title"`
}

type DiscordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []DiscordEmbed `json:"embeds"`
}

func getTime() (time.Time, time.Time) {
	t := time.Now()
	y := time.Now().AddDate(0, 0, -1)
	return t, y
}

func sendMessage(url string, dw *DiscordWebhook) {
	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	if resp.StatusCode == 204 {
		fmt.Println("sent", dw)
	} else {
		fmt.Println("失敗")
		fmt.Printf("%#v\n", resp)
	}
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

	dw := &DiscordWebhook{}
	dw.CreateMessage(q.User.ContrbutionsCollection.TotalCommitContributions)
	sendMessage(os.Getenv("DISCORD_WEBHOOK_TEST2"), dw)
}
