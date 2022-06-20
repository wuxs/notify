#!/usr/bin/env bash

docker buildx create --use --name=notify --driver docker-container --driver-opt image=dockerpracticesig/buildkit:master

docker buildx use notify

docker buildx build --platform linux/arm/v7,linux/arm64,linux/amd64 -t wuxs/notify . --push

docker buildx imagetools inspect wuxs/notify