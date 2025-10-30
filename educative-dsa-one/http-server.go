package educativedsa

import (
	"fmt"
	"net/http"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const form = `<html><body><form action="/" method="POST">
<h1>Statistics</h1>
<h5>Compute base statistics for a given list of numbers</h5>
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="number" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></body></html>`

const error = `<p class="error">%s</p>`

var pageBottom = ``

// Write an HTML header, parse the form, write form to writer and make request for numbers
func HomePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("mamamam", request.Method)

	switch request.Method {
	case "GET":
		fmt.Fprintf(writer, "%v", form)
	case "POST":

		pageBottom = FormatStats(statistics{})

		ProcessRequest(request)
		fmt.Fprintf(writer, "%v \n %v", form, pageBottom)
	}

}

// Capture the numbers from the request, and format the data and check for errors
func ProcessRequest(request *http.Request) (separatedInputs []float64, err string, isValid bool) {
	// write your code here

	inputsGiven := request.FormValue("numbers")

	if inputsGiven != "" {

		inputArr := strings.FieldsFunc(inputsGiven, func(r rune) bool {

			return r == ' ' || r == ','
		})

		fmt.Println("inputArr", inputArr)
	}

	fmt.Println("inputsGiven", inputsGiven)
	return nil, "", false
}

// sort the values to get mean and median
func GetStats(numbers []float64) (stats statistics) {
	// write your code here
	return stats
}

// seperate function to calculate the sum for mean
func sum(numbers []float64) (total float64) {
	// write your code here
	return 0
}

// seperate function to calculate the median
func median(numbers []float64) float64 {
	// write your code here
	return 0
}

func FormatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
