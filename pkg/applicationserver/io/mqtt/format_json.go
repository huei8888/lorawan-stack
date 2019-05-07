// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package mqtt

import (
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/formatters"
	mqtt "go.thethings.network/lorawan-stack/pkg/applicationserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/pubsub"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/pubsub/topics"
)

type json struct {
	topics.Layout
	formatters.Formatter
}

// JSON is a format that uses the default topic layout and JSON formatter.
var JSON pubsub.Format = &json{
	Layout:    mqtt.Default,
	Formatter: formatters.JSON,
}
