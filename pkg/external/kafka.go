package external

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func CheckKafka(brokerList []string) error {
	// make a new reader that consumes from _milvus-operator, partition 0, at offset 0
	if len(brokerList) == 0 {
		return errors.New("broker list is empty")
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerList,
		Topic:   "_milvus-operator",
	})
	defer r.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := r.SetOffsetAt(ctx, time.Now())
	return errors.Wrap(err, "check consume offset from borker failed")
}
