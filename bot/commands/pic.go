package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func Pic(s *discordgo.Session, m *discordgo.MessageCreate, min int, max int) {
	// generate random numbers between 150 and 800
	var imgHeight = rand.Intn(max - min) + min  
	var imgWidth = rand.Intn(max - min) + min

	url := fmt.Sprintf("https://picsum.photos/%d/%d", imgWidth, imgHeight)

	embed := &discordgo.MessageEmbed { 
		Color:       0x808080, // Gray
		Title:       "Random Image! :city_dusk:",
		Image:       &discordgo.MessageEmbedImage{URL: url},

		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("(%d x %d)", imgWidth, imgHeight),
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}