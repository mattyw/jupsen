This example is desgined to be run on a mongo replicaset.

 * Deploy 5 mongodb units
 * Run main.go specifying the urls to use
 * It will tell you which url is the primary node
 * Use part.sh to cause a partition around the specified unit while main.go is running. Set Primary to the primary unit (ip address output from main.go)
 * See if the expected len is different from what is found in the db.
