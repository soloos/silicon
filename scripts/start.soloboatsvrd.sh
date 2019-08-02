#!/bin/bash
nohup bin/soloboatd ./scripts/conf/soloboat.json > ./logs/soloboat.log 2>&1 &
