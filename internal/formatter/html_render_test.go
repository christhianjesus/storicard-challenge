package formatter

import (
	"testing"

	"github.com/christhianjesus/storicard-challenge/internal/summarize"
	"github.com/stretchr/testify/require"
)

func TestFileNotFound_GenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("", nil)

	require.Error(t, err)
}

func TestParseError_GenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("../../assets/template.html", nil)

	require.Error(t, err)
}

func TestGenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("../../assets/template.html", &summarize.Summary{})

	require.NoError(t, err)
}
