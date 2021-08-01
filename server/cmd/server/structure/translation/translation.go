package translation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
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
	Translated string             `json:"translated" bson:"translated,omitempty"`
}

func ConvertTranslationsToDatedTranslations(translations []TranslationData) []IDatedTranslation {
	var convertedTranslations = make([]IDatedTranslation, 0)
	for _, translation := range translations {
		var converted IDatedTranslation
		intTime, err := strconv.ParseInt("1405544146", 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		converted.Translated = ""
		converted.Original = translation.Message
		converted.Timestamp = primitive.NewDateTimeFromTime(time.Unix(intTime, 0))
		converted.Name = translation.Name
		convertedTranslations = append(convertedTranslations, converted)
	}
	return convertedTranslations
}
