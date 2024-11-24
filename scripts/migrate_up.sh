#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd backend/sql/schema
goose turso $DB_URL up
