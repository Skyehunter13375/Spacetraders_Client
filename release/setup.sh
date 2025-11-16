#!/bin/bash
# ┣━━━━━━━━━━━━━━━━━━┫ Get sudo up-front for formatting ease ┣━━━━━━━━━━━━━━━━━━━┫
sudo printf ""

# ┣━━━━━━━━━━━━━━━━━━━┫ Install required packages if missing ┣━━━━━━━━━━━━━━━━━━━┫
if rpm -q golang 2>&1 | grep -q "not installed"; then
    printf "Installing Golang..."
    sudo dnf install -qy golang >/dev/null 2>&1
    printf "Done\n"
fi

if rpm -q postgresql-server 2>&1 | grep -q "not installed"; then
    printf "Installing PostgreSQL Server..."
    sudo dnf install -qy postgresql-server >/dev/null 2>&1
    printf "Done\n"
fi

# ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Store account token ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
# printf "Enter your ACCOUNT token: "
# read acct_token
# psql -d spacetraders -c "INSERT INTO tokens (type,token) VALUES ('account', '${acct_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"

# ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Store agent token ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
printf "Enter your AGENT token: "
read agent_token
psql -d spacetraders -c "INSERT INTO tokens (type,token) VALUES ('agent', '${agent_token}') ON CONFLICT (type) DO UPDATE SET token = EXCLUDED.token"

# ┣━━━━━━━━━━━━━━━━━━━━━━━━┫ Build and execute program ┣━━━━━━━━━━━━━━━━━━━━━━━━━┫
# printf "Building executable binary..."
# go build -o SpaceTraders.tui ../src/main.go
# printf "Done\n"

printf "All done..."