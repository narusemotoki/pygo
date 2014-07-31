#! /bin/bash

ab -n 5000 -c 1 -p test.json -T "application/json; charset=utf-8" http://localhost:9000/Count
