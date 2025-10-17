package localizer

import (
	"demo-todo-manager/pkg/logger"
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Localizer struct {
	locale *i18n.Localizer
}
type contextKey string

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files := []string{
		"/app/internal/lang/en.json",
		"/app/internal/lang/ru.json",
	}

	for _, file := range files {
		_, err := bundle.LoadMessageFile(file)
		if err != nil {
			logger.Log.WithField("err", err).Fatal("Error starting server: loading location files failed")
		}
	}
}

func GetContextKey() contextKey {
	const ctxKey contextKey = "locale"
	return ctxKey
}

func New(locale language.Tag) *Localizer {
	return &Localizer{
		locale: i18n.NewLocalizer(bundle, locale.String()),
	}
}

func (l *Localizer) T(messageID string, template map[string]interface{}) string {
	msg, err := l.locale.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: template,
	})
	if err != nil {
		return messageID
	}

	return msg
}
