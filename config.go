package godo

import (
    "os"
)

type Config struct {
    Version string
    Token string
    GuildID string
}

// getConfig returns the configuration for the bot.
func GetConfig() Config {
    config := Config{}

    config.Version = "0.1.0"

    // Get the token from the environment.
    config.Token = os.Getenv("DISCORD_TOKEN")
    
    // global guild id 
    config.GuildID = ""


    return config
}
    
