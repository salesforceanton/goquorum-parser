#!/bin/bash
cd ./contracts
for f in *.sol
do
	echo "Compiling $f contract..."
	C:/reps/solc/0.4.17/solc --bin --abi --optimize --overwrite -o ./output $f
done