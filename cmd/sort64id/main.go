/*
sort64id is a command line tool to generate and check sort64ids.

Usage:

  sort64id [-c] [-n number] [-v version] [ns name | [uuid ...]]

Flags:

  -c
      Checking mode, verifies given UUID(s).
  -n number
      Number of IDs to generate.
  -v version
      UUID version to use to generate ID.

*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/gofrs/uuid"
	"github.com/sdassow/go-sort64id"
)

func newId(version int) (uuid.UUID, error) {
	switch version {
	case 1:
		return uuid.NewV1()
	case 3:
		nsid, err := sort64id.FromString(flag.Arg(0))
		if err != nil {
			return uuid.Nil, err
		}
		return uuid.NewV3(nsid, flag.Arg(1)), nil
	case 4:
		return uuid.NewV4()
	case 5:
		nsid, err := sort64id.FromString(flag.Arg(0))
		if err != nil {
			return uuid.Nil, err
		}
		return uuid.NewV5(nsid, flag.Arg(1)), nil
	case 6:
		return uuid.NewV6()
	case 7:
		return uuid.NewV7()
	}
	return uuid.Nil, fmt.Errorf("unknown uuid version: %d", version)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-c] [-n amount] [-v version] [ns name]\n", path.Base(os.Args[0]))
	os.Exit(1)
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", path.Base(os.Args[0]), err.Error())
	os.Exit(2)
}

func main() {
	amount := 1
	check := false
	version := 6

	flag.BoolVar(&check, "c", false, "check and verify given ids")
	flag.IntVar(&amount, "n", 1, "number of uuids to generate")
	flag.IntVar(&version, "v", 6, "uuid version to use")
	flag.Parse()

	if check {
		for n := 0; n < flag.NArg(); n++ {
			id, err := sort64id.FromString(flag.Arg(n))
			if err != nil {
				fail(err)
			}

			if len(flag.Arg(n)) == 22 {
				if flag.Arg(n) != sort64id.ToString(id) {
					fail(fmt.Errorf("roundtrip mismatch: %s != %s\n", flag.Arg(n), sort64id.ToString(id)))
				}
			} else {
				if flag.Arg(n) != id.String() {
					fail(fmt.Errorf("roundtrip mismatch: %s != %s\n", flag.Arg(n), id.String()))
				}
			}

			fmt.Printf("%s %s %x\n", sort64id.ToString(id), id.String(), id.Bytes())
		}
		return
	} else if version == 3 || version == 5 {
		if flag.NArg() != 2 {
			usage()
		}
	} else {
		if flag.NArg() != 0 {
			usage()
		}
	}

	for n := 0; n < amount; n++ {
		id, err := newId(version)
		if err != nil {
			fail(err)
		}

		fmt.Printf("%s\n", sort64id.ToString(id))
	}
}
