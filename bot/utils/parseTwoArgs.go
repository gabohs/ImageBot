package utils

import (
	"strconv"
	"github.com/bwmarrin/discordgo"
)

func ParseTwoArgsAsInt(args []string, s *discordgo.Session, m *discordgo.MessageCreate) (int, int, bool) {
	
	arg1, err1 := strconv.Atoi(args[2]) 
	arg2, err2 := strconv.Atoi(args[3])

	if err1 != nil {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Error: Argument 1 is invalid")
		return 0, 0, false
	}
	if err2 != nil {
		s.ChannelMessageSend(m.ChannelID, "# :warning: Error: Argument 2 is invalid")
		return 0, 0, false
	}
	return arg1, arg2, true
}