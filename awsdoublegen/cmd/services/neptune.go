
// +build doublegen
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00

package main

import (
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
	"github.com/lwoggardner/awsdouble/awsdoublegen"
)

func main() {
	awsdoublegen.NewAPIInterface((*neptuneiface.NeptuneAPI)(nil)).Generate("double.go")
}

