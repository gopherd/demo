#!/bin/bash

function core:os() {
	echo $(uname)
}

function core:is_macos() {
	if [[ "$(uname)" == "Darwin" ]]; then
		return 0
	else
		return 1
	fi
}

function core:is_linux() {
	if [[ "$(uname)" == "Linux" ]]; then
		return 0
	else
		return 1
	fi
}
