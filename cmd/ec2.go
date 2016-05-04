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
	"errors"
	"github.com/stormcat24/cloudinit-helper/client/ec2"
	"os"
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
}

var ec2CmdRegion = &cobra.Command{
	Use:   "region",
	Short: "Get AWS Region in which the instance belongs.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := meta.NewClient(&metaConfig)
		az, err := c.GetAvailabilityZone()
		if err != nil {
			return err
		}
		cmd.SetOutput(os.Stdout)
		cmd.Println(az.GetRegion())
		return nil
	},
}

var ec2CmdAz = &cobra.Command{
	Use:   "az",
	Short: "Get AWS Availability Zone in which the instance belongs.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := meta.NewClient(&metaConfig)
		az, err := c.GetAvailabilityZone()
		if err != nil {
			return err
		}
		cmd.SetOutput(os.Stdout)
		cmd.Println(az.Name)
		return nil
	},
}

var ec2CmdInstanceID = &cobra.Command{
	Use:   "instance_id",
	Short: "Get InstanceID of this EC2 instance.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := meta.NewClient(&metaConfig)
		id, err := c.GetInstanceID()
		if err != nil {
			return err
		}
		cmd.SetOutput(os.Stdout)
		cmd.Println(id)
		return nil
	},
}

var ec2CmdDescribeInstanceTag = &cobra.Command{
	Use:   "describe-instance-tag",
	Short: "Describe EC2 Instance Tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			return err
		}

		if region == "" {
			return errors.New("--region is not specified.")
		}

		instanceId, err := cmd.Flags().GetString("instance-id")
		if err != nil {
			return err
		}

		if instanceId == "" {
			return errors.New("--instance-id is not specified.")
		}

		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			return err
		}

		if tag == "" {
			return errors.New("--tag is not specified.")
		}

		client := ec2.NewClient(region)
		result, err := client.DescribeInstance(instanceId)
		if err != nil {
			return err
		}

		for _, t := range result.Tags {
			if tag == *t.Key {
				cmd.SetOutput(os.Stdout)
				cmd.Println(*t.Value)
				return nil
			}
		}

		cmd.Printf("Tag=%v is not found at %v", tag, instanceId)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(ec2Cmd)

	ec2Cmd.PersistentFlags().StringVarP(&metaUrlBase, "meta-url-base", "", "http://169.254.169.254", "Instance Meta Data API URL Base")

	ec2CmdDescribeInstanceTag.Flags().StringP("region", "r", "", "Target Region")
	ec2CmdDescribeInstanceTag.Flags().StringP("instance-id", "i", "", "Target Instance ID")
	ec2CmdDescribeInstanceTag.Flags().StringP("tag", "t", "", "Target Tag")

	ec2Cmd.AddCommand(ec2CmdRegion)
	ec2Cmd.AddCommand(ec2CmdAz)
	ec2Cmd.AddCommand(ec2CmdInstanceID)
	ec2Cmd.AddCommand(ec2CmdDescribeInstanceTag)

}
