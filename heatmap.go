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
	LegendColors    []string               `js:"legend_colors"`
}

func NewHeatmapArgs(parent string, data map[string]interface{}, height int) *HeatmapArgs {
	new := &HeatmapArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Data = data
	new.Height = height
	new.Type = "heatmap"
	return new
}

func (heatmapArgs *HeatmapArgs) SetDiscreteDomain(val bool) {
	heatmapArgs.Set("discrete_domains", utils.Btoi(val))
}

func (heatmapArgs *HeatmapArgs) Render() {
	js.Global.Get("Chart").New(&heatmapArgs)
}
