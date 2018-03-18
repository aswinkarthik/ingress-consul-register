#!/bin/bash -el

helm install --name dns stable/consul --set "antiAffinity=soft"
helm install stable/nginx-ingress --name router --set "controller.ingressClass=nginx-internal"
