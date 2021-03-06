// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crane

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() { Root.AddCommand(NewCmdDigest()) }

// NewCmdDigest creates a new cobra.Command for the digest subcommand.
func NewCmdDigest() *cobra.Command {
	return &cobra.Command{
		Use:   "digest",
		Short: "Get the digest of an image",
		Args:  cobra.ExactArgs(1),
		Run:   digest,
	}
}

func digest(_ *cobra.Command, args []string) {
	ref := args[0]
	desc, err := getManifest(ref)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(desc.Digest)
}
