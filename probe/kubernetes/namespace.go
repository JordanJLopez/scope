package kubernetes

import (
	"github.com/weaveworks/scope/report"

	apiv1 "k8s.io/client-go/pkg/api/v1"
)

// GetNode returns a Namespace Node from a kubernetes namespace
//
// The GetNode function does not follow the conventions of other kubernetes api resources
// (see cronjob.go, daemonsets.go, deployments.go, pod.go, replica_set.go, service.go, statefulset.go).
// The above modules declare an interface nameed after the kubernetes resource and an implementation which includes
// the kubernetes api resource and the meta struct.
//
// Namespace does not follow the convention because of the following:
// 1. Declaring a Namespace interface clashes with the already existent kubernetes.Namespace constant in meta.go.
//    This can be solved by naming the interface differently, but there's second issue:
// 2. Defining a Namespace implementation and composing it with *apiv1.Namespace and Meta also produces name clashes
//    between *apiv1.Namespace and meta.Namespace()
func GetNode(ns *apiv1.Namespace) report.Node {
	m := meta{ns.ObjectMeta}
	return m.MetaNode(report.MakeNamespaceNodeID(m.UID()))
}
