package pageclient

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

// ListFiles returns a list of filenames provide on the
// sonar data site for a given study id
func ListFiles(studyID string) ([]string, error) {
	studyURL, err := appendURL(BaseURL, studyID)
	if err != nil {
		return nil, err
	}

	pageBody, err := getDocFromPage(studyURL)
	if err != nil {
		return nil, err
	}

	return getFileNamesFromPage(pageBody)
}

func getFileNamesFromPage(pageBody *goquery.Document) ([]string, error) {
	tableDiv := pageBody.Find("div .table-scroll").First()
	if selectionEmpty(tableDiv) {
		return nil, errors.New("Unable to locate content table")
	}

	tableContent := tableDiv.Find("table tbody tr")
	if selectionEmpty(tableContent) {
		return nil, errors.New("Unable to locate content rows")
	}

	fileNames := make([]string, 0, tableContent.Length())
	tableContent.Each(func(i int, entry *goquery.Selection) {
		fileLink := entry.Find("td a")
		if selectionEmpty(fileLink) {
			return
		}

		fileNames = append(fileNames, fileLink.Text())
	})

	return fileNames, nil
}
