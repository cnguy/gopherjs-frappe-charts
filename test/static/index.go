package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		"12am-3am", "3am-6pm", "6am-9am", "9am-12am",
		"12pm-3pm", "3pm-6pm", "6pm-9pm", "9am-12am",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			"light-blue",
			[]interface{}{25, 40, 30, 35, 8, 52, 17, -4},
		),
		charts.NewDataset(
			"Another Set",
			"violet",
			[]interface{}{25, 50, -10, 15, 18, 32, 27, 14},
		),
	}
	chartData.SpecificValues = []*charts.SpecificValue{
		charts.NewSpecificValue(
			"Altitude",
			"dashed",
			38,
		),
	}

	// Prepare constructor arguments via the NewChartArgs helper
	chartArgs := charts.NewBarChartArgs("#chart", "My Awesome Chart", chartData, 250)
	// chartArgs.Parent = "#chart"
	// chartArgs.Title = "My Awesome Chart"
	// chartArgs.IsNavigable = ...
	// chartArgs.Heatline = ...

	chart := chartArgs.Render()
	println(chart) // to suppress the annoying error
}
