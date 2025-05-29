.PHONY: test test-coverage clean

# Executa todos os testes
test:
	go test -v ./...

# Executa os testes com cobertura
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Limpa arquivos gerados
clean:
	rm -f coverage.out coverage.html 