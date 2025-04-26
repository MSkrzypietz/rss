#!/bin/bash

prod=false
positional_args=()
child_args=()

if [ -f backend/.env.migrate ]; then
    source backend/.env.migrate
fi

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --prod) prod=true ;;
        --) shift; while [[ "$#" -gt 0 ]]; do child_args+=("$1"); shift; done; break ;;
        -*) echo "Unknown option: $1"; exit 1 ;;
        *) positional_args+=("$1"); ;;
    esac
    shift
done

cd backend/sql/schema

db_url=$DB_URL_TEST
if [ "$prod" = true ]; then
  db_url=$DB_URL_PROD
fi

goose turso $db_url "${positional_args[@]}" "${child_args[@]}"
