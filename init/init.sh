#!/bin/bash

DATA=`cat aol.json`
curl -i -H "Content-type: application/json" -X POST -d "$DATA" http://localhost/commands/  

DATA=`cat fow.json`
curl -i -H "Content-type: application/json" -X POST -d "$DATA" http://localhost/commands/  

DATA=`cat dg.json`
curl -i -H "Content-type: application/json" -X POST -d "$DATA" http://localhost/commands/  
