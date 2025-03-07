package commands

import "github.com/bwmarrin/discordgo"

func init() {
    Commands["ping"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
        s.ChannelMessageSend(m.ChannelID, "Pong!")
    }
}