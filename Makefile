debug:
	dlv --headless \
		--api-version 2 \
		--listen "127.0.0.1:38515" \
		debug ./cmd/experid -- start