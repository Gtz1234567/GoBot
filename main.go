package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
	"Bot/commands" // Certifique-se de que o nome do módulo no go.mod seja "Bot"
	"Bot/interactions"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	// Obter o token do Discord a partir da variável de ambiente
	Token := os.Getenv("token") // Aqui você já usa corretamente o Getenv
	if Token == "" {
		log.Fatal("Token do Discord não encontrado")
	}

	// Cria uma nova sessão do Discord
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Erro ao criar o bot:", err)
		return
	}

	// Carregar comandos automaticamente
	loadCommands()

	// Registrar o manipulador de mensagens
	dg.AddHandler(messageCreate)

	// Registrar o manipulador de interações (botões, etc.)
	interactions.RegisterInteractionHandler(dg)

	// Abrir conexão com o Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Erro ao abrir a conexão:", err)
		return
	}

	fmt.Println("Bot está funcionando. Pressione CTRL+C para sair.")
	select {}
}

// Carregar comandos automaticamente
func loadCommands() {
	commandDir := "./commands"

	err := filepath.Walk(commandDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Se for um arquivo Go válido (exceto register.go), apenas importe indiretamente
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && info.Name() != "register.go" {
			// O init() de cada comando já registra no mapa commands.Commands
			fmt.Println("Comando carregado:", strings.TrimSuffix(info.Name(), ".go"))
		}
		return nil
	})

	if err != nil {
		log.Println("Erro ao carregar comandos:", err)
	}
}

// Processar mensagens e executar comandos
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Verificar se a mensagem começa com "g+"
	if strings.HasPrefix(strings.ToLower(m.Content), "g+") {
		command := strings.TrimPrefix(strings.ToLower(m.Content), "g+")

		// Se o comando existir, execute
		if cmd, exists := commands.Commands[command]; exists {
			cmd(s, m)
		}
	}
}