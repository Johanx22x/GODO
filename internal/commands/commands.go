package commands

import (
    "strconv"

    "github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "Pong!",
        },
    })
}

func ServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
    guild, err := s.Guild(i.GuildID)
    if err != nil {
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "Cannot get the server information!",
            },
        })
        return
    }

    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "This server is called " + guild.Name + " and has " + strconv.Itoa(guild.MemberCount) + " members!",
        },
    })
}
