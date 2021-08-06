package translation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/cmd/server/util"
	"time"
)

type TranslationData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	IsTl      bool   `json:"is_tl"`
}

type IDatedTranslation struct {
	Name       string             `json:"name" bson:"name"`
	Timestamp  primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Original   string             `json:"original" bson:"original"`
	Translated string             `json:"translated" bson:"translated"`
}

type IAnnouncingTranslation struct {
	Name       string    `json:"name" bson:"name"`
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	Original   string    `json:"original" bson:"original"`
	Translated string    `json:"translated" bson:"translated,omitempty"`
}

func ConvertDatedTranslationsToAnnouncingTranslations(translations []IDatedTranslation) []IAnnouncingTranslation {
	var convertedTranslations = make([]IAnnouncingTranslation, 0)
	for _, translation := range translations {
		var converted IAnnouncingTranslation
		converted.Translated = translation.Translated
		converted.Original = translation.Original
		converted.Timestamp = time.Unix(int64(translation.Timestamp)/1000, 0)
		converted.Name = translation.Name
		convertedTranslations = append(convertedTranslations, converted)
	}
	return convertedTranslations
}

func ConvertTranslationsToDatedTranslations(translations []TranslationData) []IDatedTranslation {
	var convertedTranslations = make([]IDatedTranslation, 0)
	for _, translation := range translations {
		var converted IDatedTranslation
		converted.Translated = ""
		converted.Original = translation.Message
		convertedDate, err := util.ConvertTimestampToPrimitiveDate(translation.Timestamp)
		if err != nil {
			continue
		}
		converted.Timestamp = convertedDate
		converted.Name = translation.Name
		convertedTranslations = append(convertedTranslations, converted)
	}
	return convertedTranslations
}
