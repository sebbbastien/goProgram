package main

import (
	gomodulea "github.com/sebbbastien/goModuleA"
	gomoduleb "github.com/sebbbastien/goModuleB"
)

func main() {
	gomodulea.FctFromModuleA()
	gomoduleb.FctFromModuleB()
}
