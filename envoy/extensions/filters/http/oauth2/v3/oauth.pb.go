// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/extensions/filters/http/oauth2/v3/oauth.proto

package oauth2v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/Nordix/go-control-plane/envoy/config/core/v3"
	v33 "github.com/Nordix/go-control-plane/envoy/config/route/v3"
	v3 "github.com/Nordix/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v32 "github.com/Nordix/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OAuth2Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The client_id to be used in the authorize calls. This value will be URL encoded when sent to the OAuth server.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The secret used to retrieve the access token. This value will be URL encoded when sent to the OAuth server.
	TokenSecret *v3.SdsSecretConfig `protobuf:"bytes,2,opt,name=token_secret,json=tokenSecret,proto3" json:"token_secret,omitempty"`
	// Configures how the secret token should be created.
	//
	// Types that are assignable to TokenFormation:
	//	*OAuth2Credentials_HmacSecret
	TokenFormation isOAuth2Credentials_TokenFormation `protobuf_oneof:"token_formation"`
	// The cookie names used in OAuth filters flow.
	CookieNames *OAuth2Credentials_CookieNames `protobuf:"bytes,4,opt,name=cookie_names,json=cookieNames,proto3" json:"cookie_names,omitempty"`
}

func (x *OAuth2Credentials) Reset() {
	*x = OAuth2Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Credentials) ProtoMessage() {}

func (x *OAuth2Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Credentials.ProtoReflect.Descriptor instead.
func (*OAuth2Credentials) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *OAuth2Credentials) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuth2Credentials) GetTokenSecret() *v3.SdsSecretConfig {
	if x != nil {
		return x.TokenSecret
	}
	return nil
}

func (m *OAuth2Credentials) GetTokenFormation() isOAuth2Credentials_TokenFormation {
	if m != nil {
		return m.TokenFormation
	}
	return nil
}

func (x *OAuth2Credentials) GetHmacSecret() *v3.SdsSecretConfig {
	if x, ok := x.GetTokenFormation().(*OAuth2Credentials_HmacSecret); ok {
		return x.HmacSecret
	}
	return nil
}

func (x *OAuth2Credentials) GetCookieNames() *OAuth2Credentials_CookieNames {
	if x != nil {
		return x.CookieNames
	}
	return nil
}

type isOAuth2Credentials_TokenFormation interface {
	isOAuth2Credentials_TokenFormation()
}

type OAuth2Credentials_HmacSecret struct {
	// If present, the secret token will be a HMAC using the provided secret.
	HmacSecret *v3.SdsSecretConfig `protobuf:"bytes,3,opt,name=hmac_secret,json=hmacSecret,proto3,oneof"`
}

func (*OAuth2Credentials_HmacSecret) isOAuth2Credentials_TokenFormation() {}

// OAuth config
//
// [#next-free-field: 11]
type OAuth2Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endpoint on the authorization server to retrieve the access token from.
	TokenEndpoint *v31.HttpUri `protobuf:"bytes,1,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	// The endpoint redirect to for authorization in response to unauthorized requests.
	AuthorizationEndpoint string `protobuf:"bytes,2,opt,name=authorization_endpoint,json=authorizationEndpoint,proto3" json:"authorization_endpoint,omitempty"`
	// Credentials used for OAuth.
	Credentials *OAuth2Credentials `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// The redirect URI passed to the authorization endpoint. Supports header formatting
	// tokens. For more information, including details on header value syntax, see the
	// documentation on :ref:`custom request headers <config_http_conn_man_headers_custom_request_headers>`.
	//
	// This URI should not contain any query parameters.
	RedirectUri string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// Matching criteria used to determine whether a path appears to be the result of a redirect from the authorization server.
	RedirectPathMatcher *v32.PathMatcher `protobuf:"bytes,5,opt,name=redirect_path_matcher,json=redirectPathMatcher,proto3" json:"redirect_path_matcher,omitempty"`
	// The path to sign a user out, clearing their credential cookies.
	SignoutPath *v32.PathMatcher `protobuf:"bytes,6,opt,name=signout_path,json=signoutPath,proto3" json:"signout_path,omitempty"`
	// Forward the OAuth token as a Bearer to upstream web service.
	ForwardBearerToken bool `protobuf:"varint,7,opt,name=forward_bearer_token,json=forwardBearerToken,proto3" json:"forward_bearer_token,omitempty"`
	// Any request that matches any of the provided matchers will be passed through without OAuth validation.
	PassThroughMatcher []*v33.HeaderMatcher `protobuf:"bytes,8,rep,name=pass_through_matcher,json=passThroughMatcher,proto3" json:"pass_through_matcher,omitempty"`
	// Optional list of OAuth scopes to be claimed in the authorization request. If not specified,
	// defaults to "user" scope.
	// OAuth RFC https://tools.ietf.org/html/rfc6749#section-3.3
	AuthScopes []string `protobuf:"bytes,9,rep,name=auth_scopes,json=authScopes,proto3" json:"auth_scopes,omitempty"`
	// Optional resource parameter for authorization request
	// RFC: https://tools.ietf.org/html/rfc8707
	Resources []string `protobuf:"bytes,10,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *OAuth2Config) Reset() {
	*x = OAuth2Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Config) ProtoMessage() {}

func (x *OAuth2Config) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Config.ProtoReflect.Descriptor instead.
func (*OAuth2Config) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *OAuth2Config) GetTokenEndpoint() *v31.HttpUri {
	if x != nil {
		return x.TokenEndpoint
	}
	return nil
}

func (x *OAuth2Config) GetAuthorizationEndpoint() string {
	if x != nil {
		return x.AuthorizationEndpoint
	}
	return ""
}

func (x *OAuth2Config) GetCredentials() *OAuth2Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *OAuth2Config) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *OAuth2Config) GetRedirectPathMatcher() *v32.PathMatcher {
	if x != nil {
		return x.RedirectPathMatcher
	}
	return nil
}

func (x *OAuth2Config) GetSignoutPath() *v32.PathMatcher {
	if x != nil {
		return x.SignoutPath
	}
	return nil
}

func (x *OAuth2Config) GetForwardBearerToken() bool {
	if x != nil {
		return x.ForwardBearerToken
	}
	return false
}

func (x *OAuth2Config) GetPassThroughMatcher() []*v33.HeaderMatcher {
	if x != nil {
		return x.PassThroughMatcher
	}
	return nil
}

func (x *OAuth2Config) GetAuthScopes() []string {
	if x != nil {
		return x.AuthScopes
	}
	return nil
}

func (x *OAuth2Config) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

// Filter config.
type OAuth2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Leave this empty to disable OAuth2 for a specific route, using per filter config.
	Config *OAuth2Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *OAuth2) Reset() {
	*x = OAuth2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2) ProtoMessage() {}

func (x *OAuth2) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2.ProtoReflect.Descriptor instead.
func (*OAuth2) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *OAuth2) GetConfig() *OAuth2Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type OAuth2Credentials_CookieNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cookie name to hold OAuth bearer token value. When the authentication server validates the
	// client and returns an authorization token back to the OAuth filter, no matter what format
	// that token is, if :ref:`forward_bearer_token <envoy_v3_api_field_extensions.filters.http.oauth2.v3.OAuth2Config.forward_bearer_token>`
	// is set to true the filter will send over the bearer token as a cookie with this name to the
	// upstream. Defaults to ``BearerToken``.
	BearerToken string `protobuf:"bytes,1,opt,name=bearer_token,json=bearerToken,proto3" json:"bearer_token,omitempty"`
	// Cookie name to hold OAuth HMAC value. Defaults to ``OauthHMAC``.
	OauthHmac string `protobuf:"bytes,2,opt,name=oauth_hmac,json=oauthHmac,proto3" json:"oauth_hmac,omitempty"`
	// Cookie name to hold OAuth expiry value. Defaults to ``OauthExpires``.
	OauthExpires string `protobuf:"bytes,3,opt,name=oauth_expires,json=oauthExpires,proto3" json:"oauth_expires,omitempty"`
}

func (x *OAuth2Credentials_CookieNames) Reset() {
	*x = OAuth2Credentials_CookieNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Credentials_CookieNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Credentials_CookieNames) ProtoMessage() {}

func (x *OAuth2Credentials_CookieNames) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Credentials_CookieNames.ProtoReflect.Descriptor instead.
func (*OAuth2Credentials_CookieNames) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OAuth2Credentials_CookieNames) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

func (x *OAuth2Credentials_CookieNames) GetOauthHmac() string {
	if x != nil {
		return x.OauthHmac
	}
	return ""
}

func (x *OAuth2Credentials_CookieNames) GetOauthExpires() string {
	if x != nil {
		return x.OauthExpires
	}
	return ""
}

var File_envoy_extensions_filters_http_oauth2_v3_oauth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x1a, 0x23,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xac, 0x04, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x67, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x67, 0x0a, 0x0b, 0x68, 0x6d, 0x61, 0x63,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x68, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x69, 0x0a, 0x0c, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76,
	0x33, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x9b, 0x01, 0x0a,
	0x0b, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0c,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52,
	0x0b, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x0a,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x09, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x48, 0x6d, 0x61, 0x63, 0x12, 0x30, 0x0a, 0x0d, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x0c, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x42, 0x16, 0x0a, 0x0f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x22, 0xa4, 0x05, 0x0a, 0x0c, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x16, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x60, 0x0a,
	0x15, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x4f, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x30, 0x0a, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x56, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75,
	0x67, 0x68, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x06, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x32, 0x12, 0x4d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0xa6, 0x01, 0x0a, 0x35, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x33, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x32, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescData = file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDesc
)

func file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDescData
}

var file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_goTypes = []interface{}{
	(*OAuth2Credentials)(nil),             // 0: envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials
	(*OAuth2Config)(nil),                  // 1: envoy.extensions.filters.http.oauth2.v3.OAuth2Config
	(*OAuth2)(nil),                        // 2: envoy.extensions.filters.http.oauth2.v3.OAuth2
	(*OAuth2Credentials_CookieNames)(nil), // 3: envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials.CookieNames
	(*v3.SdsSecretConfig)(nil),            // 4: envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	(*v31.HttpUri)(nil),                   // 5: envoy.config.core.v3.HttpUri
	(*v32.PathMatcher)(nil),               // 6: envoy.type.matcher.v3.PathMatcher
	(*v33.HeaderMatcher)(nil),             // 7: envoy.config.route.v3.HeaderMatcher
}
var file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials.token_secret:type_name -> envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	4, // 1: envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials.hmac_secret:type_name -> envoy.extensions.transport_sockets.tls.v3.SdsSecretConfig
	3, // 2: envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials.cookie_names:type_name -> envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials.CookieNames
	5, // 3: envoy.extensions.filters.http.oauth2.v3.OAuth2Config.token_endpoint:type_name -> envoy.config.core.v3.HttpUri
	0, // 4: envoy.extensions.filters.http.oauth2.v3.OAuth2Config.credentials:type_name -> envoy.extensions.filters.http.oauth2.v3.OAuth2Credentials
	6, // 5: envoy.extensions.filters.http.oauth2.v3.OAuth2Config.redirect_path_matcher:type_name -> envoy.type.matcher.v3.PathMatcher
	6, // 6: envoy.extensions.filters.http.oauth2.v3.OAuth2Config.signout_path:type_name -> envoy.type.matcher.v3.PathMatcher
	7, // 7: envoy.extensions.filters.http.oauth2.v3.OAuth2Config.pass_through_matcher:type_name -> envoy.config.route.v3.HeaderMatcher
	1, // 8: envoy.extensions.filters.http.oauth2.v3.OAuth2.config:type_name -> envoy.extensions.filters.http.oauth2.v3.OAuth2Config
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_init() }
func file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_init() {
	if File_envoy_extensions_filters_http_oauth2_v3_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Credentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Credentials_CookieNames); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OAuth2Credentials_HmacSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_oauth2_v3_oauth_proto = out.File
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_rawDesc = nil
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_goTypes = nil
	file_envoy_extensions_filters_http_oauth2_v3_oauth_proto_depIdxs = nil
}
