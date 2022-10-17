#!/usr/bin/env bash

set -e

go install github.com/golang/vscode-go/tools/installtools@latest
installtools
ln -sf ${PWD}/.devcontainer/.bash_history /commandhistory