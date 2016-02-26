#!/bin/bash

FRONTEND=104.197.52.35:80

curl http://${FRONTEND}/rpush/wombat/red
curl http://${FRONTEND}/rpush/wombat/green
curl http://${FRONTEND}/rpush/wombat/blue
curl http://${FRONTEND}/lrange/wombat
