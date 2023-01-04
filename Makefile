
generate:
	docker run --rm -v $(CURDIR):/data openapitools/openapi-generator-cli \
		generate -i https://stoplight.io/api/v1/projects/omneex/api/nodes/reference/WB-Supplier-Stats.json -g go -o /data --skip-validate-spec \
		--git-host github.com --git-user-id kvarts --git-repo-id wb-stats-go-client
