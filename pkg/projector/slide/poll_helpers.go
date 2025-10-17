package slide

import (
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/shopspring/decimal"
)

// formatDecimalAsString formats a decimal value as a string with a maximum precision of 3 decimal places.
// Trailing zeros are removed.
func formatDecimalAsString(value decimal.Decimal) string {
	return value.Round(3).String()
}

// calculatePercent calculates the percentage of a value relative to a base.
// Returns the percentage as a string with max 3 decimal places.
// If base is zero, returns an empty string.
func calculatePercent(value decimal.Decimal, base decimal.Decimal) string {
	if base.IsZero() {
		return ""
	}
	percent := value.Div(base).Mul(decimal.NewFromInt(100))
	return formatDecimalAsString(percent)
}

// shouldDisplayPercent determines whether percentage values should be displayed
// based on the poll's onehundredPercentBase and pollmethod settings.
// For vote type (Y/N/A), it checks if that vote type should be included in the percentage base.
func shouldDisplayPercent(poll dsmodels.Poll, voteType string) bool {
	base := poll.OnehundredPercentBase

	// If base is disabled, never show percentages
	if base == "disabled" {
		return false
	}

	// For cast and valid, all vote types are included
	if base == "cast" || base == "valid" {
		return true
	}

	// Check if the vote type is included in the base
	// The base can be Y, N, A, YN, YNA
	return strings.Contains(base, voteType)
}
