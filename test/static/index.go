package main

import (
	"math/rand"
	"strconv"

	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	data := make(map[string]interface{})
	currentTimestamp := 1477699200
	for i := 0; i < 375; i++ {
		data[strconv.Itoa(currentTimestamp)] = rand.Intn(10)
		currentTimestamp += 86400
	}
	chartArgs := charts.NewHeatmapArgs("#chart", data, 115)
	// Halloween colors
	// chartArgs.LegendColors = []string{"#ebedf0", "#fdf436", "#ffc700", "#ff9100", "#06001c"}
	chartArgs.Render()
}
