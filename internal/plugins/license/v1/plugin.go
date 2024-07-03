// Copyright 2024 Nukleros
// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: MIT

package v1

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	cfgv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"

	"github.com/nukleros/operator-builder/internal/plugins"
)

const pluginName = "license." + plugins.DefaultNameQualifier

//nolint:gochecknoglobals //needed for plugin architecture
var (
	pluginVersion            = plugin.Version{Number: 1}
	supportedProjectVersions = []config.Version{cfgv3.Version}
	pluginKey                = plugin.KeyFor(Plugin{})
)

var (
	_ plugin.Plugin = Plugin{}
	_ plugin.Init   = Plugin{}
)

type Plugin struct {
	initSubcommand
}

func (Plugin) Name() string                               { return pluginName }
func (Plugin) Version() plugin.Version                    { return pluginVersion }
func (Plugin) SupportedProjectVersions() []config.Version { return supportedProjectVersions }
func (p Plugin) GetInitSubcommand() plugin.InitSubcommand { return &p.initSubcommand }
