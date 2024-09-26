package main

import "strings"

func ChackFlag(messageStr string) (string, string) {
	var name string
	if strings.HasPrefix(messageStr, "-change_name") {
		flags := strings.Split(messageStr, " ")
		if len(flags) == 1 {
			return "name", "The Exact Format Is : [-change_name <Name> ]\n"
		}
		name = flags[1]
		if !NameExists(name) || len(flags) > 2 {
			return name, "Invalid or duplicate name. Try again.\n"
		} else {
			return name, ""
		}

	}
	return "", ""
}
