#!/usr/bin/env python3.9
#package main

import sys
import ast
from ast import *

import requests
r =requests.get('https://xkcd.com/1906/')
print(r.status_code)
