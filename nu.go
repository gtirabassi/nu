package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	n, t string
	g    bool
)

func init() {
	flag.BoolVar(&g, "git", true, "a git repository is created if not specified otherwhise")
	// flag.StringVar(&n, "n", "example", "a name is necessary for the creation of a new project (shorthand)")
	// flag.StringVar(&n, "name", "example", "a name is necessary for the creation of a new project")
	flag.StringVar(&t, "type", "base", "define the type of project to be created [ base | server | gae ]")
	flag.StringVar(&t, "t", "base", "define the type of project to be created [ base | server | gae ] (shorthand)")
}

func main() {
	readtpl, err := template.New("readme").Parse(readme)
	check(err)

	flag.Parse()
	n = flag.Arg(0)
	log.Println(n)
	if n == "example" {
		exitError("A name MUST be assigned to the project (-n/-name)")
	}
	v, ok := getFiles()[t]
	if !ok {
		exitError("The specified type DOES NOT exist")
	}
	n = strings.Replace(n, " ", "_", -1)
	// fmt.Println("--" + n + "--")
	if existD(n) {
		exitError("There already is a folder with the same name!! --> try again")
	}
	if g == true {
		res, er := exec.Command("/bin/sh", "-c", "git init "+n).Output()
		check(er)
		if !strings.Contains(string(res), "Initialized empty") {
			exitError(`Something went wrong with GIT:
--- git output ---
` + string(res) + `--- end output ---`)
		}
	} else {
		res, er := exec.Command("/bin/sh", "-c", "mkdir "+n).Output()
		check(er)
		if string(res) != "" {
			exitError("something wrong creating the folder for your project")
		}
	}
	for _, k := range v {
		err = ioutil.WriteFile(n+"/"+k.Name, []byte(k.Content), 0664)
		check(err)
	}
	f, err := os.Create(n + "/README.md")
	check(err)
	defer f.Close()
	readtpl.Execute(f, getMe(n))
	color.Green("Enjoy your Project!!")
	color.Green("Opening ATOM as we speak")
	time.Sleep(time.Second)
	_, er := exec.Command("/bin/sh", "-c", "atom "+n).Output()
	check(er)
	os.Exit(0)
}

func existD(p string) bool {
	st, e := os.Stat(n)
	if e != nil {
		if os.IsExist(e) {
			return true
		}
	} else {
		if st.IsDir() {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		exitError(e.Error())
	}
}

func exitError(e string) {
	color.Red(e)
	os.Exit(0)
}
