// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type CreateSaml2IdentityProviderDetails struct {

	// The OCID of your tenancy.
	CompartmentID string `json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	Name string `json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// The identity provider service or product (e.g., Oracle Identity Cloud Service).
	// Example: `IDCS`
	ProductType string `json:"productType,omitempty"`

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol string `json:"protocol,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataURL string `json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata string `json:"metadata,omitempty"`
}
