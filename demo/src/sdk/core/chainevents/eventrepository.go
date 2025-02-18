// Scry Info.  All rights reserved.
// license that can be found in the license file.

package chainevents

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dp/demo/src/sdk/core/ethereum/events"
)

type EventCallback func(event events.Event) bool

type EventRepository struct {
    mapEventSubscribe map[string]map[common.Address]EventCallback
}

func NewEventRepository() *EventRepository {
    return &EventRepository{
        mapEventSubscribe: make(map[string]map[common.Address]EventCallback),
    }
}
