#!/bin/sh
set -ev

N0=mongodb/0
N1=mongodb/1
N2=mongodb/2
N3=mongodb/3
N4=mongodb/4

for i in $N0 $N1 $N2 $N3 $N4; do
	for j in $N0 $N1 $N2 $N3 $N4; do
		if [ "$i" != "$j" ]; then
			juju jupsen heal $i $j || true
		fi
	done
done

