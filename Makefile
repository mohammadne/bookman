# build docker
build-dockerfiles:
	cd scripts && python3 build-dockerfile.py build auth user library
