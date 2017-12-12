package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

type HeatmapArgs struct {
	*js.Object
	Parent          string                 `js:"parent"`
	Data            map[string]interface{} `js:"data"`
	Type            string                 `js:"type"`
	Height          int                    `js:"height"`
	DiscreteDomains int                    `js:"discrete_domains"`
	Start           *js.Object             `js:"start"` // TODO
	LegendColors    []string               `js:"legend_colors"`
}

func NewHeatmapChart(parent string, data map[string]int) *HeatmapArgs {
	new := &HeatmapArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "heatmap"
	fixedData := make(map[string]interface{})
	for k, v := range data {
		fixedData[k] = v
	}
	new.Data = fixedData
	return new
}

func (chartArgs *HeatmapArgs) WithHeight(height int) *HeatmapArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *HeatmapArgs) SetDiscreteDomain(val bool) *HeatmapArgs {
	chartArgs.Set("discrete_domains", utils.Btoi(val))
	return chartArgs
}

func (chartArgs *HeatmapArgs) WithLegendColors(colors []string) *HeatmapArgs {
	chartArgs.LegendColors = colors
	return chartArgs
}

func (chartArgs *HeatmapArgs) Render() {
	// Set defaults that aren't handled by frappe here.
	if chartArgs.Height == 0 {
		chartArgs.Height = 115
	}
	// Create actual chart.
	js.Global.Get("Chart").New(&chartArgs)
}
