#!/bin/bash

FRONTEND=:3000

curl http://${FRONTEND}/rpush/wombat/red
curl http://${FRONTEND}/rpush/wombat/green
curl http://${FRONTEND}/rpush/wombat/blue
curl http://${FRONTEND}/lrange/wombat
