package cli

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/sibukixxx/migration-mirror/internal/cli/output"
	"github.com/sibukixxx/migration-mirror/internal/usecase"
)

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check that the host has the external tools migration-mirror needs",
		Long: `doctor inspects the local environment for the external tools migration-mirror
delegates to (wp-cli, wrangler, rclone, awscli, git, jq, php, node).

Each tool is reported as required or optional. Missing required tools cause
a non-zero exit code so this command is safe to use as a CI gate.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDoctor(cmd)
		},
	}
	return cmd
}

func runDoctor(cmd *cobra.Command) error {
	rep := usecase.RunDoctor(cmd.Context())

	if mustBool(cmd, "json") {
		if err := output.JSON(cmd.OutOrStdout(), rep); err != nil {
			return err
		}
	} else {
		w := cmd.OutOrStdout()
		if err := output.Println(w, "migration-mirror doctor"); err != nil {
			return err
		}
		for _, c := range rep.Checks {
			tag := "optional"
			if c.Required {
				tag = "required"
			}
			status := "missing"
			if c.Found {
				status = c.Path
			}
			if err := output.Printf(w, "  [%s] %-10s %s\n", tag, c.Name, status); err != nil {
				return err
			}
		}
		msg := "OK"
		if !rep.OK {
			msg = "FAIL: required tool(s) missing"
		}
		if err := output.Println(w, msg); err != nil {
			return err
		}
	}

	if !rep.OK {
		return errors.New("required tool missing")
	}
	return nil
}
