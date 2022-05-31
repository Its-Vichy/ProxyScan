/*
	https://github.com/Its-Vichy/ProxyScan
*/

package main

import (
	"bufio"
	"fmt"
	"h12.io/socks"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// uwu: https://github.com/asm-jaime/go-proxycheck/blob/master/pcheck.go
func ProxyReq(req string, proxy string) (res *http.Response, err error) {
	timeout := time.Duration(2 * time.Second)
	proxyURL, err := url.Parse("http://" + proxy)
	reqURL, err := url.Parse(req)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	res, err = client.Get(reqURL.String())
	return res, err
}

func SockProxyReq(req string, proxy string, ptype int) (res *http.Response, err error) {
	dialSocksProxy := socks.Dial(fmt.Sprintf("socks%d://%s?timeout=5s", ptype, proxy))
	tr := &http.Transport{Dial: dialSocksProxy}
	client := &http.Client{Transport: tr}
	
	res, err = client.Get(req)
	return res, err
}

func check(ip string, port int, file_path string) {
	_, err := SockProxyReq("https://google.com", ip+":"+strconv.Itoa(port), 4)

	if err == nil {
		fmt.Printf("[SOCKS4 VALID] %s:%d\n", ip, port)

		f, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		if _, err := f.WriteString(fmt.Sprintf("%s:%d\n", ip, port)); err != nil {
			log.Println(err)
		}

	} else {
		_, errr := SockProxyReq("https://google.com", ip+":"+strconv.Itoa(port), 5)

		if errr == nil {
			fmt.Printf("[SOCKS5 VALID] %s:%d\n", ip, port)

			f, errr := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if errr != nil {
				log.Println(errr)
			}

			defer f.Close()

			if _, errr := f.WriteString(fmt.Sprintf("%s:%d\n", ip, port)); errr != nil {
				log.Println(errr)
			}

		}
	}
}

func main() {
	fmt.Println("|*> ProxyScan - https://github.com/Its-Vichy/ProxyScan.")
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("|*> Usage:\n - zmap: sudo zmap -p <port> | ./proxy-scan <port> <output file>\n - check file: cat <file> | ./proxy-scan <port> <output file>")
		return
	}

	port, _ := strconv.Atoi(args[0])
	file_path := args[1]

	for {
		data := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(data)

		for scanner.Scan() {
			go check(strings.Split(scanner.Text(), ":")[0], port, file_path)
			time.Sleep(1 * time.Millisecond)
		}
	}
}
