#!/bin/bash -ex

if [ -z "$DELAY" ]; then
	DELAY=60
fi

N0=mongodb/0
N1=mongodb/1
N2=mongodb/2
N3=mongodb/3
N4=mongodb/4

GROUP1=
GROUP2=

if [ -z "$PRIMARY" ]; then
	echo "PRIMARY unset (use 'go run primary.go' to find it)"
	exit 1
fi

GROUP1=

for i in $N0 $N1 $N2 $N3 $N4; do
	if [ "$PRIMARY" != "$i" ]; then
		if [ -z "$GROUP1" ]; then
			GROUP1="$GROUP1 $i"
		else
			GROUP2="$GROUP2 $i"
		fi
	fi
done

GROUP1="$GROUP1 $PRIMARY"

echo "GROUP1=$GROUP1"
echo "GROUP2=$GROUP2"

echo "partitioning..."

# Partition GROUP1 from GROUP2
for g1 in $GROUP1; do
	for g2 in $GROUP2; do
		juju jupsen part $g1 $g2
		juju jupsen part $g2 $g1
	done
done

echo -n "partitioned"

sleep $DELAY

echo -n "healing..."

# Heal GROUP1 & GROUP2
for g1 in $GROUP1; do
	for g2 in $GROUP2; do
		juju jupsen heal $g1 $g2 || true
		juju jupsen heal $g2 $g1 || true
	done
done

echo "healed"

