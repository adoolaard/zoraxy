package dynamicproxy

import (
	"github.com/google/uuid"
	"imuslab.com/zoraxy/mod/dynamicproxy/loadbalance"
	"imuslab.com/zoraxy/mod/dynamicproxy/rewrite"
)

/*
	Default Provider

	This script provide the default options for all datatype
	provided by dynamicproxy module

*/

// GetDefaultAuthenticationProvider return a default authentication provider
func GetDefaultAuthenticationProvider() *AuthenticationProvider {
	return &AuthenticationProvider{
		AuthMethod:              AuthMethodNone,
		BasicAuthCredentials:    []*BasicAuthCredentials{},
		BasicAuthExceptionRules: []*BasicAuthExceptionRule{},
		BasicAuthGroupIDs:       []string{},
		AutheliaURL:             "",
		UseHTTPS:                false,
	}
}

// GetDefaultHeaderRewriteRules return a default header rewrite rules
func GetDefaultHeaderRewriteRules() *HeaderRewriteRules {
	return &HeaderRewriteRules{
		UserDefinedHeaders:           []*rewrite.UserDefinedHeader{},
		RequestHostOverwrite:         "",
		HSTSMaxAge:                   0,
		EnablePermissionPolicyHeader: false,
		PermissionPolicy:             nil,
		DisableHopByHopHeaderRemoval: false,
	}
}

// GetDefaultProxyEndpoint return a default proxy endpoint
func GetDefaultProxyEndpoint() ProxyEndpoint {
	randomPrefix := uuid.New().String()
	return ProxyEndpoint{
		ProxyType:              ProxyTypeHost,
		RootOrMatchingDomain:   randomPrefix + ".internal",
		MatchingDomainAlias:    []string{},
		ActiveOrigins:          []*loadbalance.Upstream{},
		InactiveOrigins:        []*loadbalance.Upstream{},
		UseStickySession:       false,
		UseActiveLoadBalance:   false,
		Disabled:               false,
		BypassGlobalTLS:        false,
		VirtualDirectories:     []*VirtualDirectoryEndpoint{},
		HeaderRewriteRules:     GetDefaultHeaderRewriteRules(),
		AuthenticationProvider: GetDefaultAuthenticationProvider(),
		RequireRateLimit:       false,
		RateLimit:              0,
		DisableUptimeMonitor:   false,
		AccessFilterUUID:       "default",
		DefaultSiteOption:      DefaultSite_InternalStaticWebServer,
		DefaultSiteValue:       "",
	}
}
