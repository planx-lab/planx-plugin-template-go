package main

import (
	"github.com/planx-lab/planx-plugin-template-go/internal/processor"
	"github.com/planx-lab/planx-plugin-template-go/internal/sink"
	"github.com/planx-lab/planx-plugin-template-go/internal/source"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.Serve(sdk.Plugin{
		ID:          "template",
		Version:     "1.0.0",
		DisplayName: "Connector Template",
		Description: "Reference multi-component connector (source + processor + sink).",
		Components: []sdk.ComponentSpec{
			{ID: "source", Kind: sdk.KindSource, DisplayName: "Template Source", Source: source.New},
			{ID: "processor", Kind: sdk.KindProcessor, DisplayName: "Template Processor", Processor: processor.New},
			{ID: "sink", Kind: sdk.KindSink, DisplayName: "Template Sink", Sink: sink.New},
		},
	})
}
