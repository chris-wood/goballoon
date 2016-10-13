#!/bin/bash

PROGRAM=$1
FOUT=$2
PASSWORD=$3
SALT=`dd count=8 bs=8 if=/dev/urandom 2> /dev/null | base64`

echo ${PASSWORD}
echo ${SALT}

SPACE_VALUES=( 1024 2048 4096 8192 16384 32768 65536 )
TIME_VALUES=( 1 2 8 16 128 256 1024 2048 4096 )

for s_cost in "${SPACE_VALUES[@]}"
do
    for t_cost in "${TIME_VALUES[@]}"
    do
        echo go run ${PROGRAM} ${PASSWORD} ${SALT} ${s_cost} ${t_cost}
        go run ${PROGRAM} ${PASSWORD} ${SALT} ${s_cost} ${t_cost} >> ${FOUT} 
    done
done
