#!/bin/bash
ps aux|grep silicon|grep bin|grep -v grep |awk '{print $2}'|xargs kill -SIGABRT 2>/dev/null
