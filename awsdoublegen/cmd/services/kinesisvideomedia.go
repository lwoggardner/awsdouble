
// +build doublegen
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00

package main

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	"github.com/lwoggardner/awsdouble/awsdoublegen"
)

func main() {
	awsdoublegen.NewAPIInterface((*kinesisvideomediaiface.KinesisVideoMediaAPI)(nil)).Generate("double.go")
}

