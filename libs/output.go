package libs

import (
	"fmt"
	"io"
	"strings"
)

func ResolveOutputFormat(requested string, fallback string) (string, error) {
	format := strings.TrimSpace(requested)
	if format == "" {
		format = strings.TrimSpace(fallback)
	}
	if format == "" {
		format = "table"
	}

	switch format {
	case "table", "text":
		return format, nil
	default:
		return "", fmt.Errorf("unsupported output format %q, use one of: table, text", format)
	}
}

func RenderText(writer io.Writer, text string) error {
	_, err := fmt.Fprintln(writer, text)
	return err
}

func RenderTable(writer io.Writer, headers []string, rows [][]string) error {
	if len(headers) == 0 {
		return fmt.Errorf("table headers are required")
	}

	widths := make([]int, len(headers))
	for index, header := range headers {
		widths[index] = len(header)
	}

	for _, row := range rows {
		for index, cell := range row {
			if index >= len(widths) {
				break
			}
			if len(cell) > widths[index] {
				widths[index] = len(cell)
			}
		}
	}

	separator := buildSeparator(widths)
	if _, err := fmt.Fprintln(writer, separator); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(writer, buildRow(headers, widths)); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(writer, separator); err != nil {
		return err
	}
	for _, row := range rows {
		if _, err := fmt.Fprintln(writer, buildRow(row, widths)); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintln(writer, separator)
	return err
}

func buildSeparator(widths []int) string {
	parts := make([]string, 0, len(widths))
	for _, width := range widths {
		parts = append(parts, strings.Repeat("-", width+2))
	}
	return "+" + strings.Join(parts, "+") + "+"
}

func buildRow(values []string, widths []int) string {
	columns := make([]string, 0, len(widths))
	for index, width := range widths {
		value := ""
		if index < len(values) {
			value = values[index]
		}
		columns = append(columns, fmt.Sprintf(" %-*s ", width, value))
	}
	return "|" + strings.Join(columns, "|") + "|"
}