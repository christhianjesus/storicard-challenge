package email

import (
	"testing"

	"github.com/christhianjesus/storicard-challenge/summarize"
	"github.com/stretchr/testify/require"
)

func TestFileNotFound_GenerateTransactionSummary(t *testing.T) {
	_, err := GenerateTransactionSummary("", nil)

	require.Error(t, err)
}

func TestParseError_GenerateTransactionSummary(t *testing.T) {
	_, err := GenerateTransactionSummary("template.html", nil)

	require.Error(t, err)
}

func TestGenerateTransactionSummary(t *testing.T) {
	_, err := GenerateTransactionSummary("template.html", &summarize.Summary{})

	require.NoError(t, err)
}
