#!/bin/bash

test -e go.mod || cd ..

go vet .
~/go/bin/errcheck
