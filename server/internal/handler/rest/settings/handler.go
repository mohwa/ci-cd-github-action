package settings

import (
	"github.com/mohwa/ci-cd-github-action/api/rest/model"
	"github.com/mohwa/ci-cd-github-action/internal/db"
)

func GetSettings() (*model.Settings, error) {
	settings, err := db.GetSettings()
	if err != nil {
		return nil, err
	}

	return &model.Settings{BgColor: settings.BgColor}, nil
}

func SetSettings(settings model.Settings) error {
	err := db.SetSettings(&db.Settings{BgColor: settings.BgColor})
	if err != nil {
		return err
	}

	return nil
}
