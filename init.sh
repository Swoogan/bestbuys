#!/bin/sh

curl -i -H "Content-type: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Age of Legends","lands":[{"name":"Rocky Coast","cost":50000,"income":500},{"name":"Fertile Valley","cost":500000,"income":3000},{"name":"Rustic Forest","cost":10000000,"income":20000},{"name":"Ancient Highlands","cost":100000000,"income":50000}]}}' http://localhost/tasks/  

curl -i -H "Content-type: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Forces of War","lands":[{"name":"Small Complex","cost":50000,"income":500},{"name":"Industrial Field","cost":500000,"income":3000},{"name":"Ocean Platform","cost":10000000,"income":20000},{"name":"Small Island","cost":100000000,"income":50000}]}}' http://localhost/tasks/  

curl -i -H "Content-type: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Dark Galaxy", "lands":[{"name":"Scorched Sands","cost":50000,"income":500},{"name":"Blistering Desert","cost":500000,"income":3000}]}}' http://localhost/tasks/  
