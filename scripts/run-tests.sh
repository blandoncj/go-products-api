#!/bin/bash

echo "🧪 Ejecutando tests de todos los servicios..."
echo ""

services=("create-service" "delete-service" "read-service" "update-service")
failed=0

for service in "${services[@]}"; do
  echo "📦 Testing $service..."
  cd "../services/$service"

  if go test ./... -v -cover; then
    echo "✅ $service: PASSED"
  else
    echo "❌ $service: FAILED"
    ((failed++))
  fi

  cd ../../scripts
  echo ""
done

if [ $failed -eq 0 ]; then
  echo "🎉 Todos los tests pasaron exitosamente!"
  exit 0
else
  echo "💥 $failed servicio(s) fallaron"
  exit 1
fi
