// Code generated by ""fitask" -type=RouteTable"; DO NOT EDIT

package awstasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// RouteTable

// JSON marshalling boilerplate
type realRouteTable RouteTable

func (o *RouteTable) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realRouteTable
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = RouteTable(r)
	return nil
}

var _ fi.HasName = &RouteTable{}

func (e *RouteTable) GetName() *string {
	return e.Name
}

func (e *RouteTable) SetName(name string) {
	e.Name = &name
}

func (e *RouteTable) String() string {
	return fi.TaskAsString(e)
}
