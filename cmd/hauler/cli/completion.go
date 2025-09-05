package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"freighter.dev/go/freighter/internal/flags"
)

func addCompletion(parent *cobra.Command, ro *flags.CliRootOpts) {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate auto-completion scripts for various shells",
	}

	cmd.AddCommand(
		addCompletionZsh(ro),
		addCompletionBash(ro),
		addCompletionFish(ro),
		addCompletionPowershell(ro),
	)

	parent.AddCommand(cmd)
}

func addCompletionZsh(ro *flags.CliRootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generates auto-completion scripts for zsh",
		Example: `To load completion run

	. <(freighter completion zsh)

	To configure your zsh shell to load completions for each session add to your zshrc

	# ~/.zshrc or ~/.profile
	command -v freighter >/dev/null && . <(freighter completion zsh)

	or write a cached file in one of the completion directories in your ${fpath}:

	echo "${fpath// /\n}" | grep -i completion
	freighter completion zsh > _freighter

	mv _freighter ~/.oh-my-zsh/completions  # oh-my-zsh
	mv _freighter ~/.zprezto/modules/completion/external/src/  # zprezto`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenZshCompletion(os.Stdout)
			// Cobra doesn't source zsh completion file, explicitly doing it here
			fmt.Println("compdef _freighter freighter")
		},
	}
	return cmd
}

func addCompletionBash(ro *flags.CliRootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generates auto-completion scripts for bash",
		Example: `To load completion run

	. <(freighter completion bash)

	To configure your bash shell to load completions for each session add to your bashrc

	# ~/.bashrc or ~/.profile
	command -v freighter >/dev/null && . <(freighter completion bash)`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenBashCompletion(os.Stdout)
		},
	}
	return cmd
}

func addCompletionFish(ro *flags.CliRootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fish",
		Short: "Generates auto-completion scripts for fish",
		Example: `To configure your fish shell to load completions for each session write this script to your completions dir:

	freighter completion fish > ~/.config/fish/completions/freighter.fish

	See http://fishshell.com/docs/current/index.html#completion-own for more details`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenFishCompletion(os.Stdout, true)
		},
	}
	return cmd
}

func addCompletionPowershell(ro *flags.CliRootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powershell",
		Short: "Generates auto-completion scripts for powershell",
		Example: `To load completion run

	. <(freighter completion powershell)

	To configure your powershell shell to load completions for each session add to your powershell profile

	Windows:

	cd "$env:USERPROFILE\Documents\WindowsPowerShell\Modules"
	freighter completion powershell >> freighter-completion.ps1

	Linux:

	cd "${XDG_CONFIG_HOME:-"$HOME/.config/"}/powershell/modules"
	freighter completion powershell >> freighter-completions.ps1`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		},
	}
	return cmd
}
