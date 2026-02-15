#!/bin/bash

if [ -z "$1" ]; then
    docker compose exec -it postgres psql -U cozy -d cozy_listings
else
    docker compose exec -t postgres psql -U cozy -d cozy_listings -c "$1"
fi
