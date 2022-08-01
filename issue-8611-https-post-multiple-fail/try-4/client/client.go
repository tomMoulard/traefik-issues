package main

import (
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var errors int

func request(id int, url string, body io.Reader) {
	logger := log.WithFields(log.Fields{
		"id": id,
	})

	logger.Infof("r: %d\n", id)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		errors++
		logger.WithFields(log.Fields{"error-count": errors}).Error(err.Error())
	}
	logger.Debug("request created")

	client := &http.Client{}

	logger.Debug("client.do")
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("sending error: %v", err)
		return
	}
	logger.Debug("done")

	defer resp.Body.Close()
}

func main() {
	log.SetLevel(log.ErrorLevel)

	// file, _ := os.Open("/dev/random")
	for i := 0; i < 1000; i++ {
		go request(i, "http://localhost:8000", strings.NewReader("file"))
	}

	request(100, "http://localhost:8000", strings.NewReader("file"))
}
