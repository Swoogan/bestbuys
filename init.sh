!/bin/sh

curl -i -X POST -d '{"name": "createGame", "data": {"name":"Age of Legends"}}' http://localhost/tasks/  
curl -i -X POST -d '{"name": "createGame", "data": {"name":"Forces of War"}}' http://localhost/tasks/  
curl -i -X POST -d '{"name": "createGame", "data": {"name":"Dark Galaxy"}}' http://localhost/tasks/  
