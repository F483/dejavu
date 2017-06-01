// +build ignore

package main

import (
	"github.com/f483/dejavu"
	"github.com/wcharczuk/go-chart"
	"net/http"
	"strconv"
	"time"
)

func runBenchmark(size uint32) float64 {

	begin := time.Now().UnixNano()
	d := dejavu.NewDeterministic(size)
	for i := 0; i < 1000000; i++ {
		if i%10 == 0 {
			s := strconv.FormatInt(int64(i-10), 10)
			d.Witness([]byte(s))
		} else {
			s := strconv.FormatInt(int64(i), 10)
			d.Witness([]byte(s))
		}
	}
	end := time.Now().UnixNano()

	return float64(end-begin) / 1000000000.0
}

func runBenchmarks() ([]float64, []float64) {
	return []float64{4096, 8192, 16384, 32768, 65536}, []float64{
		runBenchmark(4096),
		runBenchmark(8192),
		runBenchmark(16384),
		runBenchmark(32768),
		runBenchmark(65536),
	}
}

func drawChart(res http.ResponseWriter, req *http.Request) {

	xValues, yValues := runBenchmarks()

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "max entrie size",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "seconds to witness 1000000 entries",
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
