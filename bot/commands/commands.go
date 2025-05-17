package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"imgbot/bot/utils"
)

const prefix = "!imgb"

func ProcessCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")
	if (args[0] != prefix) {
		return
	}

	switch args[1] {
		case "hello":
			Hello(s, m)
		case "cmdlist":
			Cmdlist(s, m)
		
		// commands to get random images:

		// ANIMALS
		case "duck":
			Duck(s, m)
		case "dog":
			Dog(s, m)
		case "cat":
			Cat(s, m)
		case "fox":
			Fox(s, m)
		case "bunny":
			Bunny(s, m)

		// OTHER
		case "pic":
			min, max, ok := utils.ParseTwoArgsAsInt(args, s, m)
			if (ok) {
				Pic(s, m, min, max)
			}

		case "piccd":
			width, height, ok := utils.ParseTwoArgsAsInt(args, s, m)
			if (ok) {
				Piccd(s, m, width, height)
			}
	}
}