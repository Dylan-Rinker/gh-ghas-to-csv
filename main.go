package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/dylan-rinker/gh-ghas-to-csv/cmd"
)

func main() {

	// Instantiate and execute root command
	cmd := cmd.NewCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	fmt.Println("hi world, this is the gh-ghas-to-csv extension!")
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
