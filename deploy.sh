#!/bin/bash

# this is your part to fill

# Deploy invoice-app
echo -n "Deploy invoice-app -> " && kubectl apply -f invoice-app/deployment.yaml

# Deploy payment-provider
echo -n "Deploy payment-provider -> " && kubectl apply -f payment-provider/deployment.yaml