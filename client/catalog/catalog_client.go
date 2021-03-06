// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BatchDeleteCatalogObjects(params *BatchDeleteCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchDeleteCatalogObjectsOK, error)

	BatchRetrieveCatalogObjects(params *BatchRetrieveCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchRetrieveCatalogObjectsOK, error)

	BatchUpsertCatalogObjects(params *BatchUpsertCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchUpsertCatalogObjectsOK, error)

	CatalogInfo(params *CatalogInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogInfoOK, error)

	DeleteCatalogObject(params *DeleteCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCatalogObjectOK, error)

	ListCatalog(params *ListCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCatalogOK, error)

	RetrieveCatalogObject(params *RetrieveCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveCatalogObjectOK, error)

	SearchCatalogItems(params *SearchCatalogItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCatalogItemsOK, error)

	SearchCatalogObjects(params *SearchCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCatalogObjectsOK, error)

	UpdateItemModifierLists(params *UpdateItemModifierListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateItemModifierListsOK, error)

	UpdateItemTaxes(params *UpdateItemTaxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateItemTaxesOK, error)

	UpsertCatalogObject(params *UpsertCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertCatalogObjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BatchDeleteCatalogObjects batches delete catalog objects

  Deletes a set of [CatalogItem](#type-catalogitem)s based on the
provided list of target IDs and returns a set of successfully deleted IDs in
the response. Deletion is a cascading event such that all children of the
targeted object are also deleted. For example, deleting a CatalogItem will
also delete all of its [CatalogItemVariation](#type-catalogitemvariation)
children.

`BatchDeleteCatalogObjects` succeeds even if only a portion of the targeted
IDs can be deleted. The response will only include IDs that were
actually deleted.
*/
func (a *Client) BatchDeleteCatalogObjects(params *BatchDeleteCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchDeleteCatalogObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchDeleteCatalogObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BatchDeleteCatalogObjects",
		Method:             "POST",
		PathPattern:        "/v2/catalog/batch-delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchDeleteCatalogObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchDeleteCatalogObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchDeleteCatalogObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchRetrieveCatalogObjects batches retrieve catalog objects

  Returns a set of objects based on the provided ID.
Each [CatalogItem](#type-catalogitem) returned in the set includes all of its
child information including: all of its
[CatalogItemVariation](#type-catalogitemvariation) objects, references to
its [CatalogModifierList](#type-catalogmodifierlist) objects, and the ids of
any [CatalogTax](#type-catalogtax) objects that apply to it.
*/
func (a *Client) BatchRetrieveCatalogObjects(params *BatchRetrieveCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchRetrieveCatalogObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchRetrieveCatalogObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BatchRetrieveCatalogObjects",
		Method:             "POST",
		PathPattern:        "/v2/catalog/batch-retrieve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchRetrieveCatalogObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchRetrieveCatalogObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchRetrieveCatalogObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchUpsertCatalogObjects batches upsert catalog objects

  Creates or updates up to 10,000 target objects based on the provided
list of objects. The target objects are grouped into batches and each batch is
inserted/updated in an all-or-nothing manner. If an object within a batch is
malformed in some way, or violates a database constraint, the entire batch
containing that item will be disregarded. However, other batches in the same
request may still succeed. Each batch may contain up to 1,000 objects, and
batches will be processed in order as long as the total object count for the
request (items, variations, modifier lists, discounts, and taxes) is no more
than 10,000.
*/
func (a *Client) BatchUpsertCatalogObjects(params *BatchUpsertCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchUpsertCatalogObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchUpsertCatalogObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BatchUpsertCatalogObjects",
		Method:             "POST",
		PathPattern:        "/v2/catalog/batch-upsert",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchUpsertCatalogObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchUpsertCatalogObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchUpsertCatalogObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CatalogInfo catalogs info

  Retrieves information about the Square Catalog API, such as batch size
limits that can be used by the `BatchUpsertCatalogObjects` endpoint.
*/
func (a *Client) CatalogInfo(params *CatalogInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CatalogInfo",
		Method:             "GET",
		PathPattern:        "/v2/catalog/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CatalogInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CatalogInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCatalogObject deletes catalog object

  Deletes a single [CatalogObject](#type-catalogobject) based on the
provided ID and returns the set of successfully deleted IDs in the response.
Deletion is a cascading event such that all children of the targeted object
are also deleted. For example, deleting a [CatalogItem](#type-catalogitem)
will also delete all of its
[CatalogItemVariation](#type-catalogitemvariation) children.
*/
func (a *Client) DeleteCatalogObject(params *DeleteCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCatalogObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCatalogObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCatalogObject",
		Method:             "DELETE",
		PathPattern:        "/v2/catalog/object/{object_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCatalogObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCatalogObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCatalogObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCatalog lists catalog

  Returns a list of [CatalogObject](#type-catalogobject)s that includes
all objects of a set of desired types (for example, all [CatalogItem](#type-catalogitem)
and [CatalogTax](#type-catalogtax) objects) in the catalog. The `types` parameter
is specified as a comma-separated list of valid [CatalogObject](#type-catalogobject) types:
`ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, `CATEGORY`, `DISCOUNT`, `TAX`, `IMAGE`.

__Important:__ ListCatalog does not return deleted catalog items. To retrieve
deleted catalog items, use [SearchCatalogObjects](#endpoint-Catalog-SearchCatalogObjects)
and set the `include_deleted_objects` attribute value to `true`.
*/
func (a *Client) ListCatalog(params *ListCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListCatalog",
		Method:             "GET",
		PathPattern:        "/v2/catalog/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveCatalogObject retrieves catalog object

  Returns a single [CatalogItem](#type-catalogitem) as a
[CatalogObject](#type-catalogobject) based on the provided ID. The returned
object includes all of the relevant [CatalogItem](#type-catalogitem)
information including: [CatalogItemVariation](#type-catalogitemvariation)
children, references to its
[CatalogModifierList](#type-catalogmodifierlist) objects, and the ids of
any [CatalogTax](#type-catalogtax) objects that apply to it.
*/
func (a *Client) RetrieveCatalogObject(params *RetrieveCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveCatalogObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveCatalogObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveCatalogObject",
		Method:             "GET",
		PathPattern:        "/v2/catalog/object/{object_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveCatalogObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveCatalogObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveCatalogObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchCatalogItems searches catalog items

  Searches for catalog items or item variations by matching supported search attribute values, including
custom attribute values, against one or more of the specified query expressions.

This (`SearchCatalogItems`) endpoint differs from the [SearchCatalogObjects](#endpoint-Catalog-SearchCatalogObjects)
endpoint in the following aspects:

- `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
- `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
- `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
- The both endpoints use different call conventions, including the query filter formats.
*/
func (a *Client) SearchCatalogItems(params *SearchCatalogItemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCatalogItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchCatalogItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchCatalogItems",
		Method:             "POST",
		PathPattern:        "/v2/catalog/search-catalog-items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchCatalogItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchCatalogItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchCatalogItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchCatalogObjects searches catalog objects

  Searches for [CatalogObject](#type-CatalogObject) of any type by matching supported search attribute values,
excluding custom attribute values on items or item variations, against one or more of the specified query expressions.

This (`SearchCatalogObjects`) endpoint differs from the [SearchCatalogItems](#endpoint-Catalog-SearchCatalogItems)
endpoint in the following aspects:

- `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects.
- `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not.
- `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does.
- The both endpoints have different call conventions, including the query filter formats.
*/
func (a *Client) SearchCatalogObjects(params *SearchCatalogObjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchCatalogObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchCatalogObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchCatalogObjects",
		Method:             "POST",
		PathPattern:        "/v2/catalog/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchCatalogObjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchCatalogObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchCatalogObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateItemModifierLists updates item modifier lists

  Updates the [CatalogModifierList](#type-catalogmodifierlist) objects
that apply to the targeted [CatalogItem](#type-catalogitem) without having
to perform an upsert on the entire item.
*/
func (a *Client) UpdateItemModifierLists(params *UpdateItemModifierListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateItemModifierListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateItemModifierListsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateItemModifierLists",
		Method:             "POST",
		PathPattern:        "/v2/catalog/update-item-modifier-lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateItemModifierListsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateItemModifierListsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateItemModifierLists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateItemTaxes updates item taxes

  Updates the [CatalogTax](#type-catalogtax) objects that apply to the
targeted [CatalogItem](#type-catalogitem) without having to perform an
upsert on the entire item.
*/
func (a *Client) UpdateItemTaxes(params *UpdateItemTaxesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateItemTaxesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateItemTaxesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateItemTaxes",
		Method:             "POST",
		PathPattern:        "/v2/catalog/update-item-taxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateItemTaxesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateItemTaxesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateItemTaxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpsertCatalogObject upserts catalog object

  Creates or updates the target [CatalogObject](#type-catalogobject).
*/
func (a *Client) UpsertCatalogObject(params *UpsertCatalogObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpsertCatalogObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertCatalogObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpsertCatalogObject",
		Method:             "POST",
		PathPattern:        "/v2/catalog/object",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertCatalogObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpsertCatalogObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpsertCatalogObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
