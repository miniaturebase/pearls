package progress

import pearls "github.com/minibase-app/pearls/types"

// Create a new progress bar instance
func Bar(percent float64) pearls.Progress {
	return pearls.Progress{
		Percent: percent,
	}
}
