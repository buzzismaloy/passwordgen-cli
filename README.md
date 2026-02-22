# passwordgen-cli

`passwordgen-cli` is a lightweight command-line interface (CLI) tool for generating secure passwords.
It is a thin wrapper around the [`passwordgen-lib`](https://github.com/buzzismaloy/passwordgen-lib) Go library, exposing its functionality through a simple and configurable CLI interface.


![Go](https://img.shields.io/badge/Go-1.25-blue)
![Repo Size](https://img.shields.io/github/repo-size/buzzismaloy/passwordgen-cli)

---

## Features

- Generate cryptographically secure passwords
- Configurable password length
- Enable or disable:
  - digits
  - lowercase letters
  - uppercase letters
  - symbols
- Built with a minimal custom command framework (no Cobra or external CLI frameworks)

---

## Project Structure

```
.
├── main.go
├── cmd/
│ ├── root.go
│ ├── command.go
│ └── generate.go
├── internal/
│ ├── config/
│ │ └── flags.go
│ ├── service/
│ │ └── generator.go
│ └── ui/
│   └── printer.go
├── go.mod
```

---

## Usage

Clone the repository:

```bash
git clone https://github.com/buzzismaloy/passwordgen-cli.git
```

And then run the generator:

```bash
go run main.go [flags]
```

Or just build it into bin:

```bash
go build -o passwordgen

./passwordgen [flags]
```

---

## Flags


| Flag | Description | Default |
|------|-------------|----------|
| --length | Password length | passwordgen.DefaultLength |
| --digits | Include digits | true |
| --lowercase | Include lowercase letters | true |
| --uppercase | Include uppercase letters | false |
| --symbols | Include symbols | false |
| -h, --help. | Show help message | - |

---

## Examples

Generate a password with default settings:

```bash
go run main.go
```

Generate a 16-character password with uppercase letters and symbols:

```bash
go run main.go --length 16 --uppercase --symbols
```

or

```bash
go run main.go --length 16 --uppercase=true --symbols=true
```

Disable digits and use only lowercase letters:

```bash
go run main.go generate --digits=false --uppercase=false --symbols=false
```

---

## How it works

1. The CLI parses flags using Go's `flag` package.
2. Flags are mapped into an internal `config.Flags` struct.
3. `PasswordService` converts CLI flags into a `passwordgen-lib` configuration.
4. `passwordgen-lib` generates the password.
5. The CLI prints the result to stdout.

--- 

## Validation Rules

- Password length must be between `passwordgen.MinPassLength and passwordgen.MaxPassLength`.
- At least one character set must be enabled (digits, lowercase, uppercase, symbols).

If validation fails, an error is printed and the program exits.

---

## Dependency

This project depends on:

```
github.com/buzzismaloy/passwordgen-lib
```

In go.mod, the library is replaced locally:

```
replace github.com/buzzismaloy/passwordgen-lib => ../passwordgen-lib
```

---

## Notes

- This project intentionally avoids heavy CLI frameworks to keep the code minimal and educational.
- The command system is custom-built via the `Command` struct and recursive execution.
- Output is handled via a simple UI layer (`internal/ui`).

---

## License

[![License](https://img.shields.io/github/license/buzzismaloy/passwordgen-cli)](https://github.com/buzzismaloy/passwordgen-cli/blob/main/LICENSE)

