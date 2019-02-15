// Copyright © 2018 Banzai Cloud
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

package api

import pkgSecret "github.com/banzaicloud/pipeline/pkg/secret"

// CreateResourceGroupRequest describes the resource group create request
type CreateResourceGroupRequest struct {
	Name     string             `json:"name" binding:"required"`
	Location string             `json:"location" binding:"required"`
	SecretId pkgSecret.SecretID `json:"secretId" binding:"required"`
}

// CreateResourceGroupResponse describes the resource group create response
type CreateResourceGroupResponse struct {
	Name string `json:"name" binding:"required"`
}
