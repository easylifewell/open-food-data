#!/usr/bin/env bash

set -x

rm -fr output/xianguo
mkdir -p output/xianguo

while read line
do
	./main  $line   > output/xianguo/$line.json
done < xianguo_list.txt
