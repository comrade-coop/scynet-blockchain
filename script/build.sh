#!/usr/bin/env bash
set -e
set -o pipefail

cd `dirname $0`/../

function prefix {
	exec "${@:2}" 2>&1 | sed "s/^/[$(tput setaf 3 || true)$1$(tput sgr0 || true)] /"
}

prefix "DAE" go build -o build/scynetd cmd/scynetd/main.go
prefix "CLI" go build -o build/scynetcli cmd/scynetcli/main.go
prefix "COM" go generate ./cmd/scynetcomponent/
prefix "COM" go build -o build/scynetcomponent cmd/scynetcomponent/main.go
