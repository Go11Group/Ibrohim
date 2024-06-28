CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:root@localhost:5432/weather_transport?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
exp:
	export POSTGRESQL_URL=${DB_URL}

mig-create:
	@migrate create -ext sql -dir storage/migrations -seq ${1}_${2}_${3}

mig-up:
	migrate -database ${DB_URL} -path storage/migrations -verbose up

mig-down:
	migrate -database ${DB_URL} -path storage/migrations -verbose down

mig-force:
	@migrate -database ${DB_URL} -path storage/migrations -verbose force ${1}