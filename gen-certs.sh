#!/bin/bash

if ! command -v cockroach &> /dev/null
then
    echo "cockroach could not be found in your PATH. Please install it and add it to your PATH: https://www.cockroachlabs.com/docs/v22.2/install-cockroachdb-linux.html"
    exit
fi

source .env
if [[ "" = "$SERVER_ALIASES" ]]
then 
    echo "ERROR: SERVER_ALIASES is undefined in your .env file. Please state all your possible hostnames and ip-addresses of your database node";
    exit;
fi

mkdir -p certs
mkdir -p priv-certs

if ! [ -z "$(ls -A certs/)" ]; then
    echo "ERROR: Your certs directory isn't empty. To generate new certificates backup your old folders and empty it.";
    echo "Your current certificates are: "
    cockroach cert list --certs-dir=certs
    exit;
fi
if ! [ -z "$(ls -A priv-certs/)" ]; then
    echo "ERROR: Your priv-certs directory isn't empty. To generate new certificates backup your old folders and empty it.";
    echo "Your current certificates are: "
    cockroach cert list --certs-dir=certs
    exit;
fi

echo "Generating Certificates:";
cockroach cert create-ca --certs-dir=certs --ca-key=priv-certs/ca.key
cockroach cert create-node $SERVER_ALIASES --certs-dir=certs --ca-key=priv-certs/ca.key
echo "Done generating. That are all your certificates:"
cockroach cert list --certs-dir=certs
