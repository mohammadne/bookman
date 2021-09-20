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
        replacedText = replacedText.replace('${{ SERVICE }}', config.service)
        replacedText = replacedText.replace('${{ ENV }}', config.env)

    with open(fileName, "wt") as file:
        file.write(replacedText)
