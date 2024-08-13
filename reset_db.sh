#!/bin/bash

CONTAINER_NAME="facebook-clone-db"
USERNAME="postgres"
DATABASE="facebook"
INIT_SCRIPT="/docker-entrypoint-initdb.d/init.sql"

docker exec -it $CONTAINER_NAME psql -U $USERNAME -d $DATABASE -c "DO $$ DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP EXECUTE 'DROP TABLE IF EXISTS public.' || r.tablename || ' CASCADE'; END LOOP; END $$;"

docker exec -it $CONTAINER_NAME psql -U $USERNAME -d $DATABASE -f $INIT_SCRIPT