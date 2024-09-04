package events

import (
	"sot/src/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "t!"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, prefix) {
		command := strings.TrimPrefix(m.Content, prefix)

		if strings.HasPrefix(strings.Trim(command, " "), "procurar-membros") {
			commands.SearchMembers(s, m)
		}
	}
}
