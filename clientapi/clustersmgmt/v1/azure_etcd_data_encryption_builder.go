/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

// AzureEtcdDataEncryptionBuilder contains the data and logic needed to build 'azure_etcd_data_encryption' objects.
//
// Contains the necessary attributes to support data encryption for Azure based clusters.
type AzureEtcdDataEncryptionBuilder struct {
	bitmap_           uint32
	customerManaged   *AzureEtcdDataEncryptionCustomerManagedBuilder
	keyManagementMode string
}

// NewAzureEtcdDataEncryption creates a new builder of 'azure_etcd_data_encryption' objects.
func NewAzureEtcdDataEncryption() *AzureEtcdDataEncryptionBuilder {
	return &AzureEtcdDataEncryptionBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureEtcdDataEncryptionBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// CustomerManaged sets the value of the 'customer_managed' attribute to the given value.
//
// Contains the necessary attributes to support etcd data encryption with customer managed keys
// for Azure based clusters.
func (b *AzureEtcdDataEncryptionBuilder) CustomerManaged(value *AzureEtcdDataEncryptionCustomerManagedBuilder) *AzureEtcdDataEncryptionBuilder {
	b.customerManaged = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// KeyManagementMode sets the value of the 'key_management_mode' attribute to the given value.
func (b *AzureEtcdDataEncryptionBuilder) KeyManagementMode(value string) *AzureEtcdDataEncryptionBuilder {
	b.keyManagementMode = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureEtcdDataEncryptionBuilder) Copy(object *AzureEtcdDataEncryption) *AzureEtcdDataEncryptionBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.customerManaged != nil {
		b.customerManaged = NewAzureEtcdDataEncryptionCustomerManaged().Copy(object.customerManaged)
	} else {
		b.customerManaged = nil
	}
	b.keyManagementMode = object.keyManagementMode
	return b
}

// Build creates a 'azure_etcd_data_encryption' object using the configuration stored in the builder.
func (b *AzureEtcdDataEncryptionBuilder) Build() (object *AzureEtcdDataEncryption, err error) {
	object = new(AzureEtcdDataEncryption)
	object.bitmap_ = b.bitmap_
	if b.customerManaged != nil {
		object.customerManaged, err = b.customerManaged.Build()
		if err != nil {
			return
		}
	}
	object.keyManagementMode = b.keyManagementMode
	return
}
