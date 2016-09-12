#!/usr/bin/env bash

set -x

rm -fr output/zhurou
mkdir -p output/zhurou

while read line
do
	./main  $line   > output/zhurou/$line.json
done < zhurou_list.txt
