#!/bin/bash

bin/command -l command/command.log &
bin/denormalizer-server -l denormalizer/denormalizer.log &
query/mongorestd -d query -a :4040 -o query/query.log games,ro &
