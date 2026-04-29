include .env
export

service-run:	
	go run .

service-deploy:
	doker compose up -d application

service-undeploy:	
	doker compose down application

some-target:
	export env_var=hello && printenv env_var

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up	

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down