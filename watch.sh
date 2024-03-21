#!/bin/sh

while inotifywait -qqre modify ./; do
  make clean && make test
done
