#!/usr/bin/env bash

set -e

sudo rm -rf /go
mkdir -p ${PWD}/.devcontainer/.go
sudo ln -s ${PWD}/.devcontainer/.go /go 
touch ${PWD}/.devcontainer/.bash_history
ln -sf ${PWD}/.devcontainer/.bash_history ${HOME}/.bash_history

mkdir -p ${PWD}/.devcontainer/.vscode-extensions

rm -rf ${HOME}/.vscode-server/extensions
mkdir -p ${HOME}/.vscode-server
ln -s ${PWD}/.devcontainer/.vscode-extensions ${HOME}/.vscode-server/extensions

rm -rf ${HOME}/.vscode-remote/extensions
mkdir -p ${HOME}/.vscode-remote
ln -s ${PWD}/.devcontainer/.vscode-extensions ${HOME}/.vscode-remote/extensions

rm -rf ${HOME}/.cache/go-build
mkdir -p ${PWD}/.devcontainer/.go-build-cache ${HOME}/.cache
ln -s ${PWD}/.devcontainer/.go-build-cache ${HOME}/.cache/go-build

go install github.com/golang/vscode-go/tools/installtools@latest
installtools