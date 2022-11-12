#!/bin/sh

# Get project root path
SOURCE="${BASH_SOURCE[0]}"

while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done

ROOT_PATH="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && cd .. && pwd )"

echo "ROOT_PATH=$ROOT_PATH"

# Determine if we need to install swag

which swag > /dev/null

if [ "$?" != "0" ] ; then
  # Install Swag

  echo "Swag is not installed. Installing Swag..."

  export GO111MODULE=on

  go get github.com/swaggo/swag/cmd/swag

  export GO111MODULE=off
fi

# Move to main path
cd "${ROOT_PATH}"

# Regenerate Swagger docs
echo " - Generating Swagger docs..."

swag init -d "${ROOT_PATH}/cmd/api/" --parseDependency -o "${ROOT_PATH}/internal/docs"

if [ "$?" != "0" ] ; then
  exit 1
fi

echo " - Moving YAML and JSON specs to the official documentation directory..."

mv "${ROOT_PATH}/internal/docs/swagger.json" "${ROOT_PATH}/internal/docs/swagger.yaml" "${ROOT_PATH}/docs/specs/"

echo " - Done!"