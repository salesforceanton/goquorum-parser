#!/bin/bash
cd ./contracts/v2
for f in *.sol
do
	echo "Compiling $f contract..."
	C:/reps/solc/solc --bin --abi --optimize --overwrite -o ./output $f
done