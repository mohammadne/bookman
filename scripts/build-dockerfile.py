#!/usr/bin/python
from shutil import copyfile
import os


services = ["auth", "user", "library"]
environment = "prod"

pathToDir = "../build"
template = f"{pathToDir}/template.txt"

for service in services:
    outputDir = f"{pathToDir}/{service}"
    if not os.path.exists(outputDir):
        os.mkdir(outputDir)

    fileName = copyfile(template, f"{outputDir}/Dockerfile")

    with open(fileName, "rt") as file:
        content = file.read()

    content = content.replace('${{ SERVICE }}', service)
    content = content.replace('${{ ENV }}', environment)

    with open(fileName, "wt") as file:
        file.write(content)
