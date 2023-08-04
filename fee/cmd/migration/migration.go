package main

import (
	"fmt"
	"github.com/smapig/go-ddd-sample/fee/ioc"
)

func main() {
	migrator, err := ioc.InitializeSqlMigrator("")
	if err != nil {
		panic(fmt.Errorf("migartion failed: %v \n", err))
	}

	migrator.Migration()
}
