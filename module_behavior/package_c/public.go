package package_c

import (
	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_a"
	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_b"
)

func PackageCCanCallPublicMethods() {
	package_a.PackageAPublic()
	package_b.PackageBPublic()
}
