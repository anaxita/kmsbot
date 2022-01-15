MIGRATE_TOOL_URL := https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz

appConfigFile = "./src/config.yml"

dbName = `yq e '.database.name' $(appConfigFile)`
dbUser = `yq e '.database.user' $(appConfigFile)`
dbPassword = `yq e '.database.password' $(appConfigFile)`

build-prod:
	cd src && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o kmsbot.exe .
migrate-cli:
	curl -L ${MIGRATE_TOOL_URL} | tar xvz
	mv migrate.linux-amd64 $@
	chmod +x $@

migrate-new: migrate-cli
	./migrate-cli create -ext sql -dir src/migrations "$(name)"

migrate-up: migrate-cli
	./migrate-cli -database "sqlite3://src/bot.db?_auth_user=root&_auth_pass=secret&&query" -path src/migrations up

migrate-down: migrate-cli
	./migrate-cli -database "sqlite3://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&&query" -path src/migrations down 1