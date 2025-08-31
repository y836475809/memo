package main

import "test/p3/sv"

func main() {
	r := sv.SetupRouter3()
	r.Run(":8080")
}
