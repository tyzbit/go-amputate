package goamputate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

const (
	amputatorApi string = "https://www.amputatorbot.com/api/v1"
	userAgent    string = "github.com/tyzbit/go-amputate"
)

func (AmputatorBot) Convert(r AmputationRequest) ([]byte, error) {
	url := fmt.Sprintf("%v/convert?gac=%v&md=%v&q=%v", amputatorApi, r.gac, r.md, strings.Join(r.urls, ";"))
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetCanonicalUrls(body []byte) ([]string, error) {
	urls := []string{}
	_, err := jsonparser.ArrayEach(body, func(amputateObject []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonparser.ArrayEach(amputateObject, func(canonical []byte, dataType jsonparser.ValueType, offset int, err error) {
			if is_amp, _ := jsonparser.GetBoolean(canonical, "is_amp"); !is_amp {
				if url, _ := jsonparser.GetString(canonical, "url"); url != "" {
					urls = append(urls, url)
				}
			}
		}, "canonicals")
	})
	if err != nil {
		return nil, err
	}

	uniqueUrls := _removeDuplicateValues(urls)
	return uniqueUrls, nil
}

func _removeDuplicateValues(strings []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strings {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
