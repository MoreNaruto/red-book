.PHONY: compileProto upMigrations downMigrations statusMigrations

compileProto:
	protoc --go_out=. pkg/models/*.proto
upMigrations:
	cd migrations && goose up
downMigrations:
	cd migrations && goose down
statusMigrations:
	cd migrations && goose status
createMigration:
	cd migrations && goose create $(name) sql
run:
	go run cmd/main.go