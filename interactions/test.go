package interactions

import (
	"github.com/bwmarrin/discordgo"
)

// Função para registrar o manipulador de interações (como botões)
func RegisterInteractionHandler(s *discordgo.Session) {
	// Registra o evento de interação
	s.AddHandler(func(s *discordgo.Session, m *discordgo.InteractionCreate) {
		// Verifica se a interação é um componente de mensagem (botão clicado)
		if m.Type == discordgo.InteractionMessageComponent {
			// Lida com o clique do botão
			if m.MessageComponentData().CustomID == "click" {
				// Responde ao clique no botão
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