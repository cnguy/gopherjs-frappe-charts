package main

import (
	"math/rand"
	"strconv"
	"time"

	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		"1", "2", "3",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			"light-blue",
			[]interface{}{25, 40, 30},
		),
		charts.NewDataset(
			"Another Set",
			"violet",
			[]interface{}{25, 50, -10},
		),
	}

	chartArgs := charts.NewScatterChartArgs("#chart", "", chartData, 250)
	temp := 4
	chart := chartArgs.Render()
	for i := 0; i < 100; i++ {
		go func(i interface{}) {
			val := rand.Intn(5) + 2*i.(int)
			println(i, "sleeping for", val)
			time.Sleep(time.Duration(val) * time.Second)
			newLabel := strconv.Itoa(temp)
			actionCond := rand.Intn(4)
			switch actionCond {
			case 0:
				chart.RemoveDataPoint(rand.Int())
			case 1:
				chart.PopDataPoint()
			case 2:
				chart.AppendDataPoint([]interface{}{rand.Intn(1000), rand.Intn(1000)}, newLabel)
			case 3:
				chart.AddDataPoint([]interface{}{rand.Intn(200), rand.Intn(200)}, newLabel, rand.Intn(200))
			}
			temp++
			println(temp)
		}(i)
	}

	println("is this async?")
}
