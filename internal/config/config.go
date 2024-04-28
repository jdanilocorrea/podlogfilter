package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ChooseKubeConfig(reader *bufio.Reader) (string, error) {
	fmt.Println("Escolher contexto de configuração:")
	fmt.Println("[1] ~/.kube/config")
	fmt.Println("[2] ~/.k3s/k3s.yaml")
	fmt.Println("[3] Outro")
	fmt.Print("Escolha uma opção: ")

	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler a entrada: %w", err)
	}

	switch strings.TrimSpace(choice) {
	case "1":
		return checkConfigPath(filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	case "2":
		return checkConfigPath(filepath.Join(os.Getenv("HOME"), ".k3s", "k3s.yaml"))
	case "3":
		fmt.Print("Digite o caminho completo do arquivo kubeconfig: ")
		path, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("erro ao ler o caminho do arquivo: %w", err)
		}
		return checkConfigPath(strings.TrimSpace(path))
	default:
		return "", fmt.Errorf("opção inválida")
	}
}

func checkConfigPath(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("arquivo de configuração não encontrado em '%s'", path)
	}
	return path, nil
}
