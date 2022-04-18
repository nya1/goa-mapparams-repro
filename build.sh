#!/bin/bash

go build calcsvc/cmd/calc

# go build calcsvc/cmd/calc-cli
# note:
# ./calc-cli calc multiply --type 1 --metadata '{"testKey": "val"}'
# doesn't work, the URL generated is `GET /multiply/?testKey=val&type=1`
# instead of `GET /multiply/?metadata[testKey]=val&type=1`
#

