#! /bin/bash

uwsgi --http :9000 --wsgi-file main.py -L
