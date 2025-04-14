# sort64id - Sortable Base64 Universally Unique Identifiers

TL;DR: short, sortable UUIDs that work as identifiers in CSS, XML, URLs, and more.

## Example

To get three UUID V7 ids:

    $ sort64id -n 3 -v 4
    LCVinfjYJ6dfvhanqnTO1Q
    QTVYh709oHFbdZAc-qwDzj
    JH7x-aBo3PmftfFfZp85sQ

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

