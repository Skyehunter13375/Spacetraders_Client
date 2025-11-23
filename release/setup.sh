#!/bin/bash
# ┣━━━━━━━━━━━━━━━━━━┫ Get sudo up-front for formatting ease ┣━━━━━━━━━━━━━━━━━━━┫
sudo printf ""


# ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Creating log dir ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
sudo mkdir /var/log/spacetraders >/dev/null 2>&1
sudo chown $(whoami):$(whoami) /var/log/spacetraders
sudo chmod 777 /var/log/spacetraders
touch /var/log/spacetraders/activity.log
touch /var/log/spacetraders/error.log
echo "Logs are found here: /var/log/spacetraders"


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


# ┣━━━━━━━━━━━━━━━━━━━━━━━━┫ Build and execute program ┣━━━━━━━━━━━━━━━━━━━━━━━━━┫
# printf "Building executable binary..."
# go build -o SpaceTraders.tui ../src/main.go
# printf "Done\n"

echo "IMPORTANT: Make sure you build the DB and update the config_sample.yaml file!!!!"
echo
echo "All done..."
