package main

import (
	"fmt"

	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_a"
	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_b"
	"github.com/dougbarth/golang-exploration/module_behavior/package_c"
)

func main() {
	fmt.Println("--- package_a.PackageAPublic()")
	package_a.PackageAPublic()

	fmt.Println()
	fmt.Println("--- package_a.PackageACallsPrivateMethod()")
	package_a.PackageACallsPrivateMethod()
	/*
		Not allowed.

		The following error is shown even if we import "github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private"

			$ go run main.go
			package command-line-arguments
				main.go:5:2: use of internal package github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private not allowed
	*/
	//package_a_private.PackageAPrivate()

	fmt.Println()
	fmt.Println("--- package_b.PackageBPublic()")
	package_b.PackageBPublic()

	fmt.Println()
	fmt.Println("--- package_b.PackageBCallsPackageAPublic()")
	package_b.PackageBCallsPackageAPublic()

	fmt.Println()
	fmt.Println("--- package_c.PackageCCanCallPublicMethods()")
	package_c.PackageCCanCallPublicMethods()
}
