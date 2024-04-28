package logs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// DisplayLogs exibe os logs no formato de tabela.
// DisplayLogs exibe os logs no formato de tabela.
func DisplayLogs(logEntries []LogEntry) {
	// Cria um tabwriter com colunas alinhadas à esquerda e um padding adequado.
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0) // Removido AlignRight|tabwriter.Debug para o layout padrão
	defer writer.Flush()

	// Imprime o cabeçalho da tabela
	fmt.Fprintln(writer, "#\tCreatedAt\tStatus\tMethod\tPath\t")

	for i, entry := range logEntries {
		// Imprime a linha da tabela formatada com cada campo em sua respectiva coluna
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\t\n", i+1, entry.CreatedAt, strconv.Itoa(entry.Status), entry.Method, entry.Path)
		// Detalhes são impressos em uma linha separada abaixo da entrada
		fmt.Fprintf(writer, "Details: %s\n", entry.Details)
		// Insere uma linha separadora após os detalhes de cada entrada
		fmt.Fprintln(writer, strings.Repeat("-", 80))
	}
	writer.Flush() // Garante que todos os dados sejam enviados para o os.Stdout antes de retornar
}
