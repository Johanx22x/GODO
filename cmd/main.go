package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "log"

    "github.com/bwmarrin/discordgo"
    "github.com/Johanx22x/GODO"
    // "github.com/Johanx22x/GODO/internal/handler"
    "github.com/Johanx22x/GODO/internal/commands"
)

func main() {
    config := godo.GetConfig()

    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot " + config.Token)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }

    // Manage the commands handlers.
    commandsHandlers := commands.GetCommandsHandlers()
    dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        if handler, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
            handler(s, i)
        }
    })

    // Add intents. In this case, dm intents are required.
    dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

    dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
        log.Printf("Logged in as %s#%s", r.User.Username, r.User.Discriminator)
    })

    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    // Register the commands.
    commands := commands.GetCommands()
    registeredCommands := (make([]*discordgo.ApplicationCommand, len(commands)))

    for i, command := range commands {
        cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, config.GuildID, command)
        if err != nil {
            log.Panicf("Cannot register '%s' command: %s", command.Name, err)
        }
        registeredCommands[i] = cmd 
    }

    // Wait here until CTRL-C or other term signal is received.
    fmt.Println("GODO " + config.Version + " is now running.  Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Delete the commands.
    for _, command := range registeredCommands {
        err := dg.ApplicationCommandDelete(dg.State.User.ID, config.GuildID, command.ID)
        if err != nil {
            log.Panicf("Cannot delete '%s' command: %s", command.Name, err)
        }
    }

    // Cleanly close down the Discord session.
    log.Println("Gracefully shutting down...")
    dg.Close()
}
