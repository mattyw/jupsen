#!/bin/sh
set -ev
PRIMARY=mongodb/0
#Partition!
juju jupsen part $PRIMARY mongodb/1
juju jupsen part $PRIMARY mongodb/2
juju jupsen part $PRIMARY mongodb/3
juju jupsen part $PRIMARY mongodb/4
sleep 60
#Heal!
juju jupsen heal $PRIMARY mongodb/1
juju jupsen heal $PRIMARY mongodb/2
juju jupsen heal $PRIMARY mongodb/3
juju jupsen heal $PRIMARY mongodb/4
