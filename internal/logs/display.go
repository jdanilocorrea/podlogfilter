package logs

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
)

type LogEntry struct {
	CreatedAt  time.Time
	StatusCode int
	Method     string
	Path       string
	Details    string
}

func DisplayLogs(logEntries []LogEntry) {
	writer := new(tabwriter.Writer)

	// Minwidth, tabwidth, padding, padchar, flags
	writer.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer writer.Flush()

	// Cabeçalhos com cores
	headerFmt := color.New(color.FgWhite, color.Bold).SprintfFunc()
	fmt.Fprintln(writer, headerFmt("Data\tStatus Code\tMethod\tPath\tLog Details\t"))

	// Itera por cada entrada de log para imprimir na tabela
	for _, entry := range logEntries {
		statusColor := color.New(color.FgGreen)
		if entry.StatusCode >= 400 {
			statusColor = color.New(color.FgRed)
		}

		date := entry.CreatedAt.Format("2006-01-02 15:04:05")
		statusCode := statusColor.Sprintf("%d", entry.StatusCode)
		method := entry.Method
		path := entry.Path
		details := entry.Details

		// Imprime a linha formatada
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t\n",
			date, statusCode, method, path, details)
	}

	// Desenha a linha no final
	fmt.Fprintln(writer, strings.Repeat("-", 80))
}
