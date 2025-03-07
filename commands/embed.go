package commands

import "github.com/bwmarrin/discordgo"

func init() {
    Commands["embed"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
        embed := &discordgo.MessageEmbed{
            Title:       "Título da Embed",
            Description: "Isso é uma embed no Discord!",
            Thumbnail: &discordgo.MessageEmbedThumbnail{
              URL: m.Author.AvatarURL("1024"),
            },
            Image: &discordgo.MessageEmbedImage{
              URL: m.Author.AvatarURL("1024"),
            },
            Color: 0x00ff00, // Cor verde
        }
        s.ChannelMessageSendEmbed(m.ChannelID, embed)
    }
}