# SimilarityViewerUrlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorMetadataOverride** | Pointer to [**AuthorMetadataOverride**](AuthorMetadataOverride.md) |  | [optional] 
**ViewerUserId** | Pointer to **string** | viewer&#39;s user ID | [optional] 
**Locale** | Pointer to **string** | two character locale language setting (e.g. &#39;en&#39; or &#39;de&#39;) or full value | [optional] 
**ViewerDefaultPermissionSet** | Pointer to **string** | Default viewer permission set, accepts INSTRUCTOR, LEARNER, EDITOR, USER, APPLICANT, ADMINISTRATOR, UNDEFINED | [optional] 
**ViewerPermissions** | Pointer to [**ViewerPermissions**](ViewerPermissions.md) |  | [optional] 
**Eula** | Pointer to [**Eula**](Eula.md) |  | [optional] 
**Sidebar** | Pointer to [**Sidebar**](Sidebar.md) |  | [optional] 
**Similarity** | Pointer to [**SimilaritySettings**](SimilaritySettings.md) |  | [optional] 

## Methods

### NewSimilarityViewerUrlSettings

`func NewSimilarityViewerUrlSettings() *SimilarityViewerUrlSettings`

NewSimilarityViewerUrlSettings instantiates a new SimilarityViewerUrlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityViewerUrlSettingsWithDefaults

`func NewSimilarityViewerUrlSettingsWithDefaults() *SimilarityViewerUrlSettings`

NewSimilarityViewerUrlSettingsWithDefaults instantiates a new SimilarityViewerUrlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorMetadataOverride

`func (o *SimilarityViewerUrlSettings) GetAuthorMetadataOverride() AuthorMetadataOverride`

GetAuthorMetadataOverride returns the AuthorMetadataOverride field if non-nil, zero value otherwise.

### GetAuthorMetadataOverrideOk

`func (o *SimilarityViewerUrlSettings) GetAuthorMetadataOverrideOk() (*AuthorMetadataOverride, bool)`

GetAuthorMetadataOverrideOk returns a tuple with the AuthorMetadataOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMetadataOverride

`func (o *SimilarityViewerUrlSettings) SetAuthorMetadataOverride(v AuthorMetadataOverride)`

SetAuthorMetadataOverride sets AuthorMetadataOverride field to given value.

### HasAuthorMetadataOverride

`func (o *SimilarityViewerUrlSettings) HasAuthorMetadataOverride() bool`

HasAuthorMetadataOverride returns a boolean if a field has been set.

### GetViewerUserId

`func (o *SimilarityViewerUrlSettings) GetViewerUserId() string`

GetViewerUserId returns the ViewerUserId field if non-nil, zero value otherwise.

### GetViewerUserIdOk

`func (o *SimilarityViewerUrlSettings) GetViewerUserIdOk() (*string, bool)`

GetViewerUserIdOk returns a tuple with the ViewerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerUserId

`func (o *SimilarityViewerUrlSettings) SetViewerUserId(v string)`

SetViewerUserId sets ViewerUserId field to given value.

### HasViewerUserId

`func (o *SimilarityViewerUrlSettings) HasViewerUserId() bool`

HasViewerUserId returns a boolean if a field has been set.

### GetLocale

`func (o *SimilarityViewerUrlSettings) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SimilarityViewerUrlSettings) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SimilarityViewerUrlSettings) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SimilarityViewerUrlSettings) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetViewerDefaultPermissionSet

`func (o *SimilarityViewerUrlSettings) GetViewerDefaultPermissionSet() string`

GetViewerDefaultPermissionSet returns the ViewerDefaultPermissionSet field if non-nil, zero value otherwise.

### GetViewerDefaultPermissionSetOk

`func (o *SimilarityViewerUrlSettings) GetViewerDefaultPermissionSetOk() (*string, bool)`

GetViewerDefaultPermissionSetOk returns a tuple with the ViewerDefaultPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerDefaultPermissionSet

`func (o *SimilarityViewerUrlSettings) SetViewerDefaultPermissionSet(v string)`

SetViewerDefaultPermissionSet sets ViewerDefaultPermissionSet field to given value.

### HasViewerDefaultPermissionSet

`func (o *SimilarityViewerUrlSettings) HasViewerDefaultPermissionSet() bool`

HasViewerDefaultPermissionSet returns a boolean if a field has been set.

### GetViewerPermissions

`func (o *SimilarityViewerUrlSettings) GetViewerPermissions() ViewerPermissions`

GetViewerPermissions returns the ViewerPermissions field if non-nil, zero value otherwise.

### GetViewerPermissionsOk

`func (o *SimilarityViewerUrlSettings) GetViewerPermissionsOk() (*ViewerPermissions, bool)`

GetViewerPermissionsOk returns a tuple with the ViewerPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerPermissions

`func (o *SimilarityViewerUrlSettings) SetViewerPermissions(v ViewerPermissions)`

SetViewerPermissions sets ViewerPermissions field to given value.

### HasViewerPermissions

`func (o *SimilarityViewerUrlSettings) HasViewerPermissions() bool`

HasViewerPermissions returns a boolean if a field has been set.

### GetEula

`func (o *SimilarityViewerUrlSettings) GetEula() Eula`

GetEula returns the Eula field if non-nil, zero value otherwise.

### GetEulaOk

`func (o *SimilarityViewerUrlSettings) GetEulaOk() (*Eula, bool)`

GetEulaOk returns a tuple with the Eula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEula

`func (o *SimilarityViewerUrlSettings) SetEula(v Eula)`

SetEula sets Eula field to given value.

### HasEula

`func (o *SimilarityViewerUrlSettings) HasEula() bool`

HasEula returns a boolean if a field has been set.

### GetSidebar

`func (o *SimilarityViewerUrlSettings) GetSidebar() Sidebar`

GetSidebar returns the Sidebar field if non-nil, zero value otherwise.

### GetSidebarOk

`func (o *SimilarityViewerUrlSettings) GetSidebarOk() (*Sidebar, bool)`

GetSidebarOk returns a tuple with the Sidebar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidebar

`func (o *SimilarityViewerUrlSettings) SetSidebar(v Sidebar)`

SetSidebar sets Sidebar field to given value.

### HasSidebar

`func (o *SimilarityViewerUrlSettings) HasSidebar() bool`

HasSidebar returns a boolean if a field has been set.

### GetSimilarity

`func (o *SimilarityViewerUrlSettings) GetSimilarity() SimilaritySettings`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *SimilarityViewerUrlSettings) GetSimilarityOk() (*SimilaritySettings, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *SimilarityViewerUrlSettings) SetSimilarity(v SimilaritySettings)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *SimilarityViewerUrlSettings) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


