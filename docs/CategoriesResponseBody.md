# CategoriesResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryResponse**](CategoryResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewCategoriesResponseBody

`func NewCategoriesResponseBody() *CategoriesResponseBody`

NewCategoriesResponseBody instantiates a new CategoriesResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesResponseBodyWithDefaults

`func NewCategoriesResponseBodyWithDefaults() *CategoriesResponseBody`

NewCategoriesResponseBodyWithDefaults instantiates a new CategoriesResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *CategoriesResponseBody) GetCategories() []CategoryResponse`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CategoriesResponseBody) GetCategoriesOk() (*[]CategoryResponse, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CategoriesResponseBody) SetCategories(v []CategoryResponse)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CategoriesResponseBody) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPagination

`func (o *CategoriesResponseBody) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CategoriesResponseBody) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CategoriesResponseBody) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CategoriesResponseBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


