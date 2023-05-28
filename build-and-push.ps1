#! /usr/bin/pwsh

$tag = Read-Host -Prompt "<tag>:"

IF ($tag -eq "") {
    Write-Host "Usage: <tag>"
    Exit
}

go build
Write-Host "Docker build with tag: $tag"
docker build -t registry.shadowchef.co/shadowchef/storage:$tag .
docker push registry.shadowchef.co/shadowchef/storage:$tag
