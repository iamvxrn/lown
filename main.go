package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"lown/internal/core"
	"lown/internal/path"
	"lown/internal/ui"
)

const version = "0.3.0"

func printUsage() {
	banner := `Lown - Decentralized, User-Space Package Manager & Orchestrator

Usage:
  lown <command> [arguments]

Commands:
  install, i <uri>    Install package from Git URL, gh:user/repo, alias, or local path
  remove, rm <name>   Remove an installed package and its linked binary
  rollback <name>     Restore previous backup executable (.bak) for package
  sync, update [name] Pull latest changes and rebuild package(s) if updated
  list, ls            List all installed packages
  doctor              Check Lown installation status and shell environment
  completion <shell>  Generate shell auto-completion script (bash, zsh, fish)
  version             Print Lown version
  help                Show this help message

Philosophy:
  • No sudo (100% User-Space in ~/.lown/)
  • No central registry (Decentralized Git repositories)
  • Transparency (Native builds or explicit install scripts)
`
	fmt.Print(banner)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	subcommand := strings.ToLower(os.Args[1])

	switch subcommand {
	case "install", "i":
		if len(os.Args) < 3 {
			ui.LogError("Missing URI or package specification.")
			fmt.Println("Usage: lown install <uri>")
			os.Exit(1)
		}
		rawURI := os.Args[2]
		installer, err := core.NewInstaller()
		if err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}
		if err := installer.Install(rawURI); err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}

	case "remove", "rm", "uninstall":
		if len(os.Args) < 3 {
			ui.LogError("Missing package name.")
			fmt.Println("Usage: lown remove <package-name>")
			os.Exit(1)
		}
		pkgName := os.Args[2]
		installer, err := core.NewInstaller()
		if err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}
		if err := installer.Remove(pkgName); err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}

	case "rollback":
		if len(os.Args) < 3 {
			ui.LogError("Missing package name.")
			fmt.Println("Usage: lown rollback <package-name>")
			os.Exit(1)
		}
		pkgName := os.Args[2]
		installer, err := core.NewInstaller()
		if err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}
		if err := installer.Rollback(pkgName); err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}

	case "sync", "update":
		pkgName := ""
		if len(os.Args) >= 3 {
			pkgName = os.Args[2]
		}
		installer, err := core.NewInstaller()
		if err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}
		if err := installer.Sync(pkgName); err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}

	case "list", "ls":
		installer, err := core.NewInstaller()
		if err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}
		if err := installer.List(); err != nil {
			ui.LogError("%v", err)
			os.Exit(1)
		}

	case "doctor":
		runDoctor()

	case "completion":
		shell := "zsh"
		if len(os.Args) >= 3 {
			shell = os.Args[2]
		}
		printCompletion(shell)

	case "version", "-v", "--version":
		fmt.Printf("Lown v%s\n", version)

	case "help", "-h", "--help":
		printUsage()

	default:
		ui.LogError("Unknown command '%s'", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func runDoctor() {
	ui.LogInfo("Running Lown system diagnostics...")
	fmt.Printf("  • Lown Version: v%s\n", version)
	fmt.Printf("  • Lown Root:    %s\n", path.LownRoot())
	fmt.Printf("  • Bin Dir:      %s\n", path.BinDir())
	fmt.Printf("  • Apps Dir:     %s\n", path.AppsDir())
	fmt.Printf("  • Config File:  %s\n", path.ConfigPath())

	binDir := path.BinDir()
	pathEnv := os.Getenv("PATH")
	inPath := false
	for _, p := range filepath.SplitList(pathEnv) {
		if filepath.Clean(p) == filepath.Clean(binDir) {
			inPath = true
			break
		}
	}

	if inPath {
		ui.LogSuccess("PATH configuration is correct (%s is in PATH).", binDir)
	} else {
		ui.LogWarning("PATH configuration missing. '%s' is NOT in $PATH.", binDir)
		fmt.Printf("    To fix, add this line to your ~/.bashrc or ~/.zshrc:\n")
		fmt.Printf("    export PATH=\"%s:$PATH\"\n", binDir)
	}
}

func printCompletion(shell string) {
	switch strings.ToLower(shell) {
	case "bash":
		fmt.Print(`# lown bash completion
_lown_completions() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    local cmds="install remove rollback sync list doctor version help completion"
    if [ $COMP_CWORD -eq 1 ]; then
        COMPREPLY=( $(compgen -W "${cmds}" -- ${cur}) )
    fi
}
complete -F _lown_completions lown
`)
	case "zsh":
		fmt.Print(`#compdef lown
_lown() {
    local -a commands
    commands=(
        'install:Install package from Git URL or spec'
        'remove:Remove installed package'
        'rollback:Restore previous backup executable'
        'sync:Pull latest changes and rebuild package'
        'list:List all installed packages'
        'doctor:Check Lown installation status'
        'version:Print Lown version'
        'completion:Generate shell auto-completion script'
    )
    _describe 'command' commands
}
_lown "$@"
`)
	case "fish":
		fmt.Print(`# lown fish completion
complete -c lown -f
complete -c lown -n "not __fish_seen_subcommand_from install remove rollback sync list doctor version completion" -a "install remove rollback sync list doctor version completion"
`)
	default:
		ui.LogError("Unsupported shell '%s'. Supported shells: bash, zsh, fish", shell)
		os.Exit(1)
	}
}

