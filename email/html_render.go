package email

import (
	"bytes"
	"text/template"

	"github.com/christhianjesus/storicard-challenge/summarize"
)

func GenerateTransactionSummary(tmplName string, summary *summarize.Summary) (string, error) {
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		return "", err
	}

	var tplOutput bytes.Buffer
	if err = tmpl.Execute(&tplOutput, summary); err != nil {
		return "", err
	}

	return tplOutput.String(), nil
}
