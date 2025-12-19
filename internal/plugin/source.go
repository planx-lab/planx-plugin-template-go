package plugin

import (
	"context"
	"encoding/json"

	"github.com/planx-lab/planx-sdk-go/sdk"
)

type Config struct {
	Message string `json:"message"`
}

type Source struct {
	msg string
}

func New() sdk.SourceSPI {
	return &Source{}
}

func (s *Source) Init(ctx context.Context, cfg []byte) error {
	var c Config
	if err := json.Unmarshal(cfg, &c); err != nil {
		return err
	}
	s.msg = c.Message
	return nil
}

func (s *Source) ReadBatch() (sdk.Batch, error) {
	// Minimal example: emit a string
	return s.msg, nil
}

func (s *Source) Close() error {
	return nil
}
