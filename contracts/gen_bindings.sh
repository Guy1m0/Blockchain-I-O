#!/bin/bash

# Function to generate Go bindings from Solidity contract
generate_bindings() {
    local contract_name="$1"
    local folder_name="$2"

    echo "Compiling Solidity contract for $contract_name..."

    # Go to the contract subdirectory
    cd "$folder_name" || exit

    # Compile Solidity contract
    solc --bin --abi --optimize -o build "$contract_name.sol" --overwrite

    # Check if solc compilation was successful
    if [ $? -ne 0 ]; then
        echo "Solidity compilation failed for $contract_name"
        exit 1
    fi

    echo "Generating Go bindings..."

    # Generate Go bindings
    ../abigen --bin="./build/${contract_name}.bin" --abi="./build/${contract_name}.abi" --pkg="$folder_name" --out="../${folder_name}/${contract_name}_gen.go"

    # Check if abigen was successful
    if [ $? -ne 0 ]; then
        echo "Failed to generate Go bindings for $contract_name"
        exit 1
    fi

    echo "Go bindings generated successfully."

    # Go back to the contracts directory
    cd ..
}

# Main script execution starts here

# Check if a directory name was provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <contract-subdirectory>"
    exit 1
fi

contract_subdirectory="$1"

# Check if the subdirectory exists
if [ ! -d "$contract_subdirectory" ]; then
    echo "Subdirectory $contract_subdirectory does not exist"
    exit 1
fi

# Find the Solidity contract file
contract_file=$(find "$contract_subdirectory" -name "*.sol" -print -quit)

# Extract the contract name
contract_name=$(basename "$contract_file" .sol)

# Call the function to generate bindings
generate_bindings "$contract_name" "$contract_subdirectory"
