# Getting Started

First make sure to spin up the local env using:
```
make compose-up
```

Once running, we can execute the tests using:
```
make test
```

The consumer and producer pipeline is tested using Redis as the queue and data backend.

## Performance

Threads: 10

```
== RUN   TestPipelinePerformance
-----
duration: 47.794250481s
sent: 1000000
 => 20923.018772 events/second
-----
--- PASS: TestPipelinePerformance (47.79s)
```

## Commands

### Running the Tests
```
make test
```

### Spinning it up
```
make compose-up
```

### Spinning it down
```
make compose-down
```

<!-- ### Create A Migration -->
<!-- ``` -->
<!-- make migrate-create n=name -->
<!-- ``` -->

<!-- #### Example -->
<!-- ``` -->
<!-- make migrate-create name=events -->
<!-- ``` -->

<!-- ### Manually Migrating Up -->
<!-- ``` -->
<!-- make migrate-up -->
<!-- ``` -->
<!-- *note: this will migrate to the latest -->

<!-- ### Manually Migrating Down -->
<!-- ``` -->
<!-- make migrate-down -->
<!-- ``` -->
<!-- *note: this will migrate down one step, n=1 -->

<!-- ### Manualy Forcing a Migration -->
<!-- ``` -->
<!-- make migrate-force v=$(v) -->
<!-- ``` -->

<!-- #### Example -->
<!-- ``` -->
<!-- make migrate-force v=2 -->
<!-- ``` -->


TODO
* documentation
* tests
* implement lat/long comp algo
* re-implement DLQ data backend for Redis
* re-implement local in-memory data backends
