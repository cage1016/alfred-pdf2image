package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var PageRangeRegex = regexp.MustCompile(`(?m)^\d*(?:-\d*)?(?:,\s*\d*(?:-\d*)?)*$`)

func IsPageRangeValid(s string) bool {
	return PageRangeRegex.MatchString(s)
}

type Range struct {
	Start int
	End   int
}

func ParsePageNumber(input string, maxNumPage int) (*[]Range, error) {
	groups := strings.Split(input, ",")
	ranges := make([]Range, len(groups))
	for i, group := range groups {
		parts := strings.Split(group, "-")
		if len(parts) == 1 {
			ranges[i].Start, _ = strconv.Atoi(parts[0])
			ranges[i].End = ranges[i].Start

			if ranges[i].Start > maxNumPage {
				return nil, fmt.Errorf("page range \"%s\" is out of total page \"%d\"", group, maxNumPage)
			}
		} else {
			if s, err := strconv.Atoi(parts[0]); err != nil {
				ranges[i].Start = 1
			} else {
				ranges[i].Start = s
			}

			if e, err := strconv.Atoi(parts[1]); err != nil {
				ranges[i].End = maxNumPage
			} else {
				ranges[i].End = e
			}

			if ranges[i].Start > ranges[i].End {
				return nil, fmt.Errorf("page range start \"%d\" must be less or equal to end \"%d\"", ranges[i].Start, ranges[i].End)
			}
		}
	}

	return &ranges, nil
}
