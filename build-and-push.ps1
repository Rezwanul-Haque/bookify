#! /usr/bin/pwsh

$tag = Read-Host -Prompt "<tag>:"

IF ($tag -eq "") {
    Write-Host "Usage: <tag>"
    Exit
}

go build
Write-Host "Docker build with tag: $tag"
docker build -t <registry.url>/<username>/bookify:$tag .
docker push <registry.url>/<username>/bookify:$tag
