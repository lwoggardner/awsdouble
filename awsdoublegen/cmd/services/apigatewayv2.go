
// +build doublegen
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00

package main

import (
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/lwoggardner/awsdouble/awsdoublegen"
)

func main() {
	awsdoublegen.NewAPIInterface((*apigatewayv2iface.ApiGatewayV2API)(nil)).Generate("double.go")
}

