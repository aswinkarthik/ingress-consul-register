// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/api"
	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/aswinkarthik93/ingress-consul-register/pkg/engine"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting up...")
		runStart()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringVarP(&config.Cfg.IngressClass, "ingress-class", "i", "nginx", "Ingress class to watch for")
	startCmd.Flags().StringVarP(&config.Cfg.ConsulDomain, "consul-domain", "d", "consul", "Domain with which consul is configured")
	startCmd.Flags().StringVar(&config.Cfg.ControllerService, "ingress-controller-service", "", "Ingress controller service that is configured for the ingress-class")
	startCmd.Flags().StringVar(&config.Cfg.ControllerServiceNamespace, "ingress-controller-service-namespace", "default", "Namespace of the ingress controller")
}

func runStart() {
	if err := engine.Initialize(); err != nil {
		log.Fatal(err)
	}

	go engine.RunOnce()

	go engine.StartWatching()

	api.StartServer()
}
