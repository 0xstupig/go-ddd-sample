test:
	go test ./service/... -coverprofile cover.out.tmp
	go tool cover -func cover.out.tmp

mock:
	cd ./service; mockery --all --keeptree --output=mock
	cd ./infrastructure; mockery --all --keeptree --output=mock