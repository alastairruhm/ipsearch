package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	// IpipAPIURL is the ipip api url
	IpipAPIURL = "http://freeapi.ipip.net/"
)

func main() {
	input, err := parseInputParam(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	ips, err := parseInputToTargetIPS(input)
	for _, ip := range ips {
		requestAPI(ip)
		// avoid 403 when request free API
		time.Sleep(time.Second)
	}
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

// LookupIP mock test
var LookupIP = net.LookupIP

// lookupIP look up ip against an domain using net package
// It returns a array of string refer to IPs(ipv4)
func lookupIP(domain string) []string {
	var k []string
	ips, err := LookupIP(domain)
	if err != nil {
		fmt.Print(err)
	}

	// only return ipv4
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			k = append(k, ip.String())
		}
	}
	return k
}

// parseInputParam get the extract only 1 parameter
// It returns the first input parameter
func parseInputParam(input []string) (string, error) {
	if len(input) < 1 {
		return "", errors.New("lack param , an ip or url maybe")
	} else if len(input) >= 2 {
		// fmt.Println("param number wrong, only one is acceptable")
		return "", errors.New("too many params, an ip or url maybe")
	}
	return input[0], nil
}

func parseInputToTargetIPS(input string) ([]string, error) {
	if isIP(input) {
		return []string{input}, nil
	} else if isURL(input) {
		u, _ := url.Parse(input)
		var host string
		if u.Host != "" {
			host = u.Host
		} else {
			host = u.Path
		}
		fmt.Println("lookup host:", host)
		return lookupIP(host), nil
	} else {
		return nil, errors.New("input format error, should be IP or domain")
	}
}

// isIP judge whether the string  an ip or not
func isIP(ip string) bool {
	if m, _ := regexp.MatchString(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, ip); !m {
		return false
	}
	return true
}

// isDomain judge whether the string an domain or not
func isURL(raw string) bool {
	if m, _ := url.Parse(raw); m == nil {
		return false
	}
	return true
}

// requestAPI
func requestAPI(ip string) {
	resp, err := http.Get(IpipAPIURL + ip)
	if err != nil {
		fmt.Printf("http client error %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := extractData(string(body))
	location := strings.Join(s, "-")
	fmt.Println(ip + " " + location)
}
