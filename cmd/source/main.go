package main

import (
	"github.com/planx-lab/planx-plugin-template-go/internal/source"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.ServeSource(source.New)
}
