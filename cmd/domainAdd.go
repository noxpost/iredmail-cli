// Copyright © 2018 Christian Nolte
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"log"

	"github.com/drlogout/iredmail-cli/iredmail"
	"github.com/spf13/cobra"
)

// domainAddCmd represents the add command
var domainAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a domain",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a domain name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		server, err := iredmail.New()
		if err != nil {
			log.Fatal(err)
		}
		defer server.Close()

		description := cmd.Flag("description").Value.String()
		settings := cmd.Flag("settings").Value.String()
		if description == "" {
			description = args[0]
		}
		if settings == "" {
			settings = iredmail.DomainDefaultSettings
		}

		domain := iredmail.Domain{
			Domain:      args[0],
			Description: description,
			Settings:    settings,
		}

		err = server.DomainAdd(domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	domainCmd.AddCommand(domainAddCmd)

	domainAddCmd.Flags().StringP("description", "d", "", "domain description")
	domainAddCmd.Flags().StringP("settings", "s", "", "domain settings (default: default_user_quota:2048)")
}
