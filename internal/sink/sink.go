package sink

import (
	"context"

	"github.com/planx-lab/planx-sdk-go/spi"
)

type Sink struct{}

func New() spi.Sink { return &Sink{} }

func (s *Sink) Create(ctx context.Context, cfg []byte) error { return nil }
func (s *Sink) Write(ctx context.Context, data []byte) error { return nil }
func (s *Sink) Stop(ctx context.Context) error { <-ctx.Done(); return nil }
