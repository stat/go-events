#!/bin/sh

while inotifywait -qqre modify ./; do
  ENV=test make test
done
