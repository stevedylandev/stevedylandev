package main

import "fmt"

const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"
	KeyColor = Yellow
	PaletteSwatch = "●"
)

type Row struct {
	Key   string
	Value string
}

type Color struct {
	Name string
	Code string
}

var Palette = []Color{
	{"Gray", Gray},
	{"Red", Red},
	{"Green", Green},
	{"Yellow", Yellow},
	{"Blue", Blue},
	{"Magenta", Magenta},
	{"Cyan", Cyan},
	{"White", White},
}

func Colorize(color, s string) string {
	return color + s + Reset
}

func PrintPalette() {
	fmt.Print("  ")
	for _, c := range Palette {
		swatch := Colorize(c.Code, PaletteSwatch + "  ")
		fmt.Print(swatch)
	}
	fmt.Println()
}

func PrintRows(rows []Row) {
	width := labelWidth(rows)

	fmt.Println()
	for _, row := range rows {
		if row.Value == "Unknown" {
			continue
		}
		label := Colorize(Bold+KeyColor, fmt.Sprintf("%-*s", width, row.Key))
		line := "  " + label + "    " + row.Value

		fmt.Println(line)
	}
	fmt.Println()
}

func labelWidth(rows []Row) int {
	width := 0
	for _, row := range rows {
		if w := len(row.Key) + 1; w > width {
			width = w
		}
	}
	return width
}

func main() {
	rows := []Row {
		Row{Key: "name", Value: "Steve Simkins"},
		Row{Key: "title", Value: "Senior Solutions Engineer"},
		Row{Key: "job", Value: "https://stablecore.com"},
		Row{Key: "website", Value: "https://stevedylan.dev"},
		Row{Key: "dream", Value: "Enable freedom in software"},
		Row{Key: "interests", Value: "art, literature, science, nature"},
		Row{Key: "projects", Value: "gargantua.page, andromeda.build"},
	}
	PrintRows(rows)
	PrintPalette()
}
