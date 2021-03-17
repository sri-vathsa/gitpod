// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true,
	Short:  "Generates bash completion scripts",
	Long: `To load completion run

. <(gp completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(gp completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			fmt.Println("Received error in generating bash completion:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
