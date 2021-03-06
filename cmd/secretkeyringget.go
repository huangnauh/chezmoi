package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	keyring "github.com/zalando/go-keyring"
)

var keyringGetCmd = &cobra.Command{
	Use:     "get",
	Args:    cobra.NoArgs,
	Short:   "Get a value from keyring",
	PreRunE: config.ensureNoError,
	RunE:    config.runKeyringGetCmd,
}

func init() {
	keyringCmd.AddCommand(keyringGetCmd)
}

func (c *Config) runKeyringGetCmd(cmd *cobra.Command, args []string) error {
	value, err := keyring.Get(c.keyring.service, c.keyring.user)
	if err != nil {
		return err
	}
	fmt.Println(value)
	return nil
}
