package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	chartData := charts.NewChartData()
	chartData.Labels = []string{
		"1", "2", "3",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			[]interface{}{25, 40, 30},
		),
		charts.NewDataset(
			"Another Set",
			[]interface{}{25, 50, -10},
		),
	}

	_ = charts.NewBarChart("#chart", chartData).Render()
}
