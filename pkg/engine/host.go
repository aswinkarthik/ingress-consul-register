package engine

import (
	"fmt"
	"regexp"
	"strings"

	v1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

type hosts []*string

var hostRegexPattern *regexp.Regexp

func init() {
	hostRegexPattern, _ = regexp.Compile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")
}

func convertToHosts(ingresses []*v1beta1.Ingress) hosts {
	// Trying to precalculate array size
	sum := 0
	for _, ing := range ingresses {
		sum += len(ing.GetSpec().GetRules())
	}

	hosts := make([]*string, sum)
	counter := 0
	for _, ing := range ingresses {
		for _, rule := range ing.GetSpec().GetRules() {
			if rule.Host != nil {
				hosts[counter] = rule.Host
				counter++
			}
		}
	}
	return hosts[:counter]
}

func (h hosts) filterByDomain(domain string) hosts {
	validHosts := make([]*string, h.size())
	counter := 0
	for _, host := range h.items() {
		if hostRegexPattern.MatchString(*host) && strings.HasSuffix(*host, domain) {
			validHosts[counter] = host
			counter++
		}
	}
	return hosts(validHosts[:counter])
}

func (h hosts) getTags(consulDomain string, prefix string) []string {
	tags := make([]string, h.size())

	var domain string
	if !strings.HasPrefix(consulDomain, ".") {
		domain = fmt.Sprintf(".%s.%s", prefix, consulDomain)
	} else {
		domain = fmt.Sprintf(".%s%s", prefix, consulDomain)
	}

	for i, host := range h.items() {
		tags[i] = strings.Replace(*host, domain, "", -1)
	}
	return tags
}

func (h *hosts) size() int {
	return len([]*string(*h))
}

func (h *hosts) items() []*string {
	return []*string(*h)
}
