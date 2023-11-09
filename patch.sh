#!/bin/bash
API_GATEWAY="13.212.68.212"

awk '{gsub(/localhost/,"'"$API_GATEWAY"'");gsub(/go-gonojobs-template/,"REMOTE-go-gonojobs-template")}1' docs/go-gonojobs-template.postman_collection.json > docs/REMOTE-go-gonojobs-template.postman_collection.json