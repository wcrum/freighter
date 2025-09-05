package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"freighter.dev/go/freighter/pkg/consts"
)

var (
	ContentGroupVersion    = schema.GroupVersion{Group: consts.ContentGroup, Version: "v1"}
	CollectionGroupVersion = schema.GroupVersion{Group: consts.CollectionGroup, Version: "v1"}
)
