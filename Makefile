test-go:
	go test -v -cover -p=1 -count=1 ./internal/user/user_test.go
test-fail-go:
	go test -v -cover -p=1 -count=1 ./internal/user/user_test.go -testa=false