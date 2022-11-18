// Copyright 2021 The frp Authors
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

package frps

import (
	"Erfrp/pkg/config"
	"fmt"
	"os"

	"github.com/Gogods/cobra"
)

func init() {
	rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify that the configures is valid",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			fmt.Println("no config file is specified")
			return nil
		}
		iniContent, err := config.GetRenderedConfFromFile(cfgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = parseServerCommonCfg(CfgFileTypeIni, iniContent)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("frps: the configuration file %s syntax is ok\n", cfgFile)
		return nil
	},
}
