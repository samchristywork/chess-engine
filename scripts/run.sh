#!/bin/bash

test -e go.mod || cd ..

go run .
