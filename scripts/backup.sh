#!/bin/bash

# backup.sh

DATE=$(date +"%Y%m%d-%H%M")

backupDir="backups/backup-$DATE"
mkdir -p "$backupDir"

# Ejecutar mongodump dentro del contenedor
docker exec parcial_mongo mongodump --uri="mongodb://admin:admin123@mongo:27017" --out="/data/db/backup-$DATE"

# Copiar backup al host
docker cp "parcial_mongo:/data/db/backup-$DATE" "$backupDir"

echo "✅ Backup guardado en $backupDir"
