docker-compose up --build todo-app

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

docker ps
docker exec -it {container} /bin/bash
psql -U postgres
\d
