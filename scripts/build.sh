#!/bin/bash
count=10
echo "Compiling $count files"
for i in `seq $count`; do
  sleep 1
  echo "Compiling $i..."
done
echo "Done"
