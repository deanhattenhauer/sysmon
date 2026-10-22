package collect

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Load struct {
	One     float64
	Five    float64
	Fifteen float64
}

func parseLoad(content string) (Load, error) {
	fields := strings.Fields(content)
	if len(fields) < 3 {
		return Load{}, fmt.Errorf("parsing loadavg: expected at least 3 fields, got %d", len(fields))
	}

	one, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return Load{}, fmt.Errorf("parsing loadavg 1 minute: %w", err)
	}

	five, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return Load{}, fmt.Errorf("parsing loadavg 5 minute: %w", err)
	}

	fifteen, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return Load{}, fmt.Errorf("parsing loadavg 15 minute: %w", err)
	}

	return Load{One: one, Five: five, Fifteen: fifteen}, nil
}

func ReadLoad() (Load, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return Load{}, fmt.Errorf("reading load: %w", err)
	}
	return parseLoad(string(data))
}
