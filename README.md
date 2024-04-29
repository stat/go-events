# Status
It "compiles", untested

# Getting Started

## Prereqs

We need to make sure we have the following installed:
```
go 1.22
docker
```

go can be installed and managed with GVM https://github.com/moovweb/gvm

Using Homebrew:

```
brew install --cask docker
```

## Running the tests

First copy .env.local to .env so that it may be used with docker compose and make commands

Next make sure to spin up the local env using:
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

### Local
```
== RUN   TestPipelinePerformance
-----
duration: 9.521344489s
sent: 10000000
 => 1050271.840448 events/second
-----
--- PASS: TestPipelinePerformance (9.52s)

```
### Redis
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

### Spinning it up
```
make compose-up
```

### Spinning it down
```
make compose-down
```

### Running the Tests
```
make test
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
* test coverage
* implement cassandra as the raw stream data store for the tape w/received_at
* implmement postgres as the persistant cached data store against redis
* implemment data expirations/cleaners within redis
* swaggo annotations
* technical diagram
