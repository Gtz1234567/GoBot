package commands

import "github.com/bwmarrin/discordgo"

// Mapa global de comandos
var Commands = map[string]func(*discordgo.Session, *discordgo.MessageCreate){}