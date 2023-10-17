package email

import (
	"testing"

	"github.com/christhianjesus/storicard-challenge/summarize"
	"github.com/stretchr/testify/require"
)

func TestFileNotFound_GenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("", nil)

	require.Error(t, err)
}

func TestParseError_GenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("template.html", nil)

	require.Error(t, err)
}

func TestGenerateTransactionEmail(t *testing.T) {
	_, err := GenerateTransactionEmail("template.html", &summarize.Summary{})

	require.NoError(t, err)
}
