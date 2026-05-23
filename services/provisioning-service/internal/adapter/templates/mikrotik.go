package templates

import (
	"bytes"
	"text/template"
)

type MikroTikConfig struct {
	Username string
	Password string
	Profile  string
}

const pppoeUserTemplate = \`
/ppp secret add name={{.Username}} password={{.Password}} profile={{.Profile}} service=pppoe
\`

func GenerateMikroTikPPPoE(cfg MikroTikConfig) (string, error) {
	tmpl, err := template.New("mikrotik").Parse(pppoeUserTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", err
	}

	return buf.String(), nil
}
