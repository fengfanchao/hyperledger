/*
Copyright IBM Corp. 2017 All Rights Reserved.

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

package util

import (
	"github.com/hyperledger/fabric/gossip/common"
	proto "github.com/hyperledger/fabric/protos/gossip"
)

type MembershipStore map[string]*proto.SignedGossipMessage

// msgByID returns a message stored by a certain ID, or nil
// if such an ID isn't found
func (m MembershipStore) MsgByID(pkiID common.PKIidType) *proto.SignedGossipMessage {
	if msg, exists := m[string(pkiID)]; exists {
		return msg
	}
	return nil
}

// Put associates msg with the given pkiID
func (m MembershipStore) Put(pkiID common.PKIidType, msg *proto.SignedGossipMessage) {
	m[string(pkiID)] = msg
}

// Remove removes a message with a given pkiID
func (m MembershipStore) Remove(pkiID common.PKIidType) {
	delete(m, string(pkiID))
}

// ToSlice returns a slice backed by the elements
// of the MembershipStore
func (m MembershipStore) ToSlice() []*proto.SignedGossipMessage {
	members := make([]*proto.SignedGossipMessage, len(m))
	i := 0
	for _, member := range m {
		members[i] = member
		i++
	}
	return members
}
