#!/usr/bin/env bash

go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -e
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/github.com/openkruise/controllermesh-api
cp -r ./{ctrlmesh,hack,vendor} "${TMP_DIR}"/src/github.com/openkruise/controllermesh-api/

(cd "${TMP_DIR}"/src/github.com/openkruise/controllermesh-api; \
    GOPATH=${TMP_DIR} GO111MODULE=off /bin/bash vendor/k8s.io/code-generator/generate-groups.sh all \
    github.com/openkruise/controllermesh-api/client github.com/openkruise/controllermesh-api "ctrlmesh:v1alpha1" -h ./hack/boilerplate.go.txt)

mkdir -p ./client
rm -rf ./client/{clientset,informers,listers}
mv "${TMP_DIR}"/src/github.com/openkruise/controllermesh-api/client/* ./client/
