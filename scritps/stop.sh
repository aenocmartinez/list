#!/bin/bash

docker kill listapp-container
docker kill listapp-db-container

docker rmi listapp:dev
docker rmi mysql:5.7 -f