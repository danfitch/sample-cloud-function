package p

import (
	"context"
	"log"
	"os"
	"time"
)

var (
	stdLogger = log.New(os.Stdout, "", 0)
	logger    = log.New(os.Stderr, "", 0)
)

// PubSubMessage is the decoded message from the pub/sub channel
type PubSubMessage struct {
	Data        []byte            `json:"data"`
	Attributes  map[string]string `json:"attributes"`
	MessageID   string            `json:"messageId"`
	PublishTime time.Time         `json:"publishTime"`
}

// SampleCF is a sample cloud function triggered by a PubSubMessage
func SampleCF(ctx context.Context, msg PubSubMessage) error {
	stdLogger.Println("SampleCF")
	return nil
}
