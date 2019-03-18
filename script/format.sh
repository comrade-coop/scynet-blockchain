#!/usr/bin/env bash
set -e
set -o pipefail

cd `dirname $0`/../

function prefix {
	exec "${@:2}" 2>&1 | sed "s/^/[$(tput setaf 3 || true)$1$(tput sgr0 || true)] /"
}

prefix "FMT" go fmt ./
prefix "IMP" goimports -local github.com/comrade-coop -w .
