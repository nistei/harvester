/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	BackingImage() BackingImageController
	BackingImageDataSource() BackingImageDataSourceController
	Setting() SettingController
	Volume() VolumeController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) BackingImage() BackingImageController {
	return NewBackingImageController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImage"}, "backingimages", true, c.controllerFactory)
}
func (c *version) BackingImageDataSource() BackingImageDataSourceController {
	return NewBackingImageDataSourceController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImageDataSource"}, "backingimagedatasources", true, c.controllerFactory)
}
func (c *version) Setting() SettingController {
	return NewSettingController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Setting"}, "settings", true, c.controllerFactory)
}
func (c *version) Volume() VolumeController {
	return NewVolumeController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Volume"}, "volumes", true, c.controllerFactory)
}
