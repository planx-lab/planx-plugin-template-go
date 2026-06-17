package source

import (
	"context"
	"time"

	"github.com/planx-lab/planx-sdk-go/spi"
)

type Source struct{}

func New() spi.Source { return &Source{} }

func (s *Source) Create(ctx context.Context, cfg []byte) error { return nil }
func (s *Source) Start(ctx context.Context) error              { return nil }
func (s *Source) Read(ctx context.Context) ([]byte, error)     { return nil, nil }
func (s *Source) Ack(ctx context.Context) error                { return nil }
func (s *Source) Stop(ctx context.Context) error                { <-ctx.Done(); return nil }
