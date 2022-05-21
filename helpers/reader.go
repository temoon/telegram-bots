package helpers

import (
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/temoon/telegram-bots/config"
)

func ReadTemplate(data interface{}, filename ...string) (text string, err error) {
	var tmplText string
	if tmplText, err = ReadStaticFile(filename...); err != nil {
		return
	}

	var tmpl *template.Template
	if tmpl, err = template.New("static").Parse(tmplText); err != nil {
		return
	}

	textBuilder := strings.Builder{}
	if err = tmpl.Execute(&textBuilder, data); err != nil {
		return
	}

	return textBuilder.String(), nil
}

func ReadStaticFile(filename ...string) (text string, err error) {
	var data []byte
	if data, err = ReadRawStaticFile(filename...); err != nil {
		return
	}

	return string(data), nil
}

func ReadRawStaticFile(filename ...string) (data []byte, err error) {
	filename = append([]string{config.GetBotStaticRoot()}, filename...)
	if data, err = os.ReadFile(path.Join(filename...)); err != nil {
		return
	}

	return
}
