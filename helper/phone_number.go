package helper

import (
	"strings"
)

func BeautifyIDNumber(mdn string, zero bool) string {
	check := true

	for check {
		check = false

		// Remove non-numeric prefix
		if len(mdn) > 0 && !isNumeric(string(mdn[0])) {
			mdn = mdn[1:]
			check = true
		}

		// Remove '62' prefix
		if strings.HasPrefix(mdn, "62") {
			mdn = mdn[2:]
			check = true
		}

		// Remove leading '0's
		for strings.HasPrefix(mdn, "0") {
			mdn = mdn[1:]
			check = true
		}
	}

	if zero {
		mdn = "0" + mdn
	} else {
		mdn = "62" + mdn
	}

	return mdn
}

// isNumeric checks if a string is numeric
func isNumeric(str string) bool {
	return str >= "0" && str <= "9"
}
