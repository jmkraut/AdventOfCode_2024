package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	day := "1"
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day), nil)
	req.Header.Set("Cookie", "session=")

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(fmt.Sprintf("day_%s", day), 0755)
	os.Chdir(fmt.Sprintf("day_%s", day))

	if err != nil {
		panic(err)
	}

	file, err := os.Create("input.txt")
	defer file.Close()

	file.Write(body)
	file.Sync()
}
