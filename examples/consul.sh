#!/bin/bash -el

helm install --name dns stable/consul --set "antiAffinity=soft"
