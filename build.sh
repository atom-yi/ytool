#!/bin/bash

echo "build start..."

# target
rm -rf target
mkdir target

# build webapp
cd webapp || exit
npm run build
mv dist ../target/
cd ../target || exit
mv dist webapp
cd ../

# build service
cd service || exit
go build
mv ytool ../target/
mv ytool.exe ../target/

echo "build end."