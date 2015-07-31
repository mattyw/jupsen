#!/bin/sh
set -ev
juju ssh mongodb/0 "mongo --eval 'printjson(db.isMaster())'"
juju ssh mongodb/1 "mongo --eval 'printjson(db.isMaster())'"
juju ssh mongodb/2 "mongo --eval 'printjson(db.isMaster())'"
juju ssh mongodb/3 "mongo --eval 'printjson(db.isMaster())'"
