#!/bin/sh

docker build -t choskyo-blog .
docker run -it --rm --name cs-blog -p 80:8080 choskyo-blog