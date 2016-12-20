package main

var (
	base = Files{"main.go",
		`
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!!")
}
`}

	server = Files{"main.go", `
  package main

  import (
  	"io"
  	"net/http"
  )

  func hello(w http.ResponseWriter, r *http.Request) {
  	io.WriteString(w, "Hello world!")
  }

  func main() {
  	http.HandleFunc("/", hello)
  	http.ListenAndServe(":8000", nil)
  }
  `}

	yaml = Files{"app.yaml", `
application: <application number>
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
`}
	app = Files{"hello.go", `
  package hello

import (
	"fmt"
	"net/http"
)

// init is run before the application starts serving.
func init() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Go app")
}
`}

	codecon = Files{"main.go",
		`
	// cat data || go run main.go

	package main

	import (
		"bufio"
		"log"
		"os"
		"strconv"
	)

	var (
		c, s, max int
		lala      []string
		ner       *bufio.Scanner
		loop      bool
		f         *os.File
		fixed     int
	)

	func main() {
		if os.Getenv("CODECON") == "local" {
			var e error
			f, e = os.Open("data")
			if e != nil {
				panic(e)
			}
			ner = bufio.NewScanner(f)
		} else {
			ner = bufio.NewScanner(os.Stdin)
		}
		if fixed != 0 {
			max = fixed
		}
		for ner.Scan() && loop {
			// log.Println(ner.Text(), "   ", max)
			if fixed == 0 && c == 0 {
				setMax(ner.Text())
				c++
				break
			}
			if c >= max {
				endScanner()
				break
			}
			do(ner.Text())
			c++
		}
		f.Close()
		final()
	}

	func setMax(s string) {
		max, _ = strconv.Atoi(s)
		max++
	}

	func endScanner() {
		loop = false
	}

	func init() {
		loop = true
		fixed = 2
	}

	func do(p string) {
		lala = append(lala, p)
	}

	func final() {
		log.Println(lala)
	}

	`}
	data = Files{"data", `
gwgwwgbwwwbgbwgwwbwwb
gggwwbwwgbwwwwgwwgwbbbwgwwwbgwwwb
`}
)
