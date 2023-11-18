package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"regexp"
)

func main() {
	var rootCmd = &cobra.Command{Use: "huncodingCli"}

	var nome, email, senha string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Crie um novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if nome == "" {
				fmt.Println("Nome não pode estar vazio.")
				return
			}

			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
			if !emailRegex.MatchString(email) {
				fmt.Println("Email inválido. Por favor, insira um email válido.")
				return
			}

			if len(senha) < 6 {
				fmt.Println("A senha deve ter pelo menos 6 caracteres.")
				return
			}

			fmt.Printf("Nome: %s\nEmail: %s\nSenha: %s\n", nome, email, senha)
		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuário")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Endereço de email")
	cmd.Flags().StringVarP(&senha, "senha", "s", "", "Senha do usuário")

	rootCmd.AddCommand(cmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Não foi possível executar!")
	}
}
