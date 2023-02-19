#!/bin/bash
unset SECRETT
#echo "before pin"
#ls -la pin.txt
timeout 3 python3 main.py &> pin.txt
export SECRETT=$(tail -1 pin.txt | awk '{print $NF}')
#echo "done"
echo $SECRETT



