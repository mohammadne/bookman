#!/usr/bin/python
from shutil import copyfile
import os


class Config:
    def __init__(self, service, env):
        self.service = service
        self.env = env


configs = [
    Config("auth", "prod"),
    Config("user", "prod"),
    Config("library", "prod"),
]

pathToDir = "../build"
template = f"{pathToDir}/template.txt"

for config in configs:
    outputDir = f"{pathToDir}/{config.service}"
    if not os.path.exists(outputDir):
        os.mkdir(outputDir)

    fileName = copyfile(template, f"{outputDir}/Dockerfile")

    with open(fileName, "rt") as file:
        content = file.read()

    content = content.replace('${{ SERVICE }}', config.service)
    content = content.replace('${{ ENV }}', config.env)

    with open(fileName, "wt") as file:
        file.write(content)
