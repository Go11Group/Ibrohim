CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:root@localhost:5432/e_commerce_auth?sslmode=disable

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs
proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
exp:
	export POSTGRESQL_URL=${DB_URL}
test:
	go test ./storage/postgres
mig-cr:
	@migrate create -ext sql -dir migrations -seq ${1}_${2}_${3}

mig-up:
	@migrate -database ${DB_URL} -path migrations -verbose up

mig-down:
	@migrate -database ${DB_URL} -path migrations -verbose down

mig-force:
	@migrate -database ${DB_URL} -path migrations -verbose force ${1}