package parser

import "time"

func ParseTime(rawTime string) (*time.Time, error) {
	var temp *time.Time
	tempParsed, err := time.Parse(time.RFC1123, rawTime) // ex on javascript: new Date("2023-09-3").toUTCString()
	if err != nil {
		return nil, err
	}
	temp = &tempParsed

	return temp, nil
}
