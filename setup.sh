#!/bin/bash
# Get sudo creds for install so it doesn't break alignment
sudo printf ""

# Installing packages
printf "Installing Golang..."
sudo dnf install -qy golang >/dev/null 2>&1
printf "Done\n"

printf "Installing SQLlite3..."
sudo dnf install -qy sqlite-devel >/dev/null 2>&1
printf "Done\n"

# Setting up the DB files
if [[ ! -e src/DB/spacetraders.db ]]; then
    printf "Creating src/DB/spacetraders.db..."
    sqlite3 src/DB/spacetraders.db < src/DB/spacetraders.schema
    printf "Done\n"
fi

# Get Account and Agent data for the user and store
printf "Enter your ACCOUNT token: "
read acct_token

sqlite3 src/DB/spacetraders.db "INSERT INTO tokens (type,token) VALUES ('account', '${acct_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"

printf "Enter your AGENT token: "
read agent_token

sqlite3 src/DB/spacetraders.db "INSERT INTO tokens (type,token) VALUES ('agent', '${agent_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"
