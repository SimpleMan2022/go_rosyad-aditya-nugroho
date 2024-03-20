#!/bin/sh

echo "Enter project name: "
read projectName

echo "Choose project type (simple, api): "
read projectType

if [ "$projectType" = "simple" ]; then
    echo "Creating $projectName ..."
    mkdir "$projectName"
    cd "$projectName" || exit
    go mod init "$projectName"
    touch main.go
    echo "Project $projectName created successfully"
elif [ "$projectType" = "api" ]; then
    echo "Creating $projectName ..."
    mkdir "$projectName"
    cd "$projectName" || exit
    go mod init "$projectName"
    mkdir routes models repositories services configs controllers
    touch main.go
   echo "Project $projectName created successfully"
else
    echo "Invalid project type"
fi
