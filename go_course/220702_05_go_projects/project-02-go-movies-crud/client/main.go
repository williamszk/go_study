package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func main() {

	// following a simple tutorial on how to include insert wait
	// in the terminal
	// https://stackoverflow.com/questions/18936662/how-to-wait-for-command-line-input-in-go

	fmt.Println("Choose an option: get, get all, delete")

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	command, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("\n")

	commStr := string(command)

	if runtime.GOOS == "windows" {
		commStr = commStr[:len(commStr)-2]
	} else if runtime.GOOS == "linux" {
		commStr = commStr[:len(commStr)-1]
	} else {
		fmt.Errorf("We could not identify the OS. Stop the program here.")
		return
	}

	if commStr == "get all" {

		func() {
			// ==================================== //
			// GET request for all movies
			// ==================================== //
			url := "http://localhost:8000/movies"

			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Response from get all movies")
			fmt.Println(string(body))
		}()

	} else if commStr == "get" {

		func() {
			// ==================================== //
			// GET request for movie with id=1
			// ==================================== //
			url := "http://localhost:8000/movies/1"

			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Response from get movie with id=1")
			fmt.Println(string(body))
		}()

		func() {
			// ==================================== //
			// GET request for movie with id=2 (alternative implementation)
			// ==================================== //
			url := "http://localhost:8000/movies/2"

			// create a request
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("There is a problem with the request:", err)
			}

			// fetch request
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println("There is a problem with the response:", err)
			}
			defer res.Body.Close()

			// read response body
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("There is a problem with read the body:", err)
			}

			fmt.Println("Response from GET (alternative implementation) movie with id=2")
			fmt.Println(string(body))

		}()

	} else if commStr == "delete" {

		func() {
			// ==================================== //
			// DELETE request id=1
			// ==================================== //
			url := "http://localhost:8000/movies/1"

			// create a request
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				fmt.Println("There is a problem with the request:", err)
			}

			// fetch request
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println("There is a problem with the response:", err)
			}
			defer res.Body.Close()

			// read response body
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("There is a problem with read the body:", err)
			}

			fmt.Println("Response from DELETE movie with id=1")
			fmt.Println(string(body))
		}()

	} else {
		fmt.Println("This is an invalid command")
		fmt.Printf("Unrecognized command: '%s'\n", commStr)
	}
}
