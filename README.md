# ingress-consul-register

A Tool to register the hostnames that are specified as Rules in an Ingress resource with the Ingress controller in Consul.

[![Build Status](https://travis-ci.org/aswinkarthik93/ingress-consul-register.svg?branch=master)](https://travis-ci.org/aswinkarthik93/ingress-consul-register)

## Description

Typically, Ingress and Ingress Controllers can be used to route traffic to the underlying service.

- The `Ingress` resource will be used to define rules as following:

```yaml
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - backend:
          serviceName: s1
          servicePort: 80
```

An Ingress Controller will process this rules and serve the request appropriately. The domain `foo.bar.com` with be associated with the public IP address of the Ingress Controller. When multiple such rules are used, a name based virtual hosting can be achieved.

- `ingress-consul-register` watches for such `Ingress` resources of an `IngressController` and registers the `host` of Ingress and the `ipaddress` of the IngressController in Consul.


## Inner working

Lets take an example,

- Imagine there are 2 services, `s1` and `s2`
- The following Ingress resource is used to define the rules

```yaml
spec:
  rules:
  - host: s1.service.consul
    http:
      paths:
      - backend:
          serviceName: s1
          servicePort: 80
  - host: s2.service.consul
    http:
      paths:
      - backend:
          serviceName: s2
          servicePort: 80
```

- Lets assume an IngressController is deployed and it runs with following IP address `10.104.244.148` (Let's assume cluster IP for now)
- `ingress-consul-register` will watch for this and do the following `PUT` request to consul

```JSON
{
  "ID": "...",
  "Name": "...",
  "Tags": [
    "s1",
    "s2"
  ],
  "Address": "10.104.244.148",
  "Check": "...",
  "Checks": "..."
}
```

## Use cases

- Use an internal `IngressController` for all internal inter-service communication.

If you want all your internal services to make use of the Internal Ingress controller for communication (over Kubernetes Service, for some reason), you could configure your `kube-dns` to resolve domains that end `.service.consul` against the Consul.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
data:
  stubDomains: |
    {"service.consul": ["${CONSUL_IPADDR}"]}
```

This would mean `s1.service.consul` would resolve to the IP address of the controller (10.104.244.148), and the ingress controller would serve the request.

- Exposing internal services within a VPN

There might be cases where you would need to expose the services within a VPN to developers, internal apps. Instead of manually adding to an Internal DNS that you would have in your internal infrastructure, the following can be done. Make the `IngressController` accessible via a VPN eg. using Google's Internal Loadbalancer. Expose the Consul DNS Service (Runs on port 8600) as a UDP DNS Service in Kubernetes (with an Internal Loadbalancer IP). Now make your VPN server push the ${KUBE_SVC_CONSUL_DNS_INTERNAL_IP} as the DNS Resolver. Which means anyone connecting to the VPN, will have access to service `s1` through the DNS entry `s1.service.consul`. As more `Ingress` resources are added, `ingress-consul-register` will watch for it and register it to consul. As `s3` service is deployed, a new DNS entry is exposed `s3.service.consul`.
