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

### Create A Migration
```
make migrate-create n=name
```

#### Example
```
make migrate-create name=events
```

### Manually Migrating Up
```
make migrate-up
```
*note: this will migrate to the latest

### Manually Migrating Down
```
make migrate-down
```
*note: this will migrate down one step, n=1

### Manualy Forcing a Migration
```
make migrate-force v=$(v)
```

#### Example
```
make migrate-force v=2
```


### Thoughts
* PostGIS
* additional db constraints

### Notes
* I may have been a bit too verbose on this

TODO
* expand test suite
* implement better reconsiliation algos
* implement channels
* add go function header documentation
* continue adding to this readme
