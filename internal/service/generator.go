package service

import (
	"github.com/buzzismaloy/passwordgen-cli/internal/config"
	"github.com/buzzismaloy/passwordgen-lib/passwordgen"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) Generate(f config.Flags) (string, error) {
	cfg := passwordgen.NewConfig(
		passwordgen.WithLength(f.Length),
		passwordgen.WithDigits(f.Digits),
		passwordgen.WithLowercase(f.Lowercase),
		passwordgen.WithUppercase(f.Uppercase),
		passwordgen.WithSymbols(f.Symbols),
	)

	gen, err := passwordgen.NewGenerator(*cfg, nil)
	if err != nil {
		return "", err
	}

	return gen.Generate()
}
