#!/bin/bash

mongo localhost/command clean.js
mongo localhost/query --eval "db.games.remove();"
