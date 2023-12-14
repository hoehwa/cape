/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/jipilmuk/cape/cmd"
	_ "github.com/jipilmuk/cape/cmd/fundamental"
	_ "github.com/jipilmuk/cape/cmd/law"
	_ "github.com/jipilmuk/cape/cmd/paradigm"
	_ "github.com/jipilmuk/cape/cmd/principle"
)

func main() {
	cmd.Execute()
}
