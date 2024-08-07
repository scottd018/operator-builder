// Copyright 2024 Nukleros
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/nukleros/operator-builder/internal/plugins/workload"
	"github.com/nukleros/operator-builder/pkg/cli"
)

func main() {
	command, err := cli.NewKubebuilderCLI(workload.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	if command == nil {
		log.Println("skipping command execution...")
		os.Exit(0)
	}

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
