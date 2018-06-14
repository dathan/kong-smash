package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// StructuredResponse common response which contains a map payload
type StructuredResponse struct {
	Status  string                 `json:"status"`
	Msg     string                 `json:"msg"`
	Payload map[string]interface{} `json:"payload"`
}

func main() {

	http.HandleFunc("/simulate", SimulatePause)
	//http.HandleFunc("/smashme", SimulatePause)
	error := http.ListenAndServe(":8282", nil)
	if error != nil {
		fmt.Printf("ERROR : %s", error.Error())
	}

}

// SimulatePause is a handler for the service
func SimulatePause(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	switch r.Method {
	case "GET":
		js := StructuredResponse{
			Status: "success",
		}

		delay := r.FormValue("delay")
		duration := time.Millisecond * 100
		if len(delay) > 1 {
			dvalue, err := strconv.ParseInt(delay, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			if dvalue >= 0 {
				duration = time.Millisecond * time.Duration(dvalue)
			}
		}

		time.Sleep(duration)
		dur := timeTrack(t0, "SimulatePause")
		js.Payload = make(map[string]interface{})
		js.Payload["duration"] = dur

		jbyte, err := json.Marshal(js)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jbyte)

		return

	default:
		http.Error(w, "UNKNOWN METHOD", http.StatusInternalServerError)
		return
	}
	//return

}

func timeTrack(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	return elapsed
}
