// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
	"github.com/stormcat24/cloudinit-helper/client/meta"
)

var (
	metaUrlBase string
	metaConfig meta.Config
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Get EC2 Information",
	Long: `Get EC2 Information at the instance.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		metaConfig = meta.Config{
			UseMock: UseMock,
			UrlBase: metaUrlBase,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var ec2CmdRegion = &cobra.Command{
	Use:   "region",
	Short: "Get AWS Region in which the instance belongs.",
	Run: func(cmd *cobra.Command, args []string) {
		c := meta.NewClient(&metaConfig)
		az, err := c.GetAvailabilityZone()
		if err != nil {
			cmd.Out().Write([]byte(err.Error()))
		}
		cmd.Printf(az.GetRegion())
	},
}

var ec2CmdAz = &cobra.Command{
	Use:   "az",
	Short: "Get AWS Availability Zone in which the instance belongs.",
	Run: func(cmd *cobra.Command, args []string) {
		c := meta.NewClient(&metaConfig)
		az, err := c.GetAvailabilityZone()
		if err != nil {
			cmd.Out().Write([]byte(err.Error()))
		}
		cmd.Printf(az.Name)
	},
}

var ec2CmdInstanceID = &cobra.Command{
	Use:   "instance_id",
	Short: "Get InstanceID of this EC2 instance.",
	Run: func(cmd *cobra.Command, args []string) {
		c := meta.NewClient(&metaConfig)
		id, err := c.GetInstanceID()
		if err != nil {
			cmd.Out().Write([]byte(err.Error()))
		}
		cmd.Printf(id)
	},
}

func init() {
	RootCmd.AddCommand(ec2Cmd)

	ec2Cmd.PersistentFlags().StringVarP(&metaUrlBase, "meta-url-base", "", "http://169.254.169.254", "Instance Meta Data API URL Base")

	ec2Cmd.AddCommand(ec2CmdRegion)
	ec2Cmd.AddCommand(ec2CmdAz)
	ec2Cmd.AddCommand(ec2CmdInstanceID)

}
