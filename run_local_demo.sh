#!/bin/bash

# This script starts the api locally
# and also compiles the front end and launches a local environment for it.

go run api/server.go &
echo "Waiting 3 secs to let server start"
sleep 3
npm --prefix ./front/ run dev && fg
echo "Running Front @ :8080"
