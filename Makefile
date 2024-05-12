build:
	docker build -t golang-local:latest .
	docker compose up -d --remove-orphans
