package pageclient

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ListStudies returns a list of studies
// currently available from OpenData
func ListStudies() ([]string, error) {
	pageBody, err := getDocFromPage(BaseURL)
	if err != nil {
		return nil, err
	}

	studyLinks := pageBody.Find(".pageContent .grid-container .cell a")
	if selectionEmpty(studyLinks) {
		return nil, errors.New("Unable to locate studies")
	}

	studyNames := make([]string, 0, studyLinks.Length())
	studyLinks.Each(func(i int, entry *goquery.Selection) {
		link, exists := entry.Attr("href")
		if !exists {
			return
		}

		link = strings.Trim(link, "/")
		studyNames = append(studyNames, link)
	})

	return studyNames, nil
}
