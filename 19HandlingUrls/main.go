package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://onedrive.live.com/redir?resid=53E95F3256831099%21110&page=Edit&wd=target%28Go%20lang.one%7C40c15ce2-1e95-4953-982b-7ab781be85ab%2FUntitled%20Page%7C9237df09-5bdd-4e36-8071-27586ab0497d%2F%29&wdorigin=NavigationUrl"

func main() {
	fmt.Println("Priyansh is learning the handling the urls")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println("Printing scheme " + result.Scheme)
	fmt.Println("Printing host " + result.Host)
	fmt.Println("Printing path " + result.Path)
	fmt.Println("Printing port " + result.Port())
	fmt.Println("Printing RawQuery " + result.RawQuery)

	qparam := result.Query()

	fmt.Printf("The type of the query parameter %T\n", qparam)

	fmt.Println(qparam)
}
