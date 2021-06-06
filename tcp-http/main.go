package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

func router(conn net.Conn, url string, method string) {
	var body string
	if url == "/" {
		body = "HOME"
	} else if url == "/rohan" {
		body = "ROHAN"
	} else if url == "/kumar" {
		body = "KUMAR"
	} else if url == "/toppr" {
		body = "TOPPR"
	}

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", 1000)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	tpl.ExecuteTemplate(conn, "index.gohtml", body)

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()
	var headers []string = []string{}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(string(ln))
		if len(ln) == 0 {
			break
		}
		headers = append(headers, string(ln))
	}
	var url, method string = "/", "NONE"
	if len(headers) > 0 {
		method = strings.TrimSpace(strings.Split(headers[0], " ")[0])
		url = strings.TrimSpace(strings.Split(headers[0], " ")[1])
		fmt.Printf("Method : %s\n", method)
		fmt.Printf("URL : %s\n", url)
	}

	router(conn, url, method)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		serve(conn)
	}
}
