package commands

import (
    "github.com/bwmarrin/discordgo"
)

// RegisterCommands registers all commands.
func GetCommands() []*discordgo.ApplicationCommand {
    return []*discordgo.ApplicationCommand{
        {
            Name:        "ping",
            Description: "Replies with Pong!",
        },
        {
            Name:        "serverinfo",
            Description: "Displays the server information",
        },
    }
}

// GetCommandsHandlers returns all commands handlers.
func GetCommandsHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
    return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
        "ping":   Ping,
        "serverinfo": ServerInfo,
    }
}
