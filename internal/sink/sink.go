package sink

import (
	"context"
	"fmt"

	"github.com/planx-lab/planx-sdk-go/sdk"
)

// Sink is a minimal SINK component template. Replace WriteBatch with real logic.
type Sink struct{}

func New() sdk.SinkSPI { return &Sink{} }

func (s *Sink) Init(_ context.Context, _ []byte) error { return nil }

func (s *Sink) WriteBatch(batch sdk.Batch) error {
	fmt.Printf("[SINK] template received: %v\n", batch)
	return nil
}

func (s *Sink) Close() error { return nil }
