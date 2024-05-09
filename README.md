docker-compose up todo-app

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

docker ps
docker exec -it bf39dfdd9f49 /bin/bash
psql -U postgres
\d