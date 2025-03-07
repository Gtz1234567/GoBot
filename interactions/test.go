package interactions

import (
	"github.com/bwmarrin/discordgo"
)

var messageAuthorID = make(map[string]string) // Mapa para armazenar o autor da mensagem

func RegisterInteractionHandler(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.InteractionCreate) {
		if m.Type == discordgo.InteractionMessageComponent {
			customID := m.MessageComponentData().CustomID

			if customID == "click" {
				// Verifica se o ID da mensagem está no mapa
				authorID, exists := messageAuthorID[m.Message.ID]

				if !exists || authorID != m.Member.User.ID {
					// Se não for o autor da mensagem, responde que a interação não é para ele
					s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Essa interação não é para você!",
							Flags:   discordgo.MessageFlagsEphemeral, // Mensagem visível apenas para o usuário que tentou interagir
						},
					})
					return
				}

				// Se for o autor, responde normalmente
				s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Você clicou no botão, " + m.Member.User.Username,
					},
				})
			}
		}
	})
}

// Função para registrar o autor de uma mensagem
func RegisterMessageAuthor(messageID, authorID string) {
	messageAuthorID[messageID] = authorID
}