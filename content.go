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
)
