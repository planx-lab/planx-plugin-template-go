package source

import (
	"context"
	"io"

	"github.com/planx-lab/planx-sdk-go/sdk"
)

// Source is a minimal SOURCE component template. Replace ReadBatch with real logic.
type Source struct{}

func New() sdk.SourceSPI { return &Source{} }

func (s *Source) Init(_ context.Context, _ []byte) error { return nil }

// ReadBatch returns EOF immediately. Replace with a real read loop.
func (s *Source) ReadBatch() (sdk.Batch, error) { return nil, io.EOF }

func (s *Source) Close() error { return nil }
