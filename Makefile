# コンテナ起動
up:
	docker-compose up

# DBを作成
createdb:
	docker exec -it shopping_db createdb --username=postgres --owner=postgres shopping_db

# DB削除
dropdb:
	docker exec -it shopping_db dropdb --username=postgres shopping_db

.PHONY: createdb dropdb up
