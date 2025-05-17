package bot

import (
	"github.com/bwmarrin/discordgo"
	"imgbot/bot/commands"
)

func Start(token string) error {
	sess, err := discordgo.New("Bot " + token)
	if (err != nil) {
		return err
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		commands.ProcessCommands(s, m)
	})

	return sess.Open()
}