package package_a

import (
	"fmt"

	"github.com/dougbarth/golang-exploration/module_behavior/internal/package_a/internal/package_a_private"
)

func PackageAPublic() {
	fmt.Println("Called a Package A's public method")
}

func PackageACallsPrivateMethod() {
	package_a_private.PackageAPrivate()
}
