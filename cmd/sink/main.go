package main

import (
	"github.com/planx-lab/planx-plugin-template-go/internal/sink"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.ServeSink(sink.New)
}
