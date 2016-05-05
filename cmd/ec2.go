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
	"github.com/stormcat24/cloudinit-helper/client/ec2meta"
	"errors"
	"github.com/stormcat24/cloudinit-helper/client/ec2"
	"os"
	"encoding/json"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Get EC2 Instance own information",
}

var ec2CmdMeta = &cobra.Command{
	Use:   "meta",
	Short: "Get metadata of EC2 Instance own information(JSON)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := ec2meta.NewClient(UseMock)
		doc, err := c.GetInstanceIdentityDocument()
		if err != nil {
			return err
		}

		data, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		cmd.SetOutput(os.Stdout)
		cmd.Println(string(data))
		return nil
	},
}

var ec2CmdDescribeTag = &cobra.Command{
	Use:   "describe-tag",
	Short: "Describe EC2 Instance Tag",
	RunE: func(cmd *cobra.Command, args []string) error {

		c := ec2meta.NewClient(UseMock)
		doc, err := c.GetInstanceIdentityDocument()
		if err != nil {
			return err
		}

		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			return err
		}

		if tag == "" {
			return errors.New("--tag is not specified.")
		}

		client := ec2.NewClient(doc.Region)
		result, err := client.DescribeInstance(doc.InstanceID)
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

		cmd.Printf("Tag=%v is not found at %v", tag, doc.InstanceID)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(ec2Cmd)

	ec2CmdDescribeTag.Flags().StringP("tag", "t", "", "Target Tag")

	ec2Cmd.AddCommand(ec2CmdMeta)
	ec2Cmd.AddCommand(ec2CmdDescribeTag)

}
