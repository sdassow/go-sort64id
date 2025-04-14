# sort64id - Sortable Base64 Universally Unique Identifiers

TL;DR: short, sortable UUIDs that work as identifiers in CSS, XML, URLs, and more.

## Example

To get three UUID V4 ids on the command line:

    $ sort64id -n 3 -v 4
    LCVinfjYJ6dfvhanqnTO1Q
    QTVYh709oHFbdZAc-qwDzj
    JH7x-aBo3PmftfFfZp85sQ

To convert a UUID in code:
    
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
    

## Installation

    go install github.com/sdassow/go-sort64id/cmd/sort64id@latest

## Encoding

This is a compact representation of UUIDs and builds on top intead of making changes.

The three key ingredients are:
- Encode to Base64 for fewer characters.
- Pad UUID bytes before encoding and strip afterwards to use the padding and guarantee a range of start and end characters.
- Adjust Base64 alphabet to make strings lexicographically sortable.


## Background

This design is an evolution of [RUID](https://github.com/sdassow/pg-ruid), an
encoding published 12 years ago as a PostgreSQL extension.

Now the encoding takes the sort order of Base64 characters into account and
adjusts the alphabet to make the ids sortable in binary order.

An alternative alphabet of existing implementation of sortable base64 can be seen
at https://www.codeproject.com/Articles/5165340/Sortable-Base64-Encoding.

