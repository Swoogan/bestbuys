!/bin/sh

curl -i -H "Accept: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Age of Legends"}}' http://localhost/tasks/  
curl -i -H "Accept: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Forces of War"}}' http://localhost/tasks/  
curl -i -H "Accept: application/json" -X POST -d '{"name": "createGame", "data": {"name":"Dark Galaxy"}}' http://localhost/tasks/  
