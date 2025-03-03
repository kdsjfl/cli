/*
Copyright © 2019 Doppler <support@doppler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/DopplerHQ/cli/pkg/utils"
	"github.com/spf13/cobra"
)

var enclaveSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the Doppler CLI for Enclave",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		deprecatedCommand("setup")
		setup(cmd, args)
	},
}

func init() {
	enclaveSetupCmd.Flags().StringP("project", "p", "", "enclave project (e.g. backend)")
	enclaveSetupCmd.Flags().StringP("config", "c", "", "enclave config (e.g. dev)")
	enclaveSetupCmd.Flags().Bool("no-interactive", false, "do not prompt for information. if the project or config is not specified, an error will be thrown.")
	enclaveSetupCmd.Flags().Bool("no-save-token", false, "do not save the token to the config when passed via flag or environment variable.")

	// deprecated
	enclaveSetupCmd.Flags().Bool("no-prompt", false, "do not prompt for information. if the project or config is not specified, an error will be thrown.")
	if err := enclaveSetupCmd.Flags().MarkDeprecated("no-prompt", "please use --no-interactive instead"); err != nil {
		utils.HandleError(err)
	}
	if err := enclaveSetupCmd.Flags().MarkHidden("no-prompt"); err != nil {
		utils.HandleError(err)
	}

	enclaveCmd.AddCommand(enclaveSetupCmd)
}
