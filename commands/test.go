package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	// Comando de teste
	Commands["test"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Cria o botão
		button := discordgo.Button{
			Label:    "Botao",
			Style:    discordgo.PrimaryButton,
			CustomID: "click", // ID para identificar o botão
		}

		// Cria a linha de ações
		row := discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{&button},
		}

		// Envia a mensagem com o botão
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content:   "Olá <@" + m.Author.ID + ">",
			Components: []discordgo.MessageComponent{row},
		})
	}
}

