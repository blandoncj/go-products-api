pipeline {
    agent any
    
    environment {
        GO_VERSION = '1.25'
        DOCKER_REGISTRY = 'docker.io'
        DOCKER_CREDENTIALS = credentials('docker-hub-credentials')
        MONGO_URI = 'mongodb://admin:admin123@mongodb:27017/'
    }
    
    stages {
        stage('Checkout') {
            steps {
                echo '📥 Clonando repositorio...'
                checkout scm
            }
        }
        
        stage('Setup') {
            steps {
                echo '🔧 Configurando entorno...'
                sh '''
                    go version
                    docker --version
                    docker-compose --version
                '''
            }
        }
        
        stage('Start MongoDB') {
            steps {
                echo '🍃 Iniciando MongoDB para pruebas...'
                sh '''
                    docker run -d \
                        --name mongodb-jenkins \
                        -p 27017:27017 \
                        -e MONGO_INITDB_ROOT_USERNAME=admin \
                        -e MONGO_INITDB_ROOT_PASSWORD=admin123 \
                        mongo:6
                    
                    # Esperar a que MongoDB esté listo
                    echo "⏳ Esperando MongoDB..."
                    for i in {1..15}; do
                        if docker exec mongodb-jenkins mongosh --eval "db.runCommand({ ping: 1 })" > /dev/null 2>&1; then
                            echo "✅ MongoDB está listo!"
                            break
                        fi
                        echo "Intentando... ($i/15)"
                        sleep 3
                    done
                '''
            }
        }
        
        stage('Test Microservices') {
            parallel {
                stage('Test Create Service') {
                    steps {
                        dir('services/create-service') {
                            echo '🧪 Probando Create Service...'
                            sh '''
                                go mod download
                                go test ./... -v -cover -coverprofile=coverage.out
                                go tool cover -func=coverage.out
                            '''
                        }
                    }
                }
                
                stage('Test Read Service') {
                    steps {
                        dir('services/read-service') {
                            echo '🧪 Probando Read Service...'
                            sh '''
                                go mod download
                                go test ./... -v -cover -coverprofile=coverage.out
                                go tool cover -func=coverage.out
                            '''
                        }
                    }
                }
                
                stage('Test Update Service') {
                    steps {
                        dir('services/update-service') {
                            echo '🧪 Probando Update Service...'
                            sh '''
                                go mod download
                                go test ./... -v -cover -coverprofile=coverage.out
                                go tool cover -func=coverage.out
                            '''
                        }
                    }
                }
                
                stage('Test Delete Service') {
                    steps {
                        dir('services/delete-service') {
                            echo '🧪 Probando Delete Service...'
                            sh '''
                                go mod download
                                go test ./... -v -cover -coverprofile=coverage.out
                                go tool cover -func=coverage.out
                            '''
                        }
                    }
                }
            }
        }
        
        stage('Business Rules Tests') {
            steps {
                echo '📋 ========================================='
                echo '📋 EJECUTANDO PRUEBAS DE REGLAS DE NEGOCIO'
                echo '📋 ========================================='
                
                script {
                    // Regla 1: Precio no puede ser negativo
                    echo '🔍 Regla de Negocio #1: Validación de precio negativo'
                    dir('services/create-service') {
                        sh '''
                            echo "→ Test: Producto con precio negativo debe rechazarse"
                            go test -v -run TestProductService_Create_InvalidProduct
                        '''
                    }
                    
                    // Regla 2: Stock cero permitido (pre-órdenes)
                    echo '🔍 Regla de Negocio #2: Productos con stock 0 permitidos'
                    dir('services/create-service') {
                        sh '''
                            echo "→ Test: Producto con stock 0 debe poder crearse"
                            go test -v -run TestProductService_Create_StockZero
                        '''
                    }
                    
                    // Regla 3: Idempotencia en eliminación
                    echo '🔍 Regla de Negocio #3: Idempotencia en eliminación'
                    dir('services/delete-service') {
                        sh '''
                            echo "→ Test: Eliminar producto inexistente no genera error"
                            go test -v -run TestProductService_DeleteProduct_NotFound
                        '''
                    }
                    
                    // Regla 4: Propagación de errores de BD
                    echo '🔍 Regla de Negocio #4: Propagación de errores de base de datos'
                    dir('services/delete-service') {
                        sh '''
                            echo "→ Test: Errores de BD deben propagarse correctamente"
                            go test -v -run TestProductService_DeleteProduct_DatabaseError
                        '''
                    }
                    
                    echo '✅ ========================================='
                    echo '✅ TODAS LAS REGLAS DE NEGOCIO VALIDADAS'
                    echo '✅ ========================================='
                }
            }
        }
        
        stage('Security Scan') {
            steps {
                echo '🔒 Análisis de seguridad...'
                sh '''
                    # Instalar gosec si no está instalado
                    if ! command -v gosec &> /dev/null; then
                        go install github.com/securego/gosec/v2/cmd/gosec@latest
                    fi
                    
                    # Ejecutar análisis de seguridad
                    gosec -no-fail ./...
                '''
            }
        }
        
        stage('Build Docker Images') {
            steps {
                echo '🐳 Construyendo imágenes Docker...'
                sh '''
                    docker build -t products-create-service:${BUILD_NUMBER} ./services/create-service
                    docker build -t products-read-service:${BUILD_NUMBER} ./services/read-service
                    docker build -t products-update-service:${BUILD_NUMBER} ./services/update-service
                    docker build -t products-delete-service:${BUILD_NUMBER} ./services/delete-service
                    
                    echo "✅ Imágenes construidas con tag: ${BUILD_NUMBER}"
                '''
            }
        }
        
        stage('Push to Docker Hub') {
            when {
                branch 'main'
            }
            steps {
                echo '📤 Subiendo imágenes a Docker Hub...'
                sh '''
                    echo ${DOCKER_CREDENTIALS_PSW} | docker login -u ${DOCKER_CREDENTIALS_USR} --password-stdin
                    
                    docker tag products-create-service:${BUILD_NUMBER} ${DOCKER_CREDENTIALS_USR}/products-create-service:latest
                    docker tag products-read-service:${BUILD_NUMBER} ${DOCKER_CREDENTIALS_USR}/products-read-service:latest
                    docker tag products-update-service:${BUILD_NUMBER} ${DOCKER_CREDENTIALS_USR}/products-update-service:latest
                    docker tag products-delete-service:${BUILD_NUMBER} ${DOCKER_CREDENTIALS_USR}/products-delete-service:latest
                    
                    docker push ${DOCKER_CREDENTIALS_USR}/products-create-service:latest
                    docker push ${DOCKER_CREDENTIALS_USR}/products-read-service:latest
                    docker push ${DOCKER_CREDENTIALS_USR}/products-update-service:latest
                    docker push ${DOCKER_CREDENTIALS_USR}/products-delete-service:latest
                    
                    echo "✅ Imágenes publicadas"
                '''
            }
        }
        
        stage('Generate Reports') {
            steps {
                echo '📊 Generando reportes...'
                sh '''
                    echo "==================================="
                    echo "📊 REPORTE DE PIPELINE"
                    echo "==================================="
                    echo "Build Number: ${BUILD_NUMBER}"
                    echo "Branch: ${GIT_BRANCH}"
                    echo "Commit: ${GIT_COMMIT}"
                    echo "==================================="
                '''
            }
        }
    }
    
    post {
        always {
            echo '🧹 Limpiando recursos...'
            sh '''
                # Detener y eliminar contenedor de MongoDB
                docker stop mongodb-jenkins || true
                docker rm mongodb-jenkins || true
                
                # Limpiar imágenes antiguas (opcional)
                docker image prune -f
            '''
            
            // Publicar reportes de cobertura
            publishHTML([
                allowMissing: true,
                alwaysLinkToLastBuild: true,
                keepAll: true,
                reportDir: 'services/create-service',
                reportFiles: 'coverage.html',
                reportName: 'Coverage Report - Create Service'
            ])
        }
        
        success {
            echo '✅ Pipeline ejecutado exitosamente!'
            // Aquí puedes agregar notificaciones (email, Slack, etc.)
        }
        
        failure {
            echo '❌ Pipeline falló!'
            // Notificaciones de error
        }
    }
}
