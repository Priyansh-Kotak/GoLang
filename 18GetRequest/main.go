package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://onedrive.live.com/redir?resid=53E95F3256831099%21110&page=Edit&wd=target%28Go%20lang.one%7C40c15ce2-1e95-4953-982b-7ab781be85ab%2FUntitled%20Page%7C9237df09-5bdd-4e36-8071-27586ab0497d%2F%29&wdorigin=NavigationUrl"

func main() {

	fmt.Println("Priyansh is learning the get request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response is of type %T\n: ", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
