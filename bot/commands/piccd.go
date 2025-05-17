package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Piccd(s *discordgo.Session, m *discordgo.MessageCreate, width int, height int) {
	url := fmt.Sprintf("https://picsum.photos/%d/%d", width, height)

	embed := &discordgo.MessageEmbed { 
		Color:       0x808080, // Gray
		Title:       "Random Image! :city_sunset:",
		Image:       &discordgo.MessageEmbedImage{URL: url},

		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("(%d x %d)", width, height),
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}