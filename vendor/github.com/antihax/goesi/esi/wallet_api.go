/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"

	"encoding/json"
	"fmt"

	"github.com/mailru/easyjson"
)

// Linger please
var (
	_ context.Context
)

type WalletApiService service

/* WalletApiService Get a character&#39;s wallet balance
Returns a character&#39;s wallet balance  ---  This route is cached for up to 120 seconds

* @param ctx context.Context Authentication Context
@param characterId An EVE character ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return float32*/
func (a *WalletApiService) GetCharactersCharacterIdWallet(ctx context.Context, characterId int32, localVarOptionals map[string]interface{}) (float32, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     float32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/wallet/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WalletApiService Get character wallet journal
Retrieve character wallet journal  ---  This route is cached for up to 3600 seconds

* @param ctx context.Context Authentication Context
@param characterId An EVE character ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "fromId" (int64) Only show journal entries happened before the transaction referenced by this id
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetCharactersCharacterIdWalletJournal200Ok*/
func (a *WalletApiService) GetCharactersCharacterIdWalletJournal(ctx context.Context, characterId int32, localVarOptionals map[string]interface{}) ([]GetCharactersCharacterIdWalletJournal200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCharactersCharacterIdWalletJournal200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v2/characters/{character_id}/wallet/journal/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fromId"], "int64", "fromId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fromId"].(int64); localVarOk {
		localVarQueryParams.Add("from_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WalletApiService Get wallet transactions
Get wallet transactions of a character  ---  This route is cached for up to 3600 seconds

* @param ctx context.Context Authentication Context
@param characterId An EVE character ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "fromId" (int64) Only show transactions happened before the one referenced by this id
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetCharactersCharacterIdWalletTransactions200Ok*/
func (a *WalletApiService) GetCharactersCharacterIdWalletTransactions(ctx context.Context, characterId int32, localVarOptionals map[string]interface{}) ([]GetCharactersCharacterIdWalletTransactions200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCharactersCharacterIdWalletTransactions200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/wallet/transactions/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fromId"], "int64", "fromId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fromId"].(int64); localVarOk {
		localVarQueryParams.Add("from_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WalletApiService Returns a corporation&#39;s wallet balance
Get a corporation&#39;s wallets  ---  This route is cached for up to 300 seconds

* @param ctx context.Context Authentication Context
@param corporationId An EVE corporation ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetCorporationsCorporationIdWallets200Ok*/
func (a *WalletApiService) GetCorporationsCorporationIdWallets(ctx context.Context, corporationId int32, localVarOptionals map[string]interface{}) ([]GetCorporationsCorporationIdWallets200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCorporationsCorporationIdWallets200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/corporations/{corporation_id}/wallets/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WalletApiService Get corporation wallet journal
Retrieve corporation wallet journal  ---  This route is cached for up to 3600 seconds

* @param ctx context.Context Authentication Context
@param corporationId An EVE corporation ID
@param division Wallet key of the division to fetch journals from
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "fromId" (int64) Only show journal entries happened before the transaction referenced by this id
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetCorporationsCorporationIdWalletsDivisionJournal200Ok*/
func (a *WalletApiService) GetCorporationsCorporationIdWalletsDivisionJournal(ctx context.Context, corporationId int32, division int32, localVarOptionals map[string]interface{}) ([]GetCorporationsCorporationIdWalletsDivisionJournal200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCorporationsCorporationIdWalletsDivisionJournal200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/corporations/{corporation_id}/wallets/{division}/journal/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"division"+"}", fmt.Sprintf("%v", division), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if division < 1 {
		return successPayload, nil, reportError("division must be greater than 1")
	}
	if division > 7 {
		return successPayload, nil, reportError("division must be less than 7")
	}
	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fromId"], "int64", "fromId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fromId"].(int64); localVarOk {
		localVarQueryParams.Add("from_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WalletApiService Get corporation wallet transactions
Get wallet transactions of a corporation  ---  This route is cached for up to 3600 seconds

* @param ctx context.Context Authentication Context
@param corporationId An EVE corporation ID
@param division Wallet key of the division to fetch journals from
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "fromId" (int64) Only show journal entries happened before the transaction referenced by this id
    @param "token" (string) Access token to use if unable to set a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetCorporationsCorporationIdWalletsDivisionTransactions200Ok*/
func (a *WalletApiService) GetCorporationsCorporationIdWalletsDivisionTransactions(ctx context.Context, corporationId int32, division int32, localVarOptionals map[string]interface{}) ([]GetCorporationsCorporationIdWalletsDivisionTransactions200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCorporationsCorporationIdWalletsDivisionTransactions200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/corporations/{corporation_id}/wallets/{division}/transactions/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"division"+"}", fmt.Sprintf("%v", division), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if division < 1 {
		return successPayload, nil, reportError("division must be greater than 1")
	}
	if division > 7 {
		return successPayload, nil, reportError("division must be less than 7")
	}
	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fromId"], "int64", "fromId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fromId"].(int64); localVarOk {
		localVarQueryParams.Add("from_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}
	if err = easyjson.UnmarshalFromReader(localVarHttpResponse.Body, &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
