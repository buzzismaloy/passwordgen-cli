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
	help := flag.Bool("h", false, "show help")
	helpLong := flag.Bool("help", false, "show help")

	flag.Usage = func() {
		fmt.Println("passwordgen - Generate secure passwords")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  go run main.go [flags]\n  or\n  passwordgen [flags]")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  go run main.go                   # generate password with default settings")
		fmt.Println("  go run main.go --length 16 --upper --symbols  # generate password 16 chars with uppercase and symbols")
	}

	flag.Parse()

	if *help || *helpLong {
		flag.Usage()
		return
	}

	if *length < passwordgen.MinPassLength || *length > passwordgen.MaxPassLength {
		ui.PrintError(fmt.Errorf("%w; password length must be between %d and %d", passwordgen.ErrInvalidLength, passwordgen.MinPassLength, passwordgen.MaxPassLength))
		return
	}

	if !*digits && !*lower && !*upper && !*symbols {
		ui.PrintError(fmt.Errorf("%w; at least one character set must be enabled (digits, lower, upper, symbols)", passwordgen.ErrNoCharacterSet))
		return
	}

	flags := config.Flags{
		Length:    *length,
		Digits:    *digits,
		Lowercase: *lower,
		Uppercase: *upper,
		Symbols:   *symbols,
	}

	svc := service.NewPasswordService()
	pass, err := svc.Generate(flags)

	if err != nil {
		ui.PrintError(err)
		return
	}

	ui.PrintPassword(pass)
}
