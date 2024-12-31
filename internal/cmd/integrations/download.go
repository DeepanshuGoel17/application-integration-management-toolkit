// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integrations

import (
	"fmt"
	"internal/apiclient"
	"internal/client/integrations"
	"strconv"

	"github.com/spf13/cobra"
)

// DownloadVerCmd to publish an integration flow version
var DownloadVerCmd = &cobra.Command{
	Use:   "download",
	Short: "Download an integration flow version",
	Long:  "Download an integration flow version",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		cmdProject := cmd.Flag("proj")
		cmdRegion := cmd.Flag("reg")
		version := cmd.Flag("ver").Value.String()
		userLabel := cmd.Flag("user-label").Value.String()
		snapshot := cmd.Flag("snapshot").Value.String()
		latest, _ := strconv.ParseBool(cmd.Flag("latest").Value.String())

		if err = apiclient.SetRegion(cmdRegion.Value.String()); err != nil {
			return err
		}
		if err = validate(version, userLabel, snapshot, latest); err != nil {
			return err
		}
		return apiclient.SetProjectID(cmdProject.Value.String())
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		version := cmd.Flag("ver").Value.String()
		userLabel := cmd.Flag("user-label").Value.String()
		snapshot := cmd.Flag("snapshot").Value.String()
		name := cmd.Flag("name").Value.String()

		latest := ignoreLatest(version, userLabel, snapshot)

		if latest {
			apiclient.DisableCmdPrintHttpResponse()
			// list integration versions, order by state=SNAPSHOT, page size = 1 and return basic info
			respBody, err := integrations.ListVersions(name, 1, "", "state=SNAPSHOT",
				"snapshot_number", false, false, true)
			if err != nil {
				return fmt.Errorf("unable to list versions: %v", err)
			}
			version, err = getIntegrationVersion(respBody)
			if err != nil {
				return err
			}
			apiclient.EnableCmdPrintHttpResponse()
			_, err = integrations.Download(name, version)
		} else if version != "" {
			_, err = integrations.Download(name, version)
		} else if userLabel != "" {
			_, err = integrations.DownloadUserLabel(name, userLabel)
		} else if snapshot != "" {
			_, err = integrations.DownloadSnapshot(name, snapshot)
		}
		return err
	},
}

func init() {
	var name, userLabel, snapshot, version string
	var latest bool

	DownloadVerCmd.Flags().StringVarP(&name, "name", "n",
		"", "Integration flow name")
	DownloadVerCmd.Flags().StringVarP(&version, "ver", "v",
		"", "Integration flow version")
	DownloadVerCmd.Flags().StringVarP(&userLabel, "user-label", "u",
		"", "Integration flow user label")
	DownloadVerCmd.Flags().StringVarP(&snapshot, "snapshot", "s",
		"", "Integration flow snapshot number")
	DownloadVerCmd.Flags().BoolVarP(&latest, "latest", "",
		true, "Downloads the integeration version with the highest snapshot number in SNAPSHOT state; default is true")

	_ = DownloadVerCmd.MarkFlagRequired("name")
}
