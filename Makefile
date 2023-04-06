.PHONY: open_api_http
open_api_http:
	@./scripts/open_api_http.sh auth auth port internal/port

proto_auth:
	@./scripts/proto.sh auth auth


.PHONY: build_and_migrate_postgres_auth
build_and_migrate_postgres_auth:
	@./scripts/migrate_postgres_auth.sh migrate_postgres_auth packages/postgres/Dockerfile true postgres_auth

.PHONY: migrate_postgres_auth
migrate_postgres_auth:
	@./scripts/migrate_postgres_auth.sh migrate_postgres_auth packages/postgres/Dockerfile false postgres_auth