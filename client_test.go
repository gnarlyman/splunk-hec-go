package hec

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testSplunkURL   = "https://localhost:8088"
	testSplunkToken = "00000000-0000-0000-0000-000000000000"
)

var insecureClient *http.Client = &http.Client{Transport: &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}}

func TestHEC_WriteEvent(t *testing.T) {
	event := &Event{
		Index:      String("main"),
		Source:     String("test-hec-raw"),
		SourceType: String("manual"),
		Host:       String("localhost"),
		Time:       String("1485237827.123"),
		Event:      "hello, world",
	}

	c := NewClient(testSplunkURL, testSplunkToken)
	c.SetHTTPClient(insecureClient)
	err := c.WriteEvent(event)
	assert.NoError(t, err)
}

func TestHEC_WriteObjectEvent(t *testing.T) {
	event := &Event{
		Index:      String("main"),
		Source:     String("test-hec-raw"),
		SourceType: String("manual"),
		Host:       String("localhost"),
		Time:       String("1485237827.123"),
		Event: map[string]interface{}{
			"str":  "hello",
			"time": time.Now(),
		},
	}

	c := NewClient(testSplunkURL, testSplunkToken)
	c.SetHTTPClient(insecureClient)
	err := c.WriteEvent(event)
	assert.NoError(t, err)
}

func TestHEC_WriteEventBatch(t *testing.T) {
	events := []*Event{
		{Event: "event one"},
		{Event: "event two"},
	}

	c := NewClient(testSplunkURL, testSplunkToken)
	c.SetHTTPClient(insecureClient)
	err := c.WriteBatch(events)
	assert.NoError(t, err)
}

func TestHEC_WriteEventRaw(t *testing.T) {
	events := `2017-01-24T06:07:10.488Z Raw event one
2017-01-24T06:07:12.434Z Raw event two`
	metadata := EventMetadata{
		Source: String("test-hec-raw"),
	}
	c := NewClient(testSplunkURL, testSplunkToken)
	c.SetHTTPClient(insecureClient)
	err := c.WriteRaw([]byte(events), &metadata)
	assert.NoError(t, err)
}