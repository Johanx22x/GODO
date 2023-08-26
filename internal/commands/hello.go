package commands

import (
    "github.com/bwmarrin/discordgo"
)

func Hello(s *discordgo.Session, m *discordgo.MessageCreate) {
    username := m.Author.Username
    s.ChannelMessageSend(m.ChannelID, "Hello, " + username + "!")
}
