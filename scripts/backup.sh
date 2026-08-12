#!/usr/bin/env bash

set -e

set -a
source .env
set +a

BACKUP_DIR="./backups"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.sql"

mkdir -p "$BACKUP_DIR"

echo "[$(date '+%H:%M:%S')] Starting PostgreSQL backup..."

docker compose exec -T postgres \
  pg_dump \
  -U "$DB_USER" \
  -d "$DB_NAME" \
  > "$BACKUP_FILE"

echo "[$(date '+%H:%M:%S')] Backup created: $BACKUP_FILE"

find "$BACKUP_DIR" -type f -name "*.sql" -mtime +7 -delete