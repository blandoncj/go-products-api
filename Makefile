.PHONY: test test-create test-delete test-read test-update test-all test-coverage

# Ejecutar todos los tests
test-all:
	@echo "🧪 Testing create-service..."
	@cd services/create-service && go test ./... -v
	@echo "\n🧪 Testing delete-service..."
	@cd services/delete-service && go test ./... -v
	@echo "\n🧪 Testing read-service..."
	@cd services/read-service && go test ./... -v
	@echo "\n🧪 Testing update-service..."
	@cd services/update-service && go test ./... -v

# Tests individuales
test-create:
	@cd services/create-service && go test ./... -v

test-delete:
	@cd services/delete-service && go test ./... -v

test-read:
	@cd services/read-service && go test ./... -v

test-update:
	@cd services/update-service && go test ./... -v

# Tests con cobertura
test-coverage:
	@echo "📊 Generando reporte de cobertura..."
	@cd services/create-service && go test ./... -coverprofile=../../coverage-create.out
	@cd services/delete-service && go test ./... -coverprofile=../../coverage-delete.out
	@cd services/read-service && go test ./... -coverprofile=../../coverage-read.out
	@cd services/update-service && go test ./... -coverprofile=../../coverage-update.out
	@echo "✅ Reportes generados: coverage-*.out"

# Limpiar
clean:
	rm -f coverage-*.out

# Ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  make test-all       - Ejecutar todos los tests"
	@echo "  make test-create    - Ejecutar tests de create-service"
	@echo "  make test-delete    - Ejecutar tests de delete-service"
	@echo "  make test-read      - Ejecutar tests de read-service"
	@echo "  make test-update    - Ejecutar tests de update-service"
	@echo "  make test-coverage  - Generar reportes de cobertura"
	@echo "  make clean          - Limpiar archivos generados"
