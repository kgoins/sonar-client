package pageclient

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func selectionEmpty(s *goquery.Selection) bool {
	return s.Length() == 0
}

func isErrStatus(code int) bool {
	return code < http.StatusOK || code >= http.StatusBadRequest
}

func appendURL(base string, next string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	baseURL.Path = path.Join(baseURL.Path, next)
	return baseURL.String(), nil
}

func getDocFromPage(pageURL string) (*goquery.Document, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if isErrStatus(resp.StatusCode) {
		return nil, errors.New(resp.Status)
	}

	respContentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(respContentType, "text/html") {
		return nil, errors.New("Response was not HTML")
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
