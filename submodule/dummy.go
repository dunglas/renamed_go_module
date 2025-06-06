// dummy.go
package submodule

import "github.com/dunglas/go_module_to_rename"

func World() string {
	return go_module_to_rename.Hello()
}
