package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
    Commands["avatar"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
        var userID string
        if len(m.Mentions) > 0 {
            userID = m.Mentions[0].ID // Se alguém for mencionado, pega o ID da pessoa
        } else {
            userID = m.Author.ID // Se não houver menção, usa o ID do autor da mensagem
        }

        user, err := s.User(userID)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Não consegui encontrar o usuário!")
            return
        }

        // URL do avatar
        avatarURL := user.AvatarURL("1024")

        // Criando o botão de link
        button := discordgo.Button{
            Label: "Baixar Avatar",      // Texto do botão
            Style: discordgo.LinkButton, // Botão de link
            URL:   avatarURL,            // URL do avatar
        }

        // Criando a embed
        embed := &discordgo.MessageEmbed{
            Title:       user.Username + " Avatar",
            Description: "Aqui está o avatar de " + user.Username,
            Image: &discordgo.MessageEmbedImage{
                URL: avatarURL, // Exibe o avatar
            },
            Color: 16711936, // Cor verde
        }

        // Enviando a embed com o botão de link
        s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
            Embed: embed,
            Components: []discordgo.MessageComponent{
                discordgo.ActionsRow{
                    Components: []discordgo.MessageComponent{button},
                },
            },
        })
    }
}