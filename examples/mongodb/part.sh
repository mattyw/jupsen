#!/bin/sh
set -ev

#Split brain
N1=mongodb/4

N2=mongodb/1
N3=mongodb/2
N4=mongodb/3
N5=mongodb/0

#Partition N1 & N2 from the others
juju jupsen part $N1 $N2
juju jupsen part $N1 $N3
juju jupsen part $N1 $N4
juju jupsen part $N1 $N5
juju jupsen part $N2 $N1
juju jupsen part $N2 $N3
juju jupsen part $N2 $N4
juju jupsen part $N2 $N5
sleep 60
#Heal!
juju jupsen heal $N1 $N2
juju jupsen heal $N1 $N3
juju jupsen heal $N1 $N4
juju jupsen heal $N1 $N5
juju jupsen heal $N2 $N1
juju jupsen heal $N2 $N3
juju jupsen heal $N2 $N4
juju jupsen heal $N2 $N5
