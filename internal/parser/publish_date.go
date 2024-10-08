package parser

import "time"

var layouts = []string{
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	"2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04:05 -0700",
}

func parsePublishDate(date *string) (time.Time, error) {
	if date == nil {
		return time.Now(), nil
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, *date)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, err
}
