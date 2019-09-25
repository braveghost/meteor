package file

import (
	"fmt"
	"testing"
)

func TestPathExist(t *testing.T) {
	fmt.Println(PathExist("file/dir.go"))
}
