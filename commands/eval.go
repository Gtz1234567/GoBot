package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go/parser"
	"log"
	"strings"
)

// Função que será chamada quando o comando g+eval for chamado
func init() {
Commands["eval"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora mensagens enviadas pelo próprio bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Verifica se o comando começa com "g+eval"
	if !strings.HasPrefix(m.Content, "g+eval") {
		return
	}

	// Remove o prefixo "g+eval" e espaços em branco
	code := strings.TrimSpace(strings.TrimPrefix(m.Content, "g+eval"))

	if code == "" {
		s.ChannelMessageSend(m.ChannelID, "Por favor, forneça um código para avaliar.")
		return
	}

	// Cria um novo analisador de código
	node, err := parser.ParseExpr(code)
	if err != nil {
		log.Printf("Erro ao analisar código: %v", err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Erro ao avaliar o código: %v", err))
		return
	}

	// Aqui, em vez de executar o código diretamente, você pode simplesmente verificar se o código é válido
	// O código será analisado para garantir que não haja erros de sintaxe.
	result := fmt.Sprintf("Código analisado com sucesso: %s", node)

	// Retorna o resultado da avaliação
	s.ChannelMessageSend(m.ChannelID, result)
}
}