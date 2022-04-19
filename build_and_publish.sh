#!/bin/bash

sh build.sh $1
build_out=$(docker images --filter=reference="blood-bank-system-service:$1" | grep "blood-bank-system-service")
if [[ -z "$build_out" ]]
then
  echo "Exiting since build failed"
  exit 1
fi

parsed_tag="blood-bank-system-service:$1"
echo "Parsed tag is $parsed_tag"

aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin ecr_link
docker tag "$parsed_tag" ecr_link//blood-bank-system-service:$1
docker push ecr_link//blood-bank-system-service:$1
