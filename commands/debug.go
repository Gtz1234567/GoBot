package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func init() {
    Commands["debug"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
        log.Printf("Mensagem recebida: %s", m.Content)

        if len(m.Mentions) == 0 {
            log.Println("Nenhuma menção detectada.")
            s.ChannelMessageSend(m.ChannelID, "Nenhuma menção detectada.")
        } else {
            log.Printf("Menções detectadas: %d", len(m.Mentions))
            for _, mention := range m.Mentions {
                log.Printf("Usuário mencionado: %s (ID: %s)", mention.Username, mention.ID)
                s.ChannelMessageSend(m.ChannelID, "Mencionado: "+mention.Username)
            }
        }
    }
}