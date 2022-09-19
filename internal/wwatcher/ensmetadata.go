// Package ensmetadata provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package wwatcher

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for NetworkName.
const (
	Goerli  NetworkName = "goerli"
	Mainnet NetworkName = "mainnet"
	Rinkeby NetworkName = "rinkeby"
	Ropsten NetworkName = "ropsten"
)

// ENSMetadata defines model for ENSMetadata.
type ENSMetadata struct {
	Attributes   []struct {
		TraitType   string `json:"trait_type"`
		DisplayType string `json:"display_type"`
		Value       interface{}  `json:"value"`
	} `json:"attributes"`
	BackgroundImage string `json:"background_image"`
	Description     string `json:"description"`
	ImageUrl        string `json:"image_url"`
	Name            string `json:"name"`
	NameLength      int    `json:"name_length"`
	SegmentLength   int    `json:"segment_length"`
	Url             string `json:"url"`
	Version         int    `json:"version"`
}

// ContractAddress defines model for contractAddress.
type ContractAddress = string

// EnsName defines model for ensName.
type EnsName = string

// Name of the chain to query for.
type NetworkName string

// TokenId defines model for tokenId.
type TokenId = string

// GetQueryNFTParams defines parameters for GetQueryNFT.
type GetQueryNFTParams struct {
	// NFT URI as defined under CAIP-22 for erc721 assets and CAIP-29 for erc1155 assets.
	Uri *string `form:"uri,omitempty" json:"uri,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetQueryNFT request
	GetQueryNFT(ctx context.Context, params *GetQueryNFTParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetworkNameAvatarName request
	GetNetworkNameAvatarName(ctx context.Context, networkName NetworkName, name EnsName, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetworkNameAvatarNameMeta request
	GetNetworkNameAvatarNameMeta(ctx context.Context, networkName NetworkName, name EnsName, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetworkNameContractAddressTokenId request
	GetNetworkNameContractAddressTokenId(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetworkNameContractAddressTokenIdImage request
	GetNetworkNameContractAddressTokenIdImage(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNetworkNameContractAddressTokenIdRasterize request
	GetNetworkNameContractAddressTokenIdRasterize(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetQueryNFT(ctx context.Context, params *GetQueryNFTParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetQueryNFTRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetworkNameAvatarName(ctx context.Context, networkName NetworkName, name EnsName, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkNameAvatarNameRequest(c.Server, networkName, name)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetworkNameAvatarNameMeta(ctx context.Context, networkName NetworkName, name EnsName, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkNameAvatarNameMetaRequest(c.Server, networkName, name)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetworkNameContractAddressTokenId(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkNameContractAddressTokenIdRequest(c.Server, networkName, contractAddress, tokenId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetworkNameContractAddressTokenIdImage(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkNameContractAddressTokenIdImageRequest(c.Server, networkName, contractAddress, tokenId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNetworkNameContractAddressTokenIdRasterize(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNetworkNameContractAddressTokenIdRasterizeRequest(c.Server, networkName, contractAddress, tokenId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetQueryNFTRequest generates requests for GetQueryNFT
func NewGetQueryNFTRequest(server string, params *GetQueryNFTParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/queryNFT")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Uri != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "uri", runtime.ParamLocationQuery, *params.Uri); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkNameAvatarNameRequest generates requests for GetNetworkNameAvatarName
func NewGetNetworkNameAvatarNameRequest(server string, networkName NetworkName, name EnsName) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkName", runtime.ParamLocationPath, networkName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/avatar/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkNameAvatarNameMetaRequest generates requests for GetNetworkNameAvatarNameMeta
func NewGetNetworkNameAvatarNameMetaRequest(server string, networkName NetworkName, name EnsName) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkName", runtime.ParamLocationPath, networkName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/avatar/%s/meta", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkNameContractAddressTokenIdRequest generates requests for GetNetworkNameContractAddressTokenId
func NewGetNetworkNameContractAddressTokenIdRequest(server string, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkName", runtime.ParamLocationPath, networkName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "contractAddress", runtime.ParamLocationPath, contractAddress)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "tokenId", runtime.ParamLocationPath, tokenId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/%s/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkNameContractAddressTokenIdImageRequest generates requests for GetNetworkNameContractAddressTokenIdImage
func NewGetNetworkNameContractAddressTokenIdImageRequest(server string, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkName", runtime.ParamLocationPath, networkName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "contractAddress", runtime.ParamLocationPath, contractAddress)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "tokenId", runtime.ParamLocationPath, tokenId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/%s/%s/image", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNetworkNameContractAddressTokenIdRasterizeRequest generates requests for GetNetworkNameContractAddressTokenIdRasterize
func NewGetNetworkNameContractAddressTokenIdRasterizeRequest(server string, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "networkName", runtime.ParamLocationPath, networkName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "contractAddress", runtime.ParamLocationPath, contractAddress)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "tokenId", runtime.ParamLocationPath, tokenId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/%s/%s/%s/rasterize", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetQueryNFT request
	GetQueryNFTWithResponse(ctx context.Context, params *GetQueryNFTParams, reqEditors ...RequestEditorFn) (*GetQueryNFTResponse, error)

	// GetNetworkNameContractAddressTokenId request
	GetNetworkNameContractAddressTokenIdWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdResponse, error)

	// GetNetworkNameContractAddressTokenIdImage request
	GetNetworkNameContractAddressTokenIdImageWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdImageResponse, error)

	// GetNetworkNameContractAddressTokenIdRasterize request
	GetNetworkNameContractAddressTokenIdRasterizeWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdRasterizeResponse, error)
}

type GetQueryNFTResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetQueryNFTResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetQueryNFTResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetNetworkNameAvatarNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetNetworkNameAvatarNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworkNameAvatarNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}


type GetNetworkNameContractAddressTokenIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ENSMetadata
	XML200       *ENSMetadata
}

// Status returns HTTPResponse.Status
func (r GetNetworkNameContractAddressTokenIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworkNameContractAddressTokenIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetNetworkNameContractAddressTokenIdImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetNetworkNameContractAddressTokenIdImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworkNameContractAddressTokenIdImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetNetworkNameContractAddressTokenIdRasterizeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetNetworkNameContractAddressTokenIdRasterizeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNetworkNameContractAddressTokenIdRasterizeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetQueryNFTWithResponse request returning *GetQueryNFTResponse
func (c *ClientWithResponses) GetQueryNFTWithResponse(ctx context.Context, params *GetQueryNFTParams, reqEditors ...RequestEditorFn) (*GetQueryNFTResponse, error) {
	rsp, err := c.GetQueryNFT(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetQueryNFTResponse(rsp)
}

// GetNetworkNameAvatarNameWithResponse request returning *GetNetworkNameAvatarNameResponse
func (c *ClientWithResponses) GetNetworkNameAvatarNameWithResponse(ctx context.Context, networkName NetworkName, name EnsName, reqEditors ...RequestEditorFn) (*GetNetworkNameAvatarNameResponse, error) {
	rsp, err := c.GetNetworkNameAvatarName(ctx, networkName, name, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworkNameAvatarNameResponse(rsp)
}

// GetNetworkNameContractAddressTokenIdWithResponse request returning *GetNetworkNameContractAddressTokenIdResponse
func (c *ClientWithResponses) GetNetworkNameContractAddressTokenIdWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdResponse, error) {
	rsp, err := c.GetNetworkNameContractAddressTokenId(ctx, networkName, contractAddress, tokenId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworkNameContractAddressTokenIdResponse(rsp)
}

// GetNetworkNameContractAddressTokenIdImageWithResponse request returning *GetNetworkNameContractAddressTokenIdImageResponse
func (c *ClientWithResponses) GetNetworkNameContractAddressTokenIdImageWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdImageResponse, error) {
	rsp, err := c.GetNetworkNameContractAddressTokenIdImage(ctx, networkName, contractAddress, tokenId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworkNameContractAddressTokenIdImageResponse(rsp)
}

// GetNetworkNameContractAddressTokenIdRasterizeWithResponse request returning *GetNetworkNameContractAddressTokenIdRasterizeResponse
func (c *ClientWithResponses) GetNetworkNameContractAddressTokenIdRasterizeWithResponse(ctx context.Context, networkName NetworkName, contractAddress ContractAddress, tokenId TokenId, reqEditors ...RequestEditorFn) (*GetNetworkNameContractAddressTokenIdRasterizeResponse, error) {
	rsp, err := c.GetNetworkNameContractAddressTokenIdRasterize(ctx, networkName, contractAddress, tokenId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNetworkNameContractAddressTokenIdRasterizeResponse(rsp)
}

// ParseGetQueryNFTResponse parses an HTTP response from a GetQueryNFTWithResponse call
func ParseGetQueryNFTResponse(rsp *http.Response) (*GetQueryNFTResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetQueryNFTResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetNetworkNameAvatarNameResponse parses an HTTP response from a GetNetworkNameAvatarNameWithResponse call
func ParseGetNetworkNameAvatarNameResponse(rsp *http.Response) (*GetNetworkNameAvatarNameResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNetworkNameAvatarNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetNetworkNameContractAddressTokenIdResponse parses an HTTP response from a GetNetworkNameContractAddressTokenIdWithResponse call
func ParseGetNetworkNameContractAddressTokenIdResponse(rsp *http.Response) (*GetNetworkNameContractAddressTokenIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNetworkNameContractAddressTokenIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ENSMetadata
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "xml") && rsp.StatusCode == 200:
		var dest ENSMetadata
		if err := xml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.XML200 = &dest

	}

	return response, nil
}

// ParseGetNetworkNameContractAddressTokenIdImageResponse parses an HTTP response from a GetNetworkNameContractAddressTokenIdImageWithResponse call
func ParseGetNetworkNameContractAddressTokenIdImageResponse(rsp *http.Response) (*GetNetworkNameContractAddressTokenIdImageResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNetworkNameContractAddressTokenIdImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetNetworkNameContractAddressTokenIdRasterizeResponse parses an HTTP response from a GetNetworkNameContractAddressTokenIdRasterizeWithResponse call
func ParseGetNetworkNameContractAddressTokenIdRasterizeResponse(rsp *http.Response) (*GetNetworkNameContractAddressTokenIdRasterizeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNetworkNameContractAddressTokenIdRasterizeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY227bRhN+lQXzX/wBLFKUpVrWVRXFLgTEShMrF0USCCtyJG5M7jK7S8euoHcvZnkQ",
	"TzrEbdoCta+sXc75+2aG3FieiGLBgWtljTaW8gKIqPl3fE81lTegqU81xZNYihikZmDuqdaSLROd/oIH",
	"GsUhWCNrXJwTsSKz67l1ZunHGK+Uloyvre2ZtaTe3VqKhPsLT4RCVlW8Km6JuT2gyAflSRZrJnhVx+vd",
	"xQHxQCi9iKAtPi+gjC+YX1U7wVMyfU2+BSAB1RIJivmg2tR7gmtJPb2gvi9B1TI1yW5JdnvAT04jUDH1",
	"oKph/hhDJkVyW23iElYggXuwSGRYVXFD5R3oOKQekA/v3xzwQYs7aCZkjqeYkH2C2+JELL+Ap1EVi+i6",
	"FsoUjwhCzTnshpFd5JhsU3BUuJGE042rhfjGoYbXt3ikAhaTe5BsxTx6BHZYzqqKGY3ggEAiWfX5q9kt",
	"4ZkMNUwlgQh9kG3SjyLRybIl7t/Si4OFN/D5mjAJvjX6aDzJAqiy76zcEfISf64X/8x6iIwTaQrqXQat",
	"Xc1uv7/rTBKlRUS0pEwrQpci0eRqdnuk97Tg8K1ka8bzjGSpTZ/7nvZzGwhpPEirVE3UqbDMnEAtyPC9",
	"XjThlBve9/QiBL7WQaO3YQMBSdJbtA1c7VXU8BetjuM4T54OYK/sPUjVSFoeaH55DIvHUViONXV4Z7oF",
	"CeVKHIVuGafbUrMft/X67sPgYuUOhxd0+OravVy6/dWk61//dO1fvupRz7scr9z+BYyHg0bUdbt1O2gb",
	"uJo1QMCZd2eDifywxlwaNXHQ34S8y7VVAF50KaysmY9EC/I1AflIVkLaFjqSRFiaiDLOAZMmGb+D5SP+",
	"J2KlARO/FiBDVkrxHsfKzmzzETStTaB+r+dedoeD3mAwcN3z80Fv0B0ML/v9Xq87PO+6g3Pbto/mIFe9",
	"3aIhxlfCLAGCa5ypaDCiLMzSTz39M3Bl+wLjVI1OYN2CTtnjx4JxrXZ5QoxHGWwI5T6JhETghcwDrqDk",
	"0c10Tt5kpxg70wVJctyRW5D3zIMSqkdW1+7aboeGcUBtF+PshGJtoqGhnsODRiUa15ckIqaiOy07RkzS",
	"tcx6cW3+MvaMLIcqBVo5qNNW9ykpRQycxswaWed21+5aZ1ZMdWAo4JiwcbKMNtYadBNU70xe8lQhkEwT",
	"+PB+igzG5m/GKZbd+gX0u1wfGpE0Ag1SWaOPDaymOghVxIcV4+CThPsgyWQ8/bXT6xlDIL2LnkvSmEw5",
	"0tvL/NZ1B4PsGjHEULOJKB+Co2wippszRlfvWZ+xaalYcJUOrV6328Ks63kBC6z2oO2hKdcgOQ1NwUCS",
	"KymFzBDrbEpk2Trp4HI26ON2b+oRS7UR18j3bKc2HdaztOvWkm8yg1XfJabM3nLj1jKBcsL+J2GFSHN2",
	"7yJO9iLi1DpAWwDZFGhz4M9Y3jXFkwqYrpArFhqy9rv9lhoLfFdIQq3ICimW1tltPviBqySOhdTgkywB",
	"J1TZyd9jjpW6wNmp1cZ281zxdBoAN8mlcRxme77zRaWLzGlWWjbesrJsKj1RV2MQFZNi9/L1N2FzU9tT",
	"ts4mm7GH+1GlEx5G6KRqYZ6N8H8MqS126svaU221LH315L2hSwgDqoL/37sviYNupb96L3EXeRSJLF5I",
	"7E/8E78REkYk0DpWI8fxhafs0lJTmOzQmHWKLwgOindiKTxQivH1C7TB+LpjPlG0s1IXhXla8Lvd7EcS",
	"s77PP5mVVUV/MSVPXwt+IIGd4sX5II1P2SnaOTzNJJ+J/N8i8r9gpzoEe0mVBsl+PxH6JH+eZt8mvp8I",
	"7wuLz2R4JsOODD9wEGy3fwQAAP//tUcgaBYaAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
