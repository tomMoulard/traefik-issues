package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/containous/mux"
	log "github.com/sirupsen/logrus"
)

func handler(ch chan bool) func(http.ResponseWriter, *http.Request) {
	log.SetLevel(log.DebugLevel)
	logger := log.New()

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("new request")

		b, err := ioutil.ReadAll(r.Body)
		logger.WithFields(log.Fields{
			"host":    r.Host,
			"url":     r.URL,
			"headers": r.Header,
			"err":     err,
		}).Infof("%+v", string(b))
		defer r.Body.Close()

		logger.Debug("waiting for chan")
		<-ch
		logger.Debug("closing")
	}
}

func serve(ch chan bool) {
	r := mux.NewRouter()
	r.HandleFunc("/", handler(ch))
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	ch := make(chan bool)
	go serve(ch)

	for {

		time.Sleep(30 * time.Second)
		log.Info("closing handlers", time.Now().Format(time.RFC822))
		for i := 0; i < 100000; i++ {
			ch <- true
		}
	}
}
