/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package networking

import (
	"knative.dev/pkg/apis"
	"knative.dev/reconciler-test/pkg/config"
)

// Config is the Networking configuration
type Config struct {
	config.VersionSpec
	Name string `desc:"the name of the networking layer"`
}

func (c *Config) Validate() *apis.FieldError {
	errs := c.VersionSpec.Validate()

	// TODO: other networking layer
	if c.Name != "kourier" {
		errs = apis.ErrInvalidValue("c.Name", "name").Also(errs)
	}
	return errs
}
