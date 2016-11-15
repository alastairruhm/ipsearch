package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("param needed, an ip or url maybe")
		os.Exit(1)
	} else if len(os.Args) > 3 {
		fmt.Println("param number wrong, only one is acceptable")
		os.Exit(1)
	}
	ip := os.Args[1]
	resp, err := http.Get("http://freeapi.ipip.net/" + ip)
	if err != nil {
		fmt.Printf("http client error %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := extractData(string(body))
	fmt.Println(strings.Join(s, "-"))
}

// trimSlice trim cutset of element in slice
func trimSlice(src []string, cutset string) []string {
	var dest []string
	for _, s := range src {
		if strings.Trim(s, cutset) != "" {
			dest = append(dest, strings.Trim(s, cutset))
		}
	}
	return dest
}

// extractData trim leading and trailing "[" and "]"
func extractData(src string) []string {
	dest := strings.Trim(src, "[]")
	destSlice := trimSlice(strings.Split(dest, ","), "\"")
	return destSlice
}
