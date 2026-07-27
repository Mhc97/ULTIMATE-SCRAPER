package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
	
)

type Result struct {
	Lang      string `json:"lang"`
	Count     int    `json:"count"`
	MemoryMB int    `json:"memory_mb"`
	Duration  string `json:"duration"`
	Error string `json:"error,omitempty"`
}

var (
	results []Result
	mu     sync.Mutex
)

func main() {
	fmt.Printf("🚀 Ultimate Scraper v2 (6 langages)\n🧠 CPU: %d\n\n", runtime.NumCPU())

	http.HandleFunc("/scrape", scrapeHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/", homeHandler)

	fmt.Println("🌐 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		http.Error(w, "POST required", 405)
		return
	}
	results = []Result{}
	var wg sync.WaitGroup

	scrappers := []struct {
		name, cmd, script string
		mem				  int
	}{	
		{"Go", "go", "run scraper_go.go", 15},
		{"Rust", "./scraper_rust", "", 8},
		{"C++", "./scraper_cpp", "", 5},
		{"Python", "python", "scraper_py.py", 120},
		{"Node", "node", "scraper_js.js", 100},
		{"PHP", "php", "scraper_php.php", 80},
	}

	for _, s:= range scrappers {
		wg.Add(1)
		go func(s struct{ name, cmd, script string; mem int }){
			defer wg.Done()
			start := time.Now()
			cmd := exec.Command(s.cmd, s.script)
			out, err := cmd.CombinedOutput()

			res := Result{Lang: s.name, MemoryMB: s.mem, Duration: time.Since(start).String()}

			if err != nil {
				res.Error = string(out)
			}else {
				data, _ := os.ReadFile("data/" + s.name + ".json")
				var books []interface{}
				json.Unmarshal(data, &books)
				res.Count = len(books)
			}
				mu.Lock()
				results = append(results, res)
				mu.Unlock()
			}(s)
		}
		wg.Wait()
		json.NewEncoder(w).Encode(results)
	}

	func metricsHandler(w http.ResponseWriter, r *http.Request){
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"alloc_mb": m.Alloc / 1024 / 1024,
			"sys_mb": m.Sys / 1024 / 1024,
			"num_gc": m.NumGC,
			"goroutine": runtime.NumGoroutine(),
		})
	}

	func homeHandler(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "POST /scrape\nGET /metrics")
	}



