package test

import (
	"fmt"
	"testing"

	"github.com/heqzha/dcache/core"
)

func TestVClockCompare(t *testing.T) {
	n1 := core.VClock{}
	n2 := core.VClock{}

	n1.Set("a", 1)
	n1.Set("b", 1)

	n2.Set("a", 1)
	n2.Set("c", 1)

	fmt.Println(n1.ReturnVCString())
	fmt.Println(n2.ReturnVCString())
	fmt.Println(n1.Compare(n2, core.Concurrent))

	n1 = core.VClock{}
	n2 = core.VClock{}
	n1.Set("a", 1)
	n1.Set("b", 1)

	n2.Set("a", 1)
	n2.Set("b", 1)

	fmt.Println(n1.ReturnVCString())
	fmt.Println(n2.ReturnVCString())
	fmt.Println(n1.Compare(n2, core.Equal))

	n1 = core.VClock{}
	n2 = core.VClock{}
	n1.Set("a", 1)
	n1.Set("b", 1)

	n2.Set("a", 0)
	n2.Set("b", 1)

	fmt.Println(n1.ReturnVCString())
	fmt.Println(n2.ReturnVCString())
	fmt.Println(n1.Compare(n2, core.Ancestor))

	n1 = core.VClock{}
	n2 = core.VClock{}
	n1.Set("a", 1)
	n1.Set("b", 1)

	n2.Set("a", 1)
	n2.Set("b", 1)
	n2.Set("c", 3)

	fmt.Println(n1.ReturnVCString())
	fmt.Println(n2.ReturnVCString())
	fmt.Println(n1.Compare(n2, core.Descendant))
}
