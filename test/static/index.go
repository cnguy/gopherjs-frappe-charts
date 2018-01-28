package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Make bar chart
	reportCountList := []float64{17, 40, 33, 44, 126, 156,
		324, 333, 478, 495, 373}
	barChartData := charts.NewChartData()
	barChartData.Labels = []string{
		"2007", "2008", "2009", "2010", "2011", "2012",
		"2013", "2014", "2015", "2016", "2017",
	}
	barChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("Events", reportCountList),
	}
	barChartTitle := "Fireball/Bolide Events - Yearly (more than 5 reports"
	barChart := charts.NewBarChart("#chart", barChartData).
		WithTitle(barChartTitle).
		WithHeight(180).
		WithColors([]string{"orange"}).
		SetIsNavigable(true).
		SetIsSeries(true).
		Render()

	// Make line chart
	lineChartValues := []float64{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0}
	lineChartData := charts.NewChartData()
	lineChartData.Labels = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	lineChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("", lineChartValues),
		charts.NewDataset("", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 12}),
	}
	lineChart := charts.NewLineChart("#chart-2", lineChartData).
		WithHeight(180).
		WithColors([]string{"green"}).
		SetIsSeries(true).
		Render()

	// Prepare update values for event listener
	moreLineData := charts.NewUpdateValueSets(
		[]float64{4, 0, 3, 1, 1, 2, 1, 2, 1, 0, 1, 1},
		[]float64{2, 3, 3, 2, 1, 4, 0, 1, 2, 7, 11, 4},
		[]float64{7, 7, 2, 4, 0, 1, 5, 3, 1, 2, 0, 1},
		[]float64{0, 2, 6, 2, 2, 1, 2, 3, 6, 3, 7, 10},
		[]float64{9, 10, 8, 10, 6, 5, 8, 8, 24, 15, 10, 13},
		[]float64{9, 13, 16, 9, 4, 5, 7, 10, 14, 22, 23, 24},
		[]float64{20, 22, 28, 19, 28, 19, 14, 19, 51, 37, 29, 38},
		[]float64{29, 20, 22, 16, 16, 19, 24, 26, 57, 31, 46, 2},
		[]float64{36, 24, 38, 27, 15, 22, 24, 38, 32, 57, 139, 26},
		[]float64{37, 36, 32, 33, 12, 34, 52, 45, 58, 57, 64, 35},
		[]float64{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0},
	)

	testLineData := charts.NewUpdateValueSets(
		[]float64{99, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		[]float64{99, 2, 3, 4, 5, 6, 7, 58, 9, 10, 11, 12},
		[]float64{99, 2, 3, 4, 45, 6, 7, 48, 9, 10, 11, 12},
		[]float64{50, 2, 3, 14, 5, 6, 7, 8, 39, 10, 11, 12},
		[]float64{99, 2, 3, 4, 5, 6, 72, 8, 29, 10, 11, 12},
		[]float64{99, 2, 3, 4, 35, 36, 7, 8, 9, 100, 11, 12},
		[]float64{99, 2, 13, 24, 5, 6, 7, 8, 9, 10, 11, 12},
		[]float64{99, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		[]float64{99, 2, 3, 4, 5, 6, 47, 8, 9, 10, 11, 12},
		[]float64{99, 2, 53, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		[]float64{99, 22, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	)

	barChart.AddDataSelectListener(func(event *charts.DataSelectEvent) {
		updateValues := &charts.UpdateValuesArgs{
			Values: []*charts.UpdateValueSet{moreLineData[event.Index], testLineData[event.Index]},
			// keep labels same as before
			Labels: lineChartData.Labels,
		}
		lineChart.UpdateValues(updateValues)
	})
}