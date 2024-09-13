// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceSourceAPIAssociation,
			Name:    "Source API Association",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAPICache,
			TypeName: "aws_appsync_api_cache",
			Name:     "API Cache",
		},
		{
			Factory:  resourceAPIKey,
			TypeName: "aws_appsync_api_key",
			Name:     "API Key",
		},
		{
			Factory:  resourceDataSource,
			TypeName: "aws_appsync_datasource",
			Name:     "Data Source",
		},
		{
			Factory:  resourceDomainName,
			TypeName: "aws_appsync_domain_name",
			Name:     "Domain Name",
		},
		{
			Factory:  resourceDomainNameAPIAssociation,
			TypeName: "aws_appsync_domain_name_api_association",
			Name:     "Domain Name API Association",
		},
		{
			Factory:  resourceFunction,
			TypeName: "aws_appsync_function",
			Name:     "Function",
		},
		{
			Factory:  resourceGraphQLAPI,
			TypeName: "aws_appsync_graphql_api",
			Name:     "GraphQL API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceResolver,
			TypeName: "aws_appsync_resolver",
			Name:     "Resolver",
		},
		{
			Factory:  resourceType,
			TypeName: "aws_appsync_type",
			Name:     "Type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppSync
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
