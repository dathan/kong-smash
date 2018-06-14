package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dathan/casync"
)

func main() {
	var tasks []*casync.Task
	t0 := time.Now()
	requests := flag.Int("requests", 1000, "how many requests to perform")
	duration := flag.Int("delay", 0, "add a delay flag to cause the service to block in ms time")
	URL := flag.String("url", "http://localhost:8282/simulate", "url to smash")
	concurrency := flag.Int("concurrency", 4, "how many concurrent processes to run. Raise this to smash more")

	flag.Parse()
	// set up the number of tasks
	for i := 0; i < *requests; i++ {
		tsk := casync.NewTask(i, queryTask(*URL, int64(*duration)))
		tasks = append(tasks, tsk)
	}
	// set up the concurrency
	as := casync.NewAsync(*concurrency, tasks)
	as.ExecuteTasks()
	timeTrack(t0, "Execution Finished")
	// gather timings
	// write to file
	// graph
}

func queryTask(url string, d int64) func() {
	return func() {
		//setup client
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}

		if d > 0 {
			url = fmt.Sprintf("%s?delay=%d", url, d)
		}

		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		var f map[string]interface{} // generic type
		json.Unmarshal(body, &f)     // have to read the body to work around a golang bug and reuse ports
		if f == nil || f["status"] == nil || f["status"].(string) != "success" {

			if msg, ok := f["message"].(string); ok {
				fmt.Printf("ERROR %s\n", msg)
			}

			return
		}

		payload := f["payload"].(map[string]interface{})
		if v, ok := payload["duration"]; ok {
			durnano := v.(float64)
			fl := durnano / 1000000
			log.Printf("Duration: %.2f\n", fl)
		}
		return

	}
}

func timeTrack(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	return elapsed
}
