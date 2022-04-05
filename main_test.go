package goamputate

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var unAmpedUrl = "https://electrek-co.cdn.ampproject.org/c/s/electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/amp/"

var expectedUrls = []string{
	"https://electrek.co/2018/06/19/tesla-model-3-assembly-line-inside-tent-elon-musk/",
}

var mockedAmpResponse = []byte(
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

func TestAmputate(t *testing.T) {
	urls, err := Amputate([]string{unAmpedUrl}, map[string]string{})
	if err != nil {
		t.Errorf("Unable to Amputate: %v", err)
	}

	if !reflect.DeepEqual(urls, expectedUrls) {
		t.Errorf("Amputate did not return expected URL. Received: %v Expected: %v", urls, expectedUrls)
		os.Exit(1)
	}
}

func TestConvert(t *testing.T) {
	request := AmputationRequest{
		options: map[string]string{
			"gac": "true",
			"md":  "3",
		},
		urls: []string{unAmpedUrl},
	}
	response, err := Convert(request)
	if err != nil {
		t.Errorf("Error calling Convert: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(response))
}

func TestGetCanonicalUrls(t *testing.T) {
	urls, err := GetCanonicalUrls(mockedAmpResponse)
	if err != nil {
		t.Errorf("Unable to GetCanonicalUrls: %v", err)
		os.Exit(1)
	}

	if !reflect.DeepEqual(urls, expectedUrls) {
		t.Errorf("GetCanonicalUrls did not return expected object. Received: %v Expected: %v", urls, expectedUrls)
		os.Exit(1)
	}
}
