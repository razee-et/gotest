#!/bin/bash
if [ $# -lt 1 ] ; then
  echo ""
  echo "Usage:"
  echo "  protoc.sh [path to acquia-protobufs repo root]"
  echo ""
  exit 1
fi

REPO_ROOT="$( cd "$(dirname "$0")/.." || exit ; pwd -P )"
MESSAGES_PATH="$REPO_ROOT/pkg/acquia/messages"
PROTOLIST="$1/protolist"

if [ ! -f "$PROTOLIST" ] ; then
  echo "protolist file not found at '$PROTOLIST'"
  exit 1
fi

# CAREFUL!
rm -r "$MESSAGES_PATH"

if [ -d "$MESSAGES_PATH" ] ; then
  echo "WARNING: $MESSAGES_PATH was not successfully cleaned up"
fi

while IFS="" read -r p || [ -n "$p" ]
do
  echo "$p"
  protoc --go_out="paths=source_relative:$REPO_ROOT/pkg" -I="$1" "$p"
done < "$PROTOLIST"
