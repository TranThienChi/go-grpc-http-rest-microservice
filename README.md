# go-grpc-http-rest-microservice

Reference: https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9
.\server -grpc-port=9090 -http-port=8080 -db-host=localhost:3306 -db-user=root -db-password=CHUNGchi1989 -db-schema=sql_todo
.\client-grpc -server=localhost:9090
.\client-rest -server=localhost:8080

### MySQL create-db-todo
DROP DATABASE IF EXISTS sql_todo;
CREATE DATABASE sql_todo;
USE sql_todo;

CREATE TABLE ToDo (
		ID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(200) DEFAULT NULL,
    Description VARCHAR(1024) DEFAULT NULL,
    Reminder TIMESTAMP NULL DEFAULT NULL
);

### Run
docker-compose up -d --build
