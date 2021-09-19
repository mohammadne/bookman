
#!/bin/bash

render() {
sedStr="
  s!%%RUBY_VERSION%%!$version!g;
"

sed -r "$sedStr" $1
}

versions=(2.1 2.3)
for version in ${versions[*]}; do
  mkdir $version
  render Dockerfile.template > $version/Dockerfile
done