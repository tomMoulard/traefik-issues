package main

import (
	"net/http"

	"github.com/go-chi/chi"
	chitracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithServiceName("go-agent-service"),
		// tracer.WithEnv("staging"),
		tracer.WithAgentAddr("datadog:8126"),
	)

	defer tracer.Stop()

	// build chi router
	router := chi.NewRouter()
	router.Use(chitracer.Middleware(chitracer.WithServiceName("go-agent-service")))
	router.Get("/", myRequestHandler)
	http.ListenAndServe(":8080", router)
}

func myRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Chi's all good"))
}
