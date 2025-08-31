package main

import "test/p2/sv"

func main() {
	r := sv.SetupRouter2()
	r.Run(":8080")
}
