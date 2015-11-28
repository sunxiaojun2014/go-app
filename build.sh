#!/bin/bash
PRJ_PATH=`pwd`

export GOPATH=${PRJ_PATH}
export GOBIN=${PRJ_PATH}/bin

go install main
