package util

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
	"time"
)

func ConvertTimestampToPrimitiveDate(timestamp interface{}) (convertedDate primitive.DateTime, err error) {
	var convertedIntTime int64
	switch typedTimestamp := timestamp.(type) {
	case int32:
	case string:
		convertedIntTime, err = strconv.ParseInt(typedTimestamp, 10, 64)
		if err != nil {
			log.Println(err)
			return convertedDate, err
		}
	case int64:
		convertedIntTime = typedTimestamp
	default:
		return convertedDate, errors.New("Invalid Type")
	}
	if err != nil {
		log.Println(err)
		return convertedDate, err
	}
	convertedDate = primitive.NewDateTimeFromTime(time.Unix(convertedIntTime/1000, 0))
	return convertedDate, nil
}
