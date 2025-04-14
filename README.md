# sort64id - Sortable Base64 Universally Unique Identifiers

This design is an evolution of [RUID](https://github.com/sdassow/pg-ruid), an
encoding published 12 years ago as a PostgreSQL extension.

Now the encoding takes the sort order of Base64 characters into account and
adjusts the alphabet to make the ids sortable in binary order.

An alternative alphabet of existing implementation of sortable base64 can be seen
at https://www.codeproject.com/Articles/5165340/Sortable-Base64-Encoding.

