#!/bin/sh
set -ev

#Split brain
N1=mongodb/1
N2=mongodb/2
N3=mongodb/3
N4=mongodb/4
N5=mongodb/0

for i in $N1 $N2 $N3 $N4 $N5; do
	for j in $N1 $N2 $N3 $N4 $N5; do
		if [ "$i" != "$j" ]; then
			juju jupsen heal $i $j || true
		fi
	done
done

