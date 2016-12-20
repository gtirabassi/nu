package main

import "time"

const readme = `
## {{.Name}}

this is a basic introduction to my project


{{.Author}}
{{.Date}}
`

//Readme is used for filling template above
type Readme struct {
	Name   string
	Author string
	Date   string
}

func getMe(n string) Readme {
	var me Readme
	me.Author = "Giacomo Tirabassi (drymonsoon)"
	me.Name = n
	me.Date = time.Now().Format(time.RFC850)
	return me
}
