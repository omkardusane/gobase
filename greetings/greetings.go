package greetings

import "gobase_om/t1"

func Hello(name string) int {
	// a way to import functions from other submodules
	return len(name) + t1.Yellow()
}
