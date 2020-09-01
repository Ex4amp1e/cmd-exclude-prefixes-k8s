// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prefixcollector

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type clientSetKeyType string

// ClientSetKey is ClientSet key in context map
const ClientSetKey clientSetKeyType = "clientsetKey"

// FromContext returns ClientSet from context ctx
func FromContext(ctx context.Context) kubernetes.Interface {
	return ctx.Value(ClientSetKey).(kubernetes.Interface)
}
