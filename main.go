package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/glanceapp/glance/pkg/sysinfo"
)

func WithAuth(next http.Handler, auth func(string) bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authValue := r.Header.Get("Authorization")
		if !auth(authValue) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GlanceAgentHandler(w http.ResponseWriter, r *http.Request) {
	all_nil := true
	info, err := sysinfo.Collect(nil)
	if len(err) > 0 {
		for _, e := range err {
			if e != nil {
				log.Printf("glance-agent error: %v\n", err)
				all_nil = false
			}
		}

		if !all_nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(info); err != nil {
		log.Printf("glance-agent error: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {
	var (
		host   string
		port   string
		secret string
		ok     bool
	)

	if host, ok = os.LookupEnv("GLANCE_AGENT_HOST"); !ok {
		host = ""
	}

	if port, ok = os.LookupEnv("GLANCE_AGENT_PORT"); !ok {
		port = "8080"
	}

	if secret, ok = os.LookupEnv("GLANCE_AGENT_SECRET"); !ok {
		log.Fatalf("Environment variable $GLANCE_AGENT_SECRET is not available")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/sysinfo/all", GlanceAgentHandler)
	handler := WithAuth(
		mux,
		func(maybeSecret string) bool {
			return fmt.Sprintf("Bearer %s", secret) == maybeSecret
		},
	)

	address := fmt.Sprintf("%s:%s", host, port)
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatal(err)
	}
}
