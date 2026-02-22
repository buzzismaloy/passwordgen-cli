package cmd

import (
	"flag"
	"fmt"

	"github.com/buzzismaloy/passwordgen-cli/internal/config"
	"github.com/buzzismaloy/passwordgen-cli/internal/service"
	"github.com/buzzismaloy/passwordgen-cli/internal/ui"
	"github.com/buzzismaloy/passwordgen-lib/passwordgen"
)

var generateCmd = &Command{
	Use:   "generate",
	Short: "Generate password",
	Run:   runGenerate,
}

func init() {
	RootCmd.Subcommands = append(RootCmd.Subcommands, generateCmd)
}

func runGenerate() {
	length := flag.Int("length", passwordgen.DefaultLength, "password length")
	digits := flag.Bool("digits", true, "use digits")
	lower := flag.Bool("lowercase", true, "use lowercase letters")
	upper := flag.Bool("uppercase", false, "use uppercase letters")
	symbols := flag.Bool("symbols", false, "use symbols")

	flag.Parse()

	flags := config.Flags{
		Length:    *length,
		Digits:    *digits,
		Lowercase: *lower,
		Uppercase: *upper,
		Symbols:   *symbols,
	}

	flag.Usage = func() {
		fmt.Println("Usage: passwordgen generate [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	svc := service.NewPasswordService()
	pass, err := svc.Generate(flags)

	if err != nil {
		ui.PrintError(err)
		return
	}

	ui.PrintPassword(pass)
}
