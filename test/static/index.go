package main

import (
	"math/rand"
	"strconv"
	"time"

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
			[]float64{25, 40, 30},
		),
		charts.NewDataset(
			"Another Set",
			[]float64{25, 50, -10},
		),
	}

	chart := charts.NewScatterChart("#chart", chartData).
		WithHeight(250).
		WithColors([]string{"red", "green"}).
		Render()
	temp := 4
	for i := 0; i < 100; i++ {
		go func(i interface{}) {
			val := rand.Intn(3) + 2*i.(int)
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
				chart.AppendDataPoint([]float64{float64(rand.Intn(1000)), float64(rand.Intn(1000))}, newLabel)
			case 3:
				chart.AddDataPoint([]float64{float64(rand.Intn(200)), float64(rand.Intn(200))}, newLabel, rand.Intn(200))
			}
			temp++
			println(temp)
		}(i)
	}

	println("is this async?")
}
