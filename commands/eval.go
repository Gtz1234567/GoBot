package commands

import (
	"fmt"
	"go/parser"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Comando de eval básico
func init() {
	Commands["eval"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Verificar se o autor é um administrador (evitar uso indevido)
		if !isAdmin(m.Author.ID) {
			s.ChannelMessageSend(m.ChannelID, "Você não tem permissão para usar esse comando.")
			return
		}

		// Pega a expressão após o comando "eval "
		expression := strings.TrimPrefix(m.Content, "g+eval ")

		// Avaliar a expressão
		result, err := evaluateExpression(expression)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erro ao avaliar a expressão: "+err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Resultado: %s", result))
		}
	}
}

// Função que avalia a expressão simples
func evaluateExpression(expr string) (string, error) {
	// Tenta analisar e avaliar a expressão como uma expressão aritmética
	expr = strings.ReplaceAll(expr, " ", "") // Remove espaços em branco

	// Usar o parser do Go para avaliar a expressão aritmética
	node, err := parser.ParseExpr(expr)
	if err != nil {
		return "", fmt.Errorf("erro ao analisar a expressão: %v", err)
	}

	// Aqui podemos adicionar lógica para executar a expressão analisada
	// Como exemplo, vamos apenas retornar a expressão como string
	return fmt.Sprintf("%v", node), nil
}

// Função fictícia para verificar se o autor é administrador
func isAdmin(userID string) bool {
	// Aqui você pode usar uma lista de IDs de administradores ou fazer verificações de permissão
	return userID == "1318667195005407285" // Altere para seu ID de Discord
}