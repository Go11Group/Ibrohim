CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:root@localhost:5432/collaborations?sslmode=disable

tidy:
	go mod tidy
	go mod vendor
gen-proto:
	./scripts/gen-proto.sh ${CURRENT_DIR}
exp:
	export POSTGRESQL_URL=${DB_URL}

mig-create:
	@migrate create -ext sql -dir migrations -seq ${1}_${2}_${3}

mig-up:
	@migrate -database ${DB_URL} -path migrations -verbose up

mig-down:
	@migrate -database ${DB_URL} -path migrations -verbose down

mig-force:
	@migrate -database ${DB_URL} -path migrations -verbose force ${1}