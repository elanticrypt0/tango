#!/bin/bash

cd ./cli && go mod tidy
cd ../api && go mod tidy
cd ../tango_pkg && go mod tidy