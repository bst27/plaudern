#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

trap 'kill %1' SIGINT

$DIR/bin/linux-amd64/plaudern serve --config="$DIR/examples/default-config/plaudern-config.json" | sed -e 's/^/[Backend] /' &
(cd web; ng serve --host 0.0.0.0 --disableHostCheck --open | sed -e 's/^/[Angular] /')

trap - SIGINT