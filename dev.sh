#!/bin/sh

docker build -t choskyo-blog .
docker run -it --rm --name cs-blog -p 8080:8080 choskyo-blog