package main

import (
	"github.com/planx-lab/planx-plugin-template-go/internal/processor"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.ServeProcessor(processor.New)
}
