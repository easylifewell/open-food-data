#!/usr/bin/env bash

while read line
do
	dir=`dirname ${line:7}`
	echo "download $line"
	mkdir -p $dir
	cd $dir
	wget -c $line  
	cd -
done < pictures.txt
