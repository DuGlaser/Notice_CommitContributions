package main

import (
	"fmt"

	"github.com/shurcooL/githubv4"
)

func (dw *DiscordWebhook) CreateMessage(q githubv4.Int) {

	commit := fmt.Sprintf("今日のコミット数は%v回です！！", q)

	switch {
	case q <= 2:
		dw.UserName = "中野五月"
		dw.AvatarURL = "https://cdn-ak.f.st-hatena.com/images/fotolife/m/magazine_pocket/20171213/20171213201322.jpg"
		dw.Embeds = []DiscordEmbed{
			DiscordEmbed{
				Title: commit,
				Image: DiscordImage{URL: "https://media.Discordapp.net/attachments/567985071701753857/639018077589078016/S__27058352.jpg?width=1090&height=1141"},
				Color: 0xff0000,
			},
		}

	case q > 3 && q <= 5:
		dw.UserName = "中野一花"
		dw.AvatarURL = "https://cdn-ak.f.st-hatena.com/images/fotolife/m/magazine_pocket/20171213/20171213200413.jpg"
		dw.Embeds = []DiscordEmbed{
			DiscordEmbed{
				Title: commit,
				Image: DiscordImage{URL: "https://media.Discordapp.net/attachments/567985071701753857/639018284901203968/S__27058782.jpg"},
				Color: 0xffff00,
			},
		}
	case q > 5 && q <= 8:
		dw.UserName = "中野四葉"
		dw.AvatarURL = "http://chomanga.org/wp-content/uploads/2019/12/a6516e3f616a117ed66a7af940fdfed6.png"
		dw.Embeds = []DiscordEmbed{
			DiscordEmbed{
				Title: commit,
				Image: DiscordImage{URL: "https://imasoku.com/wp-content/uploads/2019/02/yZoTi8u.jpg"},
				Color: 0x008000,
			},
		}
	case q > 8 && q <= 12:
		dw.UserName = "中野三玖"
		dw.AvatarURL = "https://cdn-ak.f.st-hatena.com/images/fotolife/m/magazine_pocket/20171213/20171213200842.jpg"
		dw.Embeds = []DiscordEmbed{
			DiscordEmbed{
				Title: commit,
				Image: DiscordImage{URL: "http://phoenix-wind.com/common/img/OGP/word/gotoubun_miku_03.jpg"},
				Color: 0x0000ff,
			},
		}
	case q > 12:
		dw.UserName = "中野二乃"
		dw.AvatarURL = "https://pbs.twimg.com/media/DyehWWfVsAA6JWV.jpg"
		dw.Embeds = []DiscordEmbed{
			DiscordEmbed{
				Title: commit,
				Image: DiscordImage{URL: "https://pbs.twimg.com/media/Dgngye7U8AI9f-3?format=jpg&name=900x900"},
				Color: 0x000000,
			},
		}
	}
}
