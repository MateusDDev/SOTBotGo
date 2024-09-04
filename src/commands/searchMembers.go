package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func SearchMembers(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Fields(m.Content)
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Por favor, forneça o nome do cargo")
		return
	}

	roleName := args[2]
	guildID := m.GuildID
	allRoles, err := s.GuildRoles(guildID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Estamos com problemas para encontrar os cargos no momento")
		return
	}

	var roleID string
	for _, role := range allRoles {
		if role.Name == roleName {
			roleID = role.ID
			break
		}
	}

	if roleID == "" {
		s.ChannelMessageSend(m.ChannelID, "Cargo não encontrado: "+roleName)
		return
	}

	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Estamos com problemas para encontrar os membros no momento")
		return
	}

	var membersByRole []string
	for _, member := range members {
		for _, r := range member.Roles {
			if r == roleID {
				membersByRole = append(membersByRole, member.User.Username)
				break
			}
		}
	}

	if len(membersByRole) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Nenhum membro possui conhecimento em "+roleName+" no momento")
	} else {
		response := fmt.Sprintf("Membros com conhecimento em %s:\n", roleName)
		for _, username := range membersByRole {
			response += "- " + username + "\n"
		}
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
