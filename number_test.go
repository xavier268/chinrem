package chinrem

import (
	"fmt"
	"testing"
)

func TestCRIVisual(t *testing.T) {

	e := NewCREngine(10)
	a := e.NewCRI()

	fmt.Println(a)

}
