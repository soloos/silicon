#!/bin/bash
nohup bin/silicond ./scripts/conf/silicon.json > ./logs/silicon.log 2>&1 &
