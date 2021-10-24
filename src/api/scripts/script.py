#!/usr/bin/env python3.9
#package main

import sys
import ast

st = ast.parse('5 + 6')
pstr = ast.dump(ast.parse('5 + q'),indent=4)

print(help(pstr))