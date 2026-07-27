package main

import (
	"encoding/json"
	"io"
	"os"
	"net/http"
	"strings"
)

func main() {
	resp, _ := http.Get("https://books.toscrape.com/catalogue/page-1.html")
	body, _ := io.ReadAll(resp.Body)
	count := strings.Count(string(body), "product_pod")
	data, _ := json.MarshalIndent([]map[string]interface{}{{"count": count}}, "", " ")
	os.WriteFile("data/Go.json", data, 0644)
	println(count)
}