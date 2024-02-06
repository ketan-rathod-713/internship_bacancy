#!/bin/bash

number=10

while [ $number -gt 1 ]
do 
    echo "while loop me number is $number"
    number=$(($number-1))
done

# for loop already done in taking inputs via arguments file