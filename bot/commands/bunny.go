package commands

import (
	"imgbot/bot/utils"
	"github.com/bwmarrin/discordgo"

	"math/rand"
)

func Bunny(s *discordgo.Session, m *discordgo.MessageCreate) {
	url := "https://api.bunnies.io/v2/loop/random/?media=gif,png"

	var data struct {
		Media struct {
			Poster string `json:"poster"`
			GIF    string `json:"gif"`
		} `json:"media"`
	}

	if err := utils.MakeGetRequest(url, &data); err != nil {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Failed to get image: " + err.Error())
		return
	}

	// select a gif or a poster randomly
	var selectedMedia string
	
	if (rand.Intn(2) == 0) {
		selectedMedia = data.Media.Poster
	} else {
		selectedMedia = data.Media.GIF
	}

	if (selectedMedia == "") {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Error: No image URL provided by API")
		return
	}

	embed := &discordgo.MessageEmbed{
		Color: 0xFFFFFF,  // White
		Title: "Random Bunny! :rabbit:",
		Image: &discordgo.MessageEmbedImage{URL: selectedMedia},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}