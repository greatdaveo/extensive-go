package mylib

import (
	"fmt"
	"net/url"
)

func UrlParsing() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)
	fmt.Println("Fragment: ", parsedURL.Fragment)

	fmt.Println("-----------------------------------------------")

	rawUrl1 := "https://example.com/path?name=john&age=30"
	parsedURL1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	fmt.Println("-----------------------------------------------")

	// Building URL
	basedUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := basedUrl.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	basedUrl.RawQuery = query.Encode()
	fmt.Println("Built URL: ", basedUrl.String())

	values := url.Values{}
	// Add key-value pairs to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	fmt.Println("-----------------------------------------------")
	basedUrl1 := "https://example.com/search"
	fullURL := basedUrl1 + "?" + encodedQuery
	fmt.Println(fullURL)

}
