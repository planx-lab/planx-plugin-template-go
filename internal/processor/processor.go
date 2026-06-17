package processor

import (
	"context"

	"github.com/planx-lab/planx-sdk-go/spi"
)

type Processor struct{}

func New() spi.Processor { return &Processor{} }

func (p *Processor) Create(ctx context.Context, cfg []byte) error { return nil }
func (p *Processor) Process(ctx context.Context, data []byte) ([]byte, error) { return data, nil }
func (p *Processor) Stop(ctx context.Context) error { <-ctx.Done(); return nil }
