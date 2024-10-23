// Copyright 2024 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"fmt"

	"github.com/flowbehappy/tigate/coordinator/changefeed"
	"github.com/flowbehappy/tigate/heartbeatpb"
	"github.com/flowbehappy/tigate/pkg/messaging"
	"github.com/flowbehappy/tigate/pkg/node"
	"github.com/pingcap/log"
	"github.com/pingcap/tiflow/cdc/model"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

// AddMaintainerOperator is an operator to schedule a maintainer to a node
type AddMaintainerOperator struct {
	cf       *changefeed.Changefeed
	dest     node.ID
	finished atomic.Bool
	removed  atomic.Bool
	db       *changefeed.ChangefeedDB
}

func NewAddMaintainerOperator(
	db *changefeed.ChangefeedDB,
	cf *changefeed.Changefeed,
	dest node.ID) *AddMaintainerOperator {
	return &AddMaintainerOperator{
		cf:   cf,
		dest: dest,
		db:   db,
	}
}

func (m *AddMaintainerOperator) Check(from node.ID, status *heartbeatpb.MaintainerStatus) {
	if !m.finished.Load() && from == m.dest && status.State == heartbeatpb.ComponentState_Working {
		log.Info("maintainer report working status",
			zap.String("changefeed", m.cf.ID.String()))
		m.finished.Store(true)
	}
}

func (m *AddMaintainerOperator) Schedule() *messaging.TargetMessage {
	if m.finished.Load() || m.removed.Load() {
		return nil
	}
	return m.cf.NewAddMaintainerMessage(m.dest)
}

// OnNodeRemove is called when node offline, and the maintainer must already move to absent status and will be scheduled again
func (m *AddMaintainerOperator) OnNodeRemove(n node.ID) {
	if n == m.dest {
		m.finished.Store(true)
		m.removed.Store(true)
	}
}

func (m *AddMaintainerOperator) ID() model.ChangeFeedID {
	return m.cf.ID
}

func (m *AddMaintainerOperator) IsFinished() bool {
	return m.finished.Load()
}

func (m *AddMaintainerOperator) OnTaskRemoved() {
	m.finished.Store(true)
	m.removed.Store(true)
}

func (m *AddMaintainerOperator) Start() {
	m.db.BindChangefeedToNode("", m.dest, m.cf)
}

func (m *AddMaintainerOperator) PostFinish() {
	if !m.removed.Load() {
		m.db.MarkMaintainerReplicating(m.cf)
	}
}

func (m *AddMaintainerOperator) String() string {
	return fmt.Sprintf("add maintainer operator: %s, dest:%s",
		m.cf.ID, m.dest)
}

func (m *AddMaintainerOperator) Type() string {
	return "add"
}