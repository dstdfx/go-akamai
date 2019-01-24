#!/usr/bin/env bash

echo "==> Running go test and creating a coverage profile..."
i=0
for testingpkg in $(go list ./akamai/.../testing); do
    covpkg=${testingpkg%/testing}
    go test -v -covermode count -coverprofile "./${i}.coverprofile" -coverpkg $covpkg $testingpkg
    ((i++))
done
gocovmerge $(ls ./*.coverprofile) > coverage.out
rm *.coverprofile
