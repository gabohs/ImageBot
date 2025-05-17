package commands

import "github.com/bwmarrin/discordgo"

func Cmdlist(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed { 
		
		Title: "ImageBot Command List",
		Description: "Discord bot that sends you random images!",

		Thumbnail: &discordgo.MessageEmbedThumbnail{
        	URL: "https://cdn.discordapp.com/avatars/1373056776693481552/8c821a6292d2a59ba99b3da3e4cd62b5.webp?size=80",
    	},

		Color: 0x0000FF,

		Fields: []*discordgo.MessageEmbedField{

			&discordgo.MessageEmbedField{
				Name:   "PREFIX",
				Value:  "`!imgb`",
				Inline: false,
			},

			&discordgo.MessageEmbedField{
				Name:   "ANIMALS:",
				Inline: false,
			},

			&discordgo.MessageEmbedField{
				Name:   "Random Duck :duck:",
				Value:  "`duck`",
				Inline: true,
			},
	
			&discordgo.MessageEmbedField{
				Name:   "Random Dog :dog:",
				Value:  "`dog`",
				Inline: true,
			},

			&discordgo.MessageEmbedField{
				Name:   "Random Cat :cat:",
				Value:  "`cat`",
				Inline: true,
			},

			&discordgo.MessageEmbedField{
				Name:   "Random Fox :fox:",
				Value:  "`fox`",
				Inline: true,
			},

			&discordgo.MessageEmbedField{
				Name:   "Random Bunny :rabbit:",
				Value:  "`bunny`",
				Inline: true,
			},

			&discordgo.MessageEmbedField{
				Name:   "OTHER:",
				Inline: false,
			},
			
			&discordgo.MessageEmbedField{
				Name:   "Random Picture with selected dimensions :city_dusk:",
				Value:  "`pic <min> <max>`",
				Inline: true,
			},

			&discordgo.MessageEmbedField{
				Name:   "Random Picture with custom dimensions :city_sunset:",
				Value:  "`piccd <width> <height>`",
				Inline: true,
			},

    	},

		Footer: &discordgo.MessageEmbedFooter{
			Text : "\nDon't forget to use the prefix before the commands!",
		},

	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}