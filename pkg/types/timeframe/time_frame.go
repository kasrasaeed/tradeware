package timeframe

import (
	"fmt"
	"strconv"
)

type TimeFrame string

func TimeFrameToSecond(timeFrame TimeFrame) (int64, error) {
	switch timeFrame {
	case "d":
		return 24 * 60 * 60, nil
	case "3d":
		return 3 * 24 * 60 * 60, nil
	case "w":
		return 7 * 24 * 60 * 60, nil
	case "m":
		return 30 * 24 * 60 * 60, nil
	}
	val, err := strconv.ParseInt(string(timeFrame), 10, 64)
	if err != nil || val == 0 {
		return 0, fmt.Errorf("given timeframe is not correct")
	}
	return val * 60, nil
}
