package helpers

import (
	"github.com/cometbft/cometbft/abci/types"
)

type HumanReadableEventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index,omitempty"`
}

type HumanReadableEvent struct {
	Type       string                        `json:"type"`
	Attributes []HumanReadableEventAttribute `json:"attributes"`
}

func ReadableEvents(events []types.Event) []HumanReadableEvent {
	readableEvents := make([]HumanReadableEvent, 0, len(events))
	for _, event := range events {
		readableAttributes := make([]HumanReadableEventAttribute, 0, len(event.Attributes))
		for i, attribute := range event.Attributes {
			readableAttributes[i] = HumanReadableEventAttribute{
				Key:   attribute.Key,
				Value: attribute.Value,
				Index: attribute.Index,
			}
		}
		readableEvents = append(readableEvents, HumanReadableEvent{
			Type:       event.Type,
			Attributes: readableAttributes,
		})
	}
	return readableEvents
}
