DB_URL = postgres://postgres:root@localhost:5432/metro_system?sslmode=disable

cl:
	clear
exp:
	export POSTGRESQL_URL=${DB_URL}

mig-create:
	@migrate create -ext sql -dir migrations -seq ${1}_${2}_${3}

mig-up:
	@if [ -z "${1}" ]; then \
		migrate -database ${DB_URL} -path migrations -verbose up; \
	else \
		migrate -database ${DB_URL} -path migrations -verbose up ${1}; \
	fi

mig-down:
	@if [ -z "${1}" ]; then \
		migrate -database ${DB_URL} -path migrations -verbose down; \
	else \
		migrate -database ${DB_URL} -path migrations -verbose down ${1}; \
	fi

mig-force:
	@migrate -database ${DB_URL} -path migrations -verbose force ${1}