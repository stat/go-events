#!/bin/sh

while inotifywait -qqre modify ./; do
  make test
done
