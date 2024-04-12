#!/bin/sh

while inotifywait -qqre modify ./app ./pkg; do
  make test
done
