#!/bin/bash

# Execute the scripts in sequence
./stop.sh
./remove.sh
./init.sh
./start.sh

echo "All scripts executed successfully."
