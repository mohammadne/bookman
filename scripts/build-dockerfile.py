#!/usr/bin/python
from shutil import copyfile
import os


class Config:
    def __init__(self, service, port, env):
        self.service = service
        self.port = port
        self.env = env


configs = [
    Config("auth", "8080", "dev"),
    Config("user", "8081", "dev"),
    Config("library", "8082", "dev"),
]

pathToDir = "../build"
template = f"{pathToDir}/template.txt"

for config in configs:
    outputDir = f"{pathToDir}/{config.service}"
    if not os.path.exists(outputDir):
        os.mkdir(outputDir)

    fileName = copyfile(template, f"{outputDir}/Dockerfile")

    with open(fileName, "rt") as file:
        replacedText = file.read().replace('${{ EXPOSED_PORT }}', config.port)
        replacedText = replacedText.replace('${{ SERVICE }}', config.service)
        replacedText = replacedText.replace('${{ ENV }}', config.env)

    with open(fileName, "wt") as file:
        file.write(replacedText)
