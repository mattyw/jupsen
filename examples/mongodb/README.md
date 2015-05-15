This example is desgined to be run on a mongo replicaset.

### Deploy 5 mongodb units

juju deploy mongodb -n 5

### Craft a connection URL

* Grab the IP address of one of the MongoDB units.
* Get the replica set name of the cluster. Default is "myset".
* Use <ip>:<port>?replicaSet=<replica set name> for the URL in the following steps.

Example:

* mongodb/2 is 10.0.3.166
* replica set name is "jupsen"
* URL is 10.0.3.166:27017?replicaSet=jupsen

The `replicaSet` name is necessary for mgo to discover the other replica set
members and fail over the primary during the following tests. Without it, the
client will hang for the duration of the partition -- it won't switch to a
newly elected primary.

### Run the mongodb updates

```
go run main.go -url <URL>
```

See `go run main.go -help` for options.

This will display the primary IP address. Cross-reference with `juju status` to
find the primary unit.

#### Find the current primary

You can also check the current primary while the test is running, if you've
created a partition, for example, with:

```
go run primary.go -url <URL>
```

Match the IP with `juju status` to find the primary unit.

### Cause a temporary partition in the replica set

`PRIMARY=<primary unit> ./part.sh`

This will create a partition between two groups. The first group containing the
given `$PRIMARY` and a secondary unit. The second group containing all the other
secondaries. Creating this parition should force an election in which two
primaries arise in each partition.

The partition will be healed after `$DELAY` seconds. Default is 60.

### See if the expected len is different from what is found in the db.

So far, lost writes are easy to reproduce with a partition and `go run main.go
-majority=false` (without write majority).

With `-majority=true`, writes from a _single client_ seem to be reliable. There
are some writes that actually happen, but are reported back as an error to the
client. However, for an "at _least_ once" operation, this is acceptable.

However, whether _multiple concurrent MongoDB client connections_ will see
consistent results through partitions and primary failovers is yet to be
determined.

Other operations (CaS, for example) are also untested. Stay tuned.

