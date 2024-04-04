#!/bin/sh

while inotifywait -qqre modify ./pkg; do
  make test
done
