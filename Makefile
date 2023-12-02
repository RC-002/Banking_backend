postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres16 dropdb --username=root --owner=root bank

.PHONY: postgres createdb dropdb