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
		// STRING INPUT IS FROM HOLODEX IN MILISECOND. THEREFOR *10^6 FOR NANOUNIX STANDARD IN SYSTEM
		convertedIntTime = convertedIntTime * 1000000
	case int64:
		convertedIntTime = typedTimestamp
	default:
		return convertedDate, errors.New("Invalid Type")
	}
	if err != nil {
		log.Println(err)
		return convertedDate, err
	}

	convertedDate = primitive.NewDateTimeFromTime(time.Unix(0, convertedIntTime))
	return convertedDate, nil
}

func ConvertTimeStringToDate(timestamp string)(primitive.DateTime,error){
	tempTime, err := time.Parse(time.RFC3339, timestamp)
	convertedTime := primitive.NewDateTimeFromTime(tempTime)
	return convertedTime, err
}