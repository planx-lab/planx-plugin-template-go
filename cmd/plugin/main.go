package main

import (
	"github.com/planx-lab/planx-plugin-template-go/internal/plugin"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.ServeSource(plugin.New)
}
