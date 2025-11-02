#!/bin/bash
# Get sudo creds for install so it doesn't break alignment
sudo printf ""

# Installing packages
if rpm -q golang 2>&1 | grep -q "not installed"; then
    printf "Installing Golang..."
    sudo dnf install -qy golang >/dev/null 2>&1
    printf "Done\n"
fi

if rpm -q sqlite-devel 2>&1 | grep -q "not installed"; then
    printf "Installing SQLlite3..."
    sudo dnf install -qy sqlite-devel >/dev/null 2>&1
    printf "Done\n"
fi

# Setting up the DB files
if [[ ! -e SpaceTraders.db ]]; then
    printf "Creating SpaceTraders.db..."
    sqlite3 SpaceTraders.db < SpaceTraders.schema
    printf "Done\n"

    # Get Account and Agent data for the user and store
    printf "Enter your ACCOUNT token: "
    read acct_token

    sqlite3 SpaceTraders.db "INSERT INTO tokens (type,token) VALUES ('account', '${acct_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"

    printf "Enter your AGENT token: "
    read agent_token

    sqlite3 SpaceTraders.db "INSERT INTO tokens (type,token) VALUES ('agent', '${agent_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"
fi



# Build the program, show where it is and execute it for testing.
printf "Building executable binary..."
go build -o SpaceTraders.tui ../src/main.go
printf "Done\n"

sleep 1

./SpaceTraders.tui