package sort64id_test

import (
	"fmt"

	"github.com/gofrs/uuid"
    	"github.com/sdassow/go-sort64id"
)

func ExampleString() {
	id := uuid.NewV5(uuid.Nil, "foo")
	s64id := sort64id.ToString(id)
	fmt.Printf("s64id=%v (%v), uuid=%v (%v)", s64id, len(s64id), id.String(), len(id.String()))
	// Output: s64id=PdpADe17_j7fC_KLkZB2Bj (22), uuid=aa752cea-8222-5bc8-acd9-555b090c0ccb (36)
}
