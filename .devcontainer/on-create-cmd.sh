#!/bin/sh

set -e

sudo chmod o+rw /var/run/docker.sock

rm -rf ${HOME}/.bash_history
touch ${PWD}/.devcontainer/bash_history
ln -s ${PWD}/.devcontainer/bash_history ${HOME}/.bash_history

sudo rm -rf /go
mkdir -p ${PWD}/.devcontainer/go
sudo ln -s ${PWD}/.devcontainer/go /go

mkdir -p ~/.cache
sudo rm -rf ~/.cache/go-build
mkdir -p ${PWD}/.devcontainer/go-build-cache
sudo ln -s ${PWD}/.devcontainer/go-build-cache ~/.cache/go-build

# Vscode Remote-Container extension uses 'vscode-remote'
if [ -d ~/.vscode-server ] ; then
    rm -rf ~/.vscode-server/extensions
    mkdir -p ${PWD}/.devcontainer/vscode-extensions ~/.vscode-server
    ln -s ${PWD}/.devcontainer/vscode-extensions ~/.vscode-server/extensions
fi

# Github Codespaces uses 'vscode-remote'
if [ -d ~/.vscode-remote ] ; then
    rm -rf ~/.vscode-remote/extensions
    mkdir -p ${PWD}/.devcontainer/vscode-extensions ~/.vscode-remote
    ln -s ${PWD}/.devcontainer/vscode-extensions ~/.vscode-remote/extensions
fi