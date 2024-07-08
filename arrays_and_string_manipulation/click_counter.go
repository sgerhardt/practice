package arrays_and_string_manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

// Given input of clicks and domains, count the total click for each subdomain

//e.g.
// 10,yahoo.sports.com
// 15,yahoo.com
// 2,com.com
//
// Outputs:
// 25 com
// 25 yahoo.com
// 10 yahoo.sports.com
// 2 com.com

func getClicks(input []string) map[string]string {

	domainToClicks := map[string]string{}
	for _, row := range input {
		splitRow := strings.Split(row, ",")
		clicks := splitRow[0]
		domain := splitRow[1]

		subdomains := strings.Split(domain, ".")
		previous := ""
		for i := len(subdomains) - 1; i >= 0; i-- {
			// build-up previous string
			joined := subdomains[i]
			if previous != "" {
				joined = strings.Join([]string{joined, previous}, ".")
			}
			val, ok := domainToClicks[joined]
			if ok {
				existingVal, err := strconv.Atoi(val)
				if err != nil {
					fmt.Printf("error pasring int:%+v\n", err)
					continue
				}
				clickVal, err := strconv.Atoi(clicks)
				if err != nil {
					fmt.Printf("error pasring int:%+v\n", err)
					continue
				}
				domainToClicks[joined] = strconv.Itoa(existingVal + clickVal)
			} else {
				domainToClicks[joined] = clicks
			}
			if previous == `` {
				previous = subdomains[i]
				continue
			}
			previous = subdomains[i] + `.` + previous
		}
	}

	return domainToClicks
}
