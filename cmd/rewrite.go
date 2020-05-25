package cmd

import (
	"log"

	"github.com/michaelperel/docker-lock/rewrite"
	"github.com/spf13/cobra"
)

// NewRewriteCmd creates the command 'rewrite' used in 'docker lock rewrite'.
func NewRewriteCmd() *cobra.Command {
	rewriteCmd := &cobra.Command{
		Use:   "rewrite",
		Short: "Rewrite files referenced by a Lockfile to use image digests",
		RunE: func(cmd *cobra.Command, args []string) error {
			flags, err := rewriterFlags(cmd)
			if err != nil {
				return err
			}

			configureLogger(flags.Verbose)

			log.Printf("Found flags '%+v'.", flags)

			rewriter, err := rewrite.NewRewriter(flags)
			if err != nil {
				return err
			}

			if err := rewriter.Rewrite(); err != nil {
				return err
			}

			return nil
		},
	}
	rewriteCmd.Flags().BoolP(
		"verbose", "v", false, "Show logs",
	)
	rewriteCmd.Flags().StringP(
		"lockfile-path", "l", "docker-lock.json", "Path to Lockfile",
	)
	rewriteCmd.Flags().StringP(
		"suffix", "s", "",
		"Create new Dockerfiles and docker-compose files "+
			"with a suffix rather than overwrite existing files",
	)
	rewriteCmd.Flags().StringP(
		"tempdir", "t", "",
		"Directory where a temporary directory will be created/deleted "+
			"during a rewrite transaction",
	)

	return rewriteCmd
}

// rewriterFlags gets values from the command and uses them to
// create Flags.
func rewriterFlags(cmd *cobra.Command) (*rewrite.Flags, error) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, err
	}

	lPath, err := cmd.Flags().GetString("lockfile-path")
	if err != nil {
		return nil, err
	}

	suffix, err := cmd.Flags().GetString("suffix")
	if err != nil {
		return nil, err
	}

	tmpDir, err := cmd.Flags().GetString("tempdir")
	if err != nil {
		return nil, err
	}

	return rewrite.NewFlags(lPath, suffix, tmpDir, verbose)
}
