package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

type PieChartArgs struct {
	*js.Object
	Parent string     `js:"parent"`
	Title  string     `js:"title"`
	Data   *ChartData `js:"data"`
	Type   string     `js:"type"`
	Height int        `js:"height"`
}

func NewPieChartArgs(parent string, title string, data *ChartData, height int) *PieChartArgs {
	new := &PieChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Data = data
	new.Height = height
	new.Type = "pie"
	return new
}

func (chartArgs *PieChartArgs) Render() {
	js.Global.Get("Chart").New(&chartArgs)
}
