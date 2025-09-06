package main

import "test/p1/sv"

func main() {
	r := sv.SetupRouter1()
	r = sv.PostUser(r)
	r.Run(":8080")
}
