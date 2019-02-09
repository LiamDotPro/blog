#!/bin/sh

docker run -it --rm --name cs-blog -p 8080:8080 -v $HOME/code/blog:/go/src/blog -w /go/src/blog choskyo-blog