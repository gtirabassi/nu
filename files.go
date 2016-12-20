package main

//Files contains all the info
type Files struct {
	Name    string
	Content string
}

func getFiles() map[string][]Files {
	j := make(map[string][]Files)
	j["base"] = []Files{base}
	j["server"] = []Files{server}
	j["gae"] = []Files{yaml, app}
	return j
}
