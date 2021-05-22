#!/usr/bin/env bash
#!/bin/bash
IMAGE=$1

# get highest tag number
VERSION=$2

NEW_TAG="$VERSION"

#######   Processing  ###########
docker build -t $IMAGE:$NEW_TAG . &&  docker push $IMAGE:$NEW_TAG
echo "Rease new image with tag: $NEW_TAG"
#######   Processing  ###########

#get current hash and see if it already has a tag
echo "###############################################################"
