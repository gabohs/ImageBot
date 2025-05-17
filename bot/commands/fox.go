package commands

import (
	"github.com/bwmarrin/discordgo"
	"imgbot/bot/utils"
)

func Fox(s *discordgo.Session, m *discordgo.MessageCreate) {
	url := "https://randomfox.ca/floof/"

	var data struct {
		Image string `json:"image"`
	}

	if err := utils.MakeGetRequest(url, &data); err != nil {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Failed to get image: " + err.Error())
		return
	}

	if data.Image == "" {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Error: No image URL provided by API")
		return
	}

	embed := &discordgo.MessageEmbed { 
		Color:       0xFFA500, // Orange
		Title:       "Random Fox! :fox:",
		Image:       &discordgo.MessageEmbedImage{URL: data.Image},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}