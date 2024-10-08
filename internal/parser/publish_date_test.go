package parser

import "testing"

func TestParsePublishDate(t *testing.T) {
	testDates := []string{
		"2024-09-17T00:00:00+00:00",
		"Tue, 08 Oct 2024 13:00:00 GMT",
		"18 Apr 2024 00:00:00 +0000",
		"8 Apr 2024 00:00:00 +0000",
		"Thu, 8 Sep 2011 01:36:45 +0000",
	}

	for _, testDate := range testDates {
		_, err := parsePublishDate(&testDate)
		if err != nil {
			t.Error(err)
		}
	}
}
