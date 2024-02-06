#!/bin/bash

echo "Your first argument is $0" 
# it will give the name of the file ha ha
echo "Your second argument is $1"
# it will start from the original arguments

for argument in $@
do
    echo "argument is $argument"
done
