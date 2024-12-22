package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	numAvailableDays := 22

	sessionToken := doesSessionTokenExist("aoc_session_token")

	for i := 1; i <= numAvailableDays; i++ {

		shouldSetupDay := shouldSetupDay(fmt.Sprintf("day_%d", i))

		if shouldSetupDay {
			log.Println(fmt.Sprintf("Setting up directory and input file for day %d", i))

			// Get each days input as a string
			dayInput := getDayInputFromWeb(sessionToken, i)

			// Create the days folder structure (e.g. /day_1)
			dirFilePath := createDayDirectory(i)

			//Create the input file for that day
			inputFilePath := fmt.Sprintf("%s/input.txt", dirFilePath)
			createDayInputFile(dayInput, inputFilePath)
		}
	}
}

func shouldSetupDay(filePath string) bool {
	_, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			panic(err)
		}
	}

	return false
}

func createDayDirectory(day int) string {
	filePath := fmt.Sprintf("day_%d", day)
	err := os.MkdirAll(filePath, 0755)

	if err != nil {
		panic(err)
	}

	return filePath
}

func createDayInputFile(body []byte, path string) {
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.Write(body)
	file.Sync()
}

func getDayInputFromWeb(sessionToken string, day int) []byte {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)

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

	return body
}

func doesSessionTokenExist(envSessionTokenName string) string {
	sessionToken, sessionTokenExists := os.LookupEnv(envSessionTokenName)
	if !sessionTokenExists || sessionToken == "" {
		log.Fatal("Session token wasn't set or was empty.")
	}

	return sessionToken
}
