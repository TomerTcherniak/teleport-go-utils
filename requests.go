package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadequest(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	val, ok := headers["Teleportfile"]
	if ok {
		log.Print("Teleportfile key header is present with value", val)
		if strings.Contains(r.Header.Get("Teleportfile"), "teleport2022") {
			log.Print("queue")
			fileBytes, err := ioutil.ReadFile("download.sh")
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(fileBytes)
		} else {
			log.Print("Nope!")
		}
	} else {
		log.Print("Content-Type key header is not present")
	}
}

func authserver(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	val, ok := headers["Teleporturl"]
	if ok {
		log.Print("TeleportURL key header is present with value", val)
		if strings.Contains(r.Header.Get("Teleporturl"), "teleport2022") {
			log.Print("queue")
			val, ok := os.LookupEnv("AUTH_SERVER")
			if !ok {
				log.Print("AUTH_SERVER  not set \n")
			} else {
				log.Print("AUTH_SERVER ENV is ", val)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			valdecode := []byte(string(b64.StdEncoding.EncodeToString([]byte(val))))
			w.Write(valdecode)
		} else {
			log.Print("Nope!")
		}
	} else {
		log.Print("Content-Type key header is not present")
	}
}

func capinhashes(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	val, ok := headers["Teleportca"]
	if ok {
		log.Print("TeleportCA key header is present with value", val)
		if strings.Contains(r.Header.Get("Teleportca"), "teleport2022") {
			log.Print("queue")
			val, ok := os.LookupEnv("CA_PIN_HASHES")
			if !ok {
				log.Print("CA_PIN_HASHES  not set \n")
			} else {
				log.Print("CA_PIN_HASHES ENV is ", val)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			valdecode := []byte(string(b64.StdEncoding.EncodeToString([]byte(val))))
			w.Write(valdecode)
		} else {
			log.Print("Nope!")
		}
	} else {
		log.Print("Content-Type key header is not present")
	}
}

func healthrequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["success"] = "True"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
