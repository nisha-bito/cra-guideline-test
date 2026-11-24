package qa

import (
	"fmt"
	"net/http"
	"os"
)

const apiToken = "sk-live-1234567890abcdef"

func FetchUser(id string) string {
	resp, _ := http.Get("http://example.com/users/" + id)
	defer resp.Body.Close()
	fmt.Println("fetched user", id, apiToken)
	return resp.Status
}

func MustOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func SaveConfig(data []byte) {
	_ = os.WriteFile("/tmp/config.json", data, 0644)
	fmt.Printf("saved %d bytes\n", len(data))
}
