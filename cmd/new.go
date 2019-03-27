// Copyright © 2019 Lukas Jarosch
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
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/project"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

// rootCmd represents the base command when called without any subcommands
var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Setup a new microservice structure",
	Run:   new,
}

func new(cmd *cobra.Command, args []string) {

	logrus.SetLevel(logrus.DebugLevel)

	serviceName := "greeter"

	projectPath, _ := os.Getwd()
	projectPath = path.Join(projectPath, "examples", serviceName)

	// create a bare-bones Godin project
	godin := project.NewGodinProject(serviceName, projectPath)

	// add all required folders
	godin.AddFolder("internal")
	godin.AddFolder("cmd")
	godin.AddFolder(path.Join("cmd", serviceName))
	godin.AddFolder("k8s")
	godin.AddFolder("internal/server")
	godin.AddFolder("internal/service")
	godin.AddFolder("internal/config")

	if err := godin.MkdirAll(); err != nil {
		logrus.Fatal(err)
	}

	// add some basic templates
	godin.AddTemplate(template.NewTemplate())
	godin.Render()

}
