#!/bin/bash

if [ -f backend/.env ]; then
    source backend/.env
fi

cd backend/sql/schema
goose turso $DB_URL up
