package main

import (
	"fmt"

	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_a"
	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_b"
)

func main() {
	package_a.PackageAPublic()

	fmt.Println("------------------")
	package_a.PackageACallsPrivateMethod()
	/*
		Not allowed.

		The following error is shown even if we import "github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private"

			$ go run main.go
			package command-line-arguments
				main.go:5:2: use of internal package github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private not allowed
	*/
	//package_a_private.PackageAPrivate()

	fmt.Println("------------------")
	package_b.PackageBPublic()

	fmt.Println("------------------")
	package_b.PackageBCallsPackageAPublic()
}
