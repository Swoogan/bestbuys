#!/bin/bash

$GOPATH/bin/command -l command/command.log &
$GOPATH/bin/denormalizerd -l denormalizerd/denormalizerd.log &
query/mongorestd -d query -a :4040 -o query/query.log games,ro &
