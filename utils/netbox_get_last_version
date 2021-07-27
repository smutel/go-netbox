#!/bin/bash

VERSION="$1"
TAGS=""
TAGS_LIST=""
NEXT=""
MAX=0

while [ "$NEXT" != "null" ]
do
   i=$((i+1))
   JSON=$(curl https://registry.hub.docker.com/v2/repositories/netboxcommunity/netbox/tags/?page=$i 2>/dev/null)
   TAGS="$TAGS $(echo $JSON | jq '."results"[]["name"]')"
   NEXT=$(echo $JSON | jq '."next"')
done

for t in $TAGS; do
  tag="$(echo $t | xargs)"
  if [ "$(echo $tag | grep "$VERSION")" == "$tag" -a "$(echo $tag | grep -v "ldap")" == "$tag" -a "$(echo $tag | grep -v "develop")" == "$tag" ]; then
    TAGS_LIST="$TAGS_LIST $tag"
  fi
done

for t in $TAGS_LIST; do
  MAJOR=$(echo $t | cut -d"." -f3)
  if [ "$MAJOR" != "" ]; then
    if [ $MAJOR -gt $MAX ]; then
      MAX=$MAJOR
    fi
  fi
done

echo "$VERSION.$MAX"
