package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main()  {
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		for key, values := range request.Header {
			key := "req-" + key
			for _, value := range values {
				writer.Header().Add(key, value)
			}
		}
		version, _ := os.LookupEnv("VERSION")
		writer.Header().Set("server-version", version)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello World"))

		log.Printf("%s %s", request.Host, http.StatusOK)
	})
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalln(err)
	}
}
