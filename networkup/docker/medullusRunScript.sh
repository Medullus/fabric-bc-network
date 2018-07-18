#!/usr/bin/env bash

declare -a users=()
#create users
curl -X POST "http://leanblocks.eastus.cloudapp.azure.com/users" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"requestHeader\": {    \"caller\": \"AtlasAmericas\",    \"org\": \"Org1\"  },  \"userRegister\": {    \"secret\": \"defaultpw\",    \"userName\": \"AtlasAmaerica\"  }}"