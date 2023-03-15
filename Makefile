.PHONY: open_api_http
open_api_http:
	@./scripts/open_api_http.sh auth ports internal/oauth/port

proto_auth:
	@./scripts/proto.sh auth auth