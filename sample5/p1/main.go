package main

import "test/p1/sv"

func main() {
	r := sv.SetupRouter()
	r = sv.PostUser(r)
	r.Run(":8080")
}
