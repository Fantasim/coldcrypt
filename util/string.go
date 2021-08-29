package util

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func RemoveMultipleSpaceAndTrim(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(input, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")
	return final
}

func RemoveMultipleEndline(str string) string {
	re_inside_endl := regexp.MustCompile(`[\n]{2,}`)
	return re_inside_endl.ReplaceAllString(str, "\n")
}

func EndlineToSpace(str string) string {
	return strings.Replace(str, "\n", " ", -1)
}

func JSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func FormatInt(number int64) string {
	output := strconv.FormatInt(number, 10)
	startOffset := 3
	if number < 0 {
		startOffset++
	}
	for outputIndex := len(output); outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
	}
	return output
}

func FormatBigNumber(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
