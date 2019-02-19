package smagnet

import (
	"context"
	"encoding/json"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type SpannerMagnetMessage struct {
	ProjectID  string `json:"projectId"`
	InstanceID string `json:"instanceId"`
	NodeNumber int    `json:"nodeNumber"`
}

// HandleSpannerMagnet consumes a Pub/Sub message.
func HandleSpannerMagnet(ctx context.Context, m PubSubMessage) error {
	var msg SpannerMagnetMessage
	if err := json.Unmarshal(m.Data, &msg); err != nil {
		log.Printf("failed json.Unmarshal. body=%s, %+v", string(m.Data), err)
		return err
	}
	log.Printf("SpannerMagnetMessage is %+v", msg)

	return nil
}
