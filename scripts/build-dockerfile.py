#!/usr/bin/python
from shutil import copyfile
import sys
import os

buildDirectory = sys.argv[1]
services = sys.argv[2:]

template = f"{buildDirectory}/template.txt"

for service in services:
    outputDir = f"{buildDirectory}/{service}"
    if not os.path.exists(outputDir):
        os.mkdir(outputDir)

    fileName = copyfile(template, f"{outputDir}/Dockerfile")

    with open(fileName, "rt") as file:
        content = file.read()

    content = content.replace('${{ SERVICE }}', service)
    content = content.replace('${{ ENV }}', "prod")

    with open(fileName, "wt") as file:
        file.write(content)
