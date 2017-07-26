package main

import (
	"errors"
	"log"

	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

const (
	readFromStdin = "-"
)

type llessCmd struct {
	BG          string
	Color       string
	HTML        bool
	ShowPalette bool
	ShowVersion bool
}

func (c *llessCmd) Run(cmd *cobra.Command, args []string) {
	stdout := colorable.NewColorableStdout()

	if c.ShowVersion {
		displayVersion(stdout)
		return
	}

	var colorPalettes ColorPalettes
	if c.BG == "dark" {
		colorPalettes = DarkColorPalettes
	} else {
		colorPalettes = LightColorPalettes
	}

	if len(args) < 1 {
		err := errors.New("Have to specify at least one filename")
		log.Fatal(err)

		return
	}
	displayLoop(args[0], colorPalettes)
}

func main() {
	log.SetFlags(0)
	llessCmd := &llessCmd{}
	rootCmd := &cobra.Command{
		Use:  "lless [OPTION]... [FILE]...",
		Long: "Colorize FILE(s), or standard input, to standard output.",
		Example: `$ lless FILE1 FILE2 ...
  $ lless --bg=dark FILE1 FILE2 ... # dark background
  $ lless --html # output html
  $ lless -G String="_darkblue_" -G Plaintext="darkred" FILE # set color codes
  $ lless # read from standard input
  $ curl https://raw.githubusercontent.com/jingweno/lless/master/main.go | lless`,
		Run: llessCmd.Run,
	}

	usageTempl := `{{ $cmd := . }}
Usage:
  {{.UseLine}}

Flags:
{{.LocalFlags.FlagUsages}}
Using color is auto both by default and with --color=auto. With --color=auto,
lless emits color codes only when standard output is connected to a terminal.
Color codes can be changed with -G KEY=VALUE. List of color codes can
be found with --palette.

Examples:
  {{ .Example }}
`
	rootCmd.SetUsageTemplate(usageTempl)

	rootCmd.PersistentFlags().StringVarP(&llessCmd.BG, "bg", "", "light", `set to "light" or "dark" depending on the terminal's background`)
	rootCmd.PersistentFlags().StringVarP(&llessCmd.Color, "color", "C", "auto", `colorize the output; value can be "never", "always" or "auto"`)
	// rootCmd.PersistentFlags().VarP(&llessCmd.ColorCodes, "color-code", "G", `set color codes`)
	rootCmd.PersistentFlags().BoolVarP(&llessCmd.HTML, "html", "", false, `output html`)
	rootCmd.PersistentFlags().BoolVarP(&llessCmd.ShowPalette, "palette", "", false, `show color palettes`)
	rootCmd.PersistentFlags().BoolVarP(&llessCmd.ShowVersion, "version", "v", false, `show version`)

	rootCmd.Execute()
}
