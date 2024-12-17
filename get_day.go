package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	numAvailableDays := 17
	sessionToken := os.Getenv("aoc_session_token")

	for i := 1; i <= numAvailableDays; i++ {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", i), nil)
		req.Header.Set("Cookie", "session="+sessionToken)

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

		err = os.MkdirAll(fmt.Sprintf("day_%d", i), 0755)
		os.Chdir(fmt.Sprintf("day_%d", i))

		if err != nil {
			panic(err)
		}

		file, err := os.Create("input.txt")
		defer file.Close()

		file.Write(body)
		file.Sync()
		os.Chdir("..")
	}
}
