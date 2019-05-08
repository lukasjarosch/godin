#!/usr/bin/env bash
pushd ../../ > /dev/null && go install && popd > /dev/null && godin $1
