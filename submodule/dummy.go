// dummy.go
package submodule

import "dunglas.dev/go_module_to_rename"

func World() string {
	return go_module_to_rename.Hello()
}
