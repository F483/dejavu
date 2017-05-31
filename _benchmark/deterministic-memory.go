// +build ignore

package main

import (
	"github.com/f483/dejavu"
	"github.com/wcharczuk/go-chart"
	"net/http"
	"runtime"
	"strconv"
)

func runBenchmark(size uint) float64 {

	// max out
	d := dejavu.NewDeterministic(size)
	for i := 0; uint(i) < size; i++ {
		s := strconv.FormatInt(int64(i), 10)
		d.Witness([]byte(s))
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.Alloc) / float64(1024*1024)
}

func runBenchmarks() ([]float64, []float64) {
	return []float64{65536, 131072, 262144, 524288, 1048576}, []float64{
		runBenchmark(65536),
		runBenchmark(131072),
		runBenchmark(262144),
		runBenchmark(524288),
		runBenchmark(1048576),
	}
}

func drawChart(res http.ResponseWriter, req *http.Request) {

	xValues, yValues := runBenchmarks()

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "entries stored",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "memory usage in mb",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
