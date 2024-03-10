package main

import "github.com/adrichard/siderproject/internal"

// Feed json files to elasticsearch
func main() {
	internal.InitBootStrap().Run()
}
