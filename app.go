package main


import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Let's first read the `config.json` file
	handlerdownload := http.HandlerFunc(downloadequest)
	http.Handle("/download", handlerdownload)
	handlauthserver := http.HandlerFunc(authserver)
	http.Handle("/auth_server", handlauthserver)
	handlcapinhashes := http.HandlerFunc(capinhashes)
	http.Handle("/ca_pin_hashes", handlcapinhashes)
	handlerhealth := http.HandlerFunc(healthrequest)
	http.Handle("/health", handlerhealth)
	val, ok := os.LookupEnv("PORT")
	if !ok {
		log.Print("PORT ENV not set \n")
	} else {
		log.Print("PORT ENV is ", val)
		port := val
		log.Print("Starting Server on port ", port)
		http.ListenAndServe(":"+port, nil)
	}

}
