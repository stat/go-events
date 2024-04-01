#!/usr/bin/env bash

until printf "" 2>>/dev/null >>/dev/tcp/cassandra-00/9042; do 
  sleep 1;
  echo "Waiting for cassandra host...";
done

echo "Creating keyspace..."
cqlsh cassandra-00 -e "CREATE KEYSPACE IF NOT EXISTS $CASSANDRA_KEYSPACE WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"
echo "Complete"
