#!/usr/bin/env bash

set -x

rm -fr output/detail
mkdir -p output/detail

while read line
do
	./food  get $line   > output/detail/$line.json
done < output/catagory/shicai_list.txt
