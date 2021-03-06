// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generators

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gardener/gardener/test/integration/framework/networkpolicies"
)

func defaultRegistry() map[string]networkpolicies.CloudAwarePodInfo {
	aws := networkpolicies.AWSNetworkPolicy{}
	azure := networkpolicies.AzureNetworkPolicy{}
	gcp := networkpolicies.GCPNetworkPolicy{}
	openstack := networkpolicies.OpenStackNetworkPolicy{}
	alicloud := networkpolicies.AlicloudNetworkPolicy{}
	return map[string]networkpolicies.CloudAwarePodInfo{
		toKey(aws):       &aws,
		toKey(azure):     &azure,
		toKey(gcp):       &gcp,
		toKey(openstack): &openstack,
		toKey(alicloud):  &alicloud,
	}
}

func toKey(in interface{}) string {
	typeof := reflect.TypeOf(in)
	if typeof.Kind() != reflect.Struct {
		panic("should pass only structs")
	}
	return fmt.Sprintf("%s.%s", typeof.PkgPath(), typeof.Name())
}

// Poor-man's pretty struct printer
func prettyPrint(i interface{}) string {
	s1 := strings.ReplaceAll(fmt.Sprintf("%#v", i), ", ", ",\n")
	s2 := strings.ReplaceAll(s1, "{", "{\n")
	s3 := strings.ReplaceAll(s2, "} ", "}\n")
	s4 := strings.ReplaceAll(s3, "[", "[\n")
	s5 := strings.ReplaceAll(s4, "] ", "]\n")

	return s5
}
