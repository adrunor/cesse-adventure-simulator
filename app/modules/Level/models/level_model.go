package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Index                uint8
	RequiredExperience   uint16
	CumulativeExperience uint16
	GivenCharacteristic  uint8
}
