package package_b

import (
	"fmt"

	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_a"
)

func PackageBPublic() {
	fmt.Println("Package B public")
}

func PackageBCallsPackageAPublic() {
	package_a.PackageAPublic()
	package_a.PackageACallsPrivateMethod()

	/*
		Not allowed

		Attempting to call this method directly even with the required import shows the following error:

		package command-line-arguments
		imports github.com/dougbarth/golang-exploration/module_behavior/internal/package_b
		internal/package_b/public.go:7:2: use of internal package github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private not allowed
	*/
	// package_a_private.PackageAPrivate()
}
