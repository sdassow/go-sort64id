package sort64id

import (
	"encoding/base64"

	"github.com/gofrs/uuid"
)

// base64 alphabet ordered by ASCII values to preserve lexicographic order
const charset = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

// string padding for decoding phase
const alpad = "A"
const arpad = "z"

// binary padding for encoding phase to ensure base64 output starts with
// uppercase character and ends with an alphanumeric value
var blpad = []byte{'-'}
var brpad = []byte{'\xFF'}

// set up base64 encoder with alternative character setup
var sort64idEncoder = base64.NewEncoding(charset)

func FromString(str string) (uuid.UUID, error) {
	if len(str) != 22 {
		return uuid.FromString(str)
	}
	bin, err := sort64idEncoder.DecodeString(alpad + str + arpad)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromBytes(bin[1:17])
}

func ToString(id uuid.UUID) string {
	bin := append(append(blpad, id.Bytes()...), brpad...)
	b64 := sort64idEncoder.EncodeToString(bin)

	return b64[1:23]
}
