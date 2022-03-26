package goamputate

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	request := AmputationRequest{
		gac:  true,
		md:   3,
		urls: []string{"https://electrek-co.cdn.ampproject.org/c/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/"},
	}
	var bot AmputatorBot
	response, err := bot.Convert(request)
	if err != nil {
		t.Errorf("Error calling Convert: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(response))
}

func TestGetCanonicalUrls(t *testing.T) {
	response := []byte(
		`[
		  {
			"amp_canonical": {
			  "domain": "ampproject",
			  "is_alt": false,
			  "is_amp": true,
			  "is_cached": true,
			  "is_valid": true,
			  "type": "CANURL",
			  "url": "https://electrek-co.cdn.ampproject.org/c/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/?usqp=mq331AQIKAGwASCAAgM%3D",
			  "url_similarity": 0.7480314960629921
			},
			"canonical": {
			  "domain": "electrek",
			  "is_alt": false,
			  "is_amp": false,
			  "is_cached": null,
			  "is_valid": true,
			  "type": "REL",
			  "url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
			  "url_similarity": 0.8900523560209425
			},
			"canonicals": [
			  {
				"domain": "electrek",
				"is_alt": false,
				"is_amp": false,
				"is_cached": null,
				"is_valid": true,
				"type": "REL",
				"url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
				"url_similarity": 0.8900523560209425
			  },
			  {
				"domain": "electrek",
				"is_alt": false,
				"is_amp": false,
				"is_cached": null,
				"is_valid": true,
				"type": "REL",
				"url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
				"url_similarity": 0.8900523560209425
			  }
			],
			"origin": {
			  "domain": "google",
			  "is_amp": true,
			  "is_cached": true,
			  "is_valid": true,
			  "url": "https://www.google.com/amp/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/"
			}
		  },
		  {
			"amp_canonical": {
			  "domain": "ampproject",
			  "is_alt": false,
			  "is_amp": true,
			  "is_cached": true,
			  "is_valid": true,
			  "type": "GOOGLE_JS_REDIRECT",
			  "url": "https://electrek-co.cdn.ampproject.org/c/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/?usqp=mq331AQIKAGwASCAAgM%3D",
			  "url_similarity": 0.7480314960629921
			},
			"canonical": {
			  "domain": "electrek",
			  "is_alt": false,
			  "is_amp": false,
			  "is_cached": null,
			  "is_valid": true,
			  "type": "REL",
			  "url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
			  "url_similarity": 0.8900523560209425
			},
			"canonicals": [
			  {
				"domain": "electrek",
				"is_alt": false,
				"is_amp": false,
				"is_cached": null,
				"is_valid": true,
				"type": "REL",
				"url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
				"url_similarity": 0.8900523560209425
			  },
			  {
				"domain": "electrek",
				"is_alt": false,
				"is_amp": false,
				"is_cached": null,
				"is_valid": true,
				"type": "REL",
				"url": "https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
				"url_similarity": 0.8900523560209425
			  }
			],
			"origin": {
			  "domain": "google",
			  "is_amp": true,
			  "is_cached": true,
			  "is_valid": true,
			  "url": "https://www.google.com/amp/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/"
			}
		  }
		]`)

	urls, err := GetCanonicalUrls(response)
	if err != nil {
		t.Errorf("Unable to GetCanonicalUrls: %v", err)
		os.Exit(1)
	}

	expectedUrls := []string{
		"https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
	}
	if !reflect.DeepEqual(urls, expectedUrls) {
		t.Errorf("GetCanonicalUrls did not return expected object. Received: %v Expected: %v", urls, expectedUrls)
		os.Exit(1)
	}
}
