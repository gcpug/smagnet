package smagnet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/spanner/v1"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type SpannerMagnetMessage struct {
	ProjectID  string `json:"projectId"`
	InstanceID string `json:"instanceId"`
	NodeNumber int64  `json:"nodeNumber"`
}

// HandleSpannerMagnet consumes a Pub/Sub message.
func HandleSpannerMagnet(ctx context.Context, m PubSubMessage) error {
	var msg SpannerMagnetMessage
	if err := json.Unmarshal(m.Data, &msg); err != nil {
		log.Printf("failed json.Unmarshal. body=%s, %+v", string(m.Data), err)
		return err
	}
	log.Printf("SpannerMagnetMessage is %+v", msg)

	client, err := google.DefaultClient(ctx, spanner.SpannerAdminScope)
	if err != nil {
		log.Printf("google.DefaultClient, %+v", err)
		return err
	}
	admin, err := spanner.New(client)
	if err != nil {
		log.Printf("spanner.New, %+v", err)
		return err
	}
	resp, err := admin.Projects.Instances.Patch(fmt.Sprintf("projects/%s/instances/%s", msg.ProjectID, msg.InstanceID), &spanner.UpdateInstanceRequest{
		Instance: &spanner.Instance{
			NodeCount: msg.NodeNumber,
		},
		FieldMask: "nodeCount",
	}).Do()
	if err != nil {
		log.Printf("failed change spanner node count. %+v", err)
		return err
	}
	fmt.Printf("Response : %+v", resp)

	return nil
}
