#!/bin/bash

# Command line argument for operation type
operation=$1
# Database name
dbname="Night_Service"
# Backup directory
backup_dir="/recycle"
# Current datetime
datetime=$(date '+%Y%m%d-%H%M%S')
# File type
type="mongodb"

# Check if the backup directory exists, if not, create it
sudo mkdir -p "$backup_dir"

case $operation in
  backup)
    # Construct the backup file path
    backup_file="$backup_dir/db_dump.$type.$datetime"
    # Run mongodump to backup the database
    sudo mongodump --db=$dbname --out=$backup_file
    echo "Backup successful: $backup_file"
    ;;
  restore)
    # Construct the restore file path (assuming the restore file is named with the type but without a datetime)
    restore_file="./db_backup.$type/Night_Service"
    # Run mongorestore to restore the database
    sudo mongorestore --db=$dbname $restore_file
    echo "Restore successful from: $restore_file"
    ;;
  *)
    echo "Invalid operation: $operation"
    echo "Usage: $0 [backup|restore]"
    exit 1
    ;;
esac
