/*
Copyright 2022 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package volcenginecloud

import (
	"github.com/openkruise/kruise-game/cloudprovider"
	"k8s.io/klog/v2"
)

const (
	VolcengineCloud = "VolcengineCloud"
)

var (
	volcengineCloudProvider = &Provider{
		plugins: make(map[string]cloudprovider.Plugin),
	}
)

type Provider struct {
	plugins map[string]cloudprovider.Plugin
}

func (vp *Provider) Name() string {
	return VolcengineCloud
}

func (vp *Provider) ListPlugins() (map[string]cloudprovider.Plugin, error) {
	if vp.plugins == nil {
		return make(map[string]cloudprovider.Plugin), nil
	}

	return vp.plugins, nil
}

// register plugin of cloud provider and different cloud providers
func (vp *Provider) registerPlugin(plugin cloudprovider.Plugin) {
	name := plugin.Name()
	if name == "" {
		klog.Fatal("empty plugin name")
	}
	vp.plugins[name] = plugin
}

func NewVolcengineCloudProvider() (cloudprovider.CloudProvider, error) {
	return volcengineCloudProvider, nil
}
