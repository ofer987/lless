package main

import (
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

	var s *Stream
	var err error
	if len(args) == 0 {
		s, err = NewStdinStream(colorPalettes)
	} else {
		s, err = NewFileStream(args[0], colorPalettes)
	}
	if err != nil {
		log.Fatal(err)
	}

	t := s.Read()
	s.Close()

	displayText(t)
}

func main() {
	log.SetFlags(0)
	llessCmd := &llessCmd{}
	rootCmd := &cobra.Command{
		Use:  "lless [OPTION]... [FILE]...",
		Long: "Colorize FILE(s), or standard input, to standard output.",
		Example: `$ lless FILE ...
  $ lless --bg=dark FILE ... # dark background
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
	rootCmd.PersistentFlags().BoolVarP(&llessCmd.ShowVersion, "version", "v", false, `show version`)

	rootCmd.Execute()
}
