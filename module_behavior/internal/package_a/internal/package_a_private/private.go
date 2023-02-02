package package_a_private

import (
	"fmt"
)

func PackageAPrivate() {
	fmt.Println("Package A private method, only callable by Package A or root module")
}
