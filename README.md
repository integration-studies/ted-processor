## Run Postgres container

First we need to create a container instance for postgres database

```shell script

docker run -d -p 5432:5432 --name tp_database -e POSTGRES_PASSWORD=admin  postgres

```