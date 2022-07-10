package main

func main() {
	func romanToInt("III")
}
/*func romanToInt(s string) int {
    sum := 0
    m := make(map[string]int)
        m["I"] = 1
        m["V"] = 5
        m["X"] = 10
        m["L"] = 50
        m["C"] = 100
        m["D"] = 500
        m["M"] = 1000

    for idx, val := range s {
        if s[idx] = "I" && s[idx+1] = "V" {
            sum = sum +4
        }
        else if s[idx] = "I" && s[idx+1] = "X" {
            sum = sum +9
        }
        else if s[idx] = "X" && s[idx+1] = "L" {
            sum = sum +40
        }
        else if s[idx] = "X" && s[idx+1] = "C" {
            sum = sum +90
        }
        else if s[idx] = "C" && s[idx+1] = "D" {
            sum = sum +400
        }
        else if s[idx] = "C" && s[idx+1] = "M" {
            sum = sum +900
        }
        sum = sum + m[string(val)]
    }
    return sum
}*/
var romanInts = map[byte]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var exceptions = map[string]int {
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	n := len(s)
	result := 0
	i := 0

	for ; i < n - 1; i++ {
		if val, found := exceptions[s[i:i+2]]; found {
			result += val
			i++
		} else {
			result += romanInts[s[i]]
		}
	}

	if i == n - 1 {
		result += romanInts[s[i]]
	}

	return result
}