package commands

import (
	"github.com/bwmarrin/discordgo"
	"imgbot/bot/utils"
)

func Dog(s *discordgo.Session, m *discordgo.MessageCreate) {
	url := "https://random.dog/woof.json"

	var data struct {
		URL string `json:"url"`
	}

	if err := utils.MakeGetRequest(url, &data); err != nil {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Failed to get image: " + err.Error())
		return
	}

	if data.URL == "" {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Error: No image URL provided by API")
		return
	}

	embed := &discordgo.MessageEmbed{
		Color: 0x00FFCC,  // Aqua
		Title: "Random Dog! :dog:",
		Image: &discordgo.MessageEmbedImage{URL: data.URL},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}