package shorten

import (
	"errors"
	"net/url"
	"strings"
)

type urlValidator struct {
	path string
}

func (checker urlValidator) check() (string, error) {
	unescapedUrl, _ := url.PathUnescape(checker.path)
	if !strings.Contains(unescapedUrl, ".") {
		return "", errors.New("not a valid url")
	}

	unescapedUrl = addHttpIfNotExists(unescapedUrl)
	data, err := url.Parse(unescapedUrl)

	if err != nil {
		return "", err
	} else if (data.Scheme != "http" && data.Scheme != "https") || data.Host == ""  {
		return "", errors.New("not a valid url")
	} else {
		return unescapedUrl, nil
	}
}

func addHttpIfNotExists(url string) string{
	if !strings.Contains(url, "http") && !strings.Contains(url, "https"){
		url = "http://" + url
	}
	return url
}