#Jupsen

## Introduction

Jupsen is a shameless copy of Jepsen: https://github.com/aphyr/jepsen tailored for juju environments https://github.com/juju/juju.

Jupsen is a juju plugin that includes a number of sub commands that are used to cause problems in juju networks.
You can see it in action here https://www.youtube.com/watch?v=xp54gcMMUfQ

As an example you can create a partition between a wordpress and mysql unit by running part:
```
juju jupsen part wordpress/0 mysql/0
```

You can fix this partition by running
```
juju jupsen heal wordpress/0 mysql/0
```

## Install

```
go install github.com/mattyw/jupsen/...
```

## Running
```
juju jupsen help
```
