#!/bin/sh
set -ev

juju jupsen part mongodb/0 mongodb/2
juju jupsen part mongodb/0 mongodb/3
juju jupsen part mongodb/1 mongodb/2
juju jupsen part mongodb/1 mongodb/3
read -p "Hit the any key to start healing"
juju jupsen heal mongodb/0 mongodb/2
juju jupsen heal mongodb/0 mongodb/3
juju jupsen heal mongodb/1 mongodb/2
juju jupsen heal mongodb/1 mongodb/3
