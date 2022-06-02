build-prod:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o kmsbot.exe .

migrate-new: migrate-cli
	./migrate-cli create -ext sql -dir migrations "$(name)"

migrate-up: migrate-cli
	./migrate-cli -database "sqlite3://bot.db?_auth_user=root&_auth_pass=secret&&query" -path migrations up

migrate-down: migrate-cli
	./migrate-cli -database "sqlite3://bot.db?_auth_user=root&_auth_pass=secret&&query" -path migrations down 1