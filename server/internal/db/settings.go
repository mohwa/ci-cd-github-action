package db

import (
	"fmt"

	"gorm.io/gorm"
)

const settingsTableName = "settings"

func SettingsTable() (*gorm.DB, error) {
	if !isExistTable(settingsTableName) {
		return nil, fmt.Errorf("%w: %s", ErrorNoTable, settingsTableName)
	}

	return db.Table(settingsTableName).Session(&gorm.Session{}), nil
}

type Settings struct {
	ID      int    `gorm:"primaryKey"`
	BgColor string `gorm:"type:varchar(10); column:bgColor"`
}

func GetSettings() (*Settings, error) {
	table, err := SettingsTable()
	if err != nil {
		return nil, ErrorTableOpenFail
	}

	var ret Settings

	result := table.First(&ret)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ret, nil
}

func SetSettings(settings *Settings) error {
	table, err := SettingsTable()
	if err != nil {
		return ErrorTableOpenFail
	}

	var ret Settings

	if firstResult := table.First(&ret); firstResult.Error != nil {
		return firstResult.Error
	}

	var result *gorm.DB
	if len(ret.BgColor) == 0 {
		result = table.Create(&settings)
	} else {
		result = table.Where("id = ?", ret.ID).Updates(&settings)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
