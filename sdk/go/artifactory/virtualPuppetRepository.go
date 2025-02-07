// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Puppet Repository Resource
//
// Provides an Artifactory virtual repository resource with specific puppet features.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewVirtualPuppetRepository(ctx, "foo-puppet", &artifactory.VirtualPuppetRepositoryArgs{
// 			Description:     pulumi.String("A test virtual repo"),
// 			ExcludesPattern: pulumi.String("com/google/**"),
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("foo-puppet"),
// 			Notes:           pulumi.String("Internal description"),
// 			Repositories:    pulumi.StringArray{},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/virtualPuppetRepository:VirtualPuppetRepository foo foo
// ```
type VirtualPuppetRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringOutput `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualPuppetRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualPuppetRepository(ctx *pulumi.Context,
	name string, args *VirtualPuppetRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualPuppetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualPuppetRepository
	err := ctx.RegisterResource("artifactory:index/virtualPuppetRepository:VirtualPuppetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualPuppetRepository gets an existing VirtualPuppetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualPuppetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualPuppetRepositoryState, opts ...pulumi.ResourceOption) (*VirtualPuppetRepository, error) {
	var resource VirtualPuppetRepository
	err := ctx.ReadResource("artifactory:index/virtualPuppetRepository:VirtualPuppetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualPuppetRepository resources.
type virtualPuppetRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key *string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualPuppetRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualPuppetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPuppetRepositoryState)(nil)).Elem()
}

type virtualPuppetRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualPuppetRepository resource.
type VirtualPuppetRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualPuppetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPuppetRepositoryArgs)(nil)).Elem()
}

type VirtualPuppetRepositoryInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput
	ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput
}

func (*VirtualPuppetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPuppetRepository)(nil)).Elem()
}

func (i *VirtualPuppetRepository) ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput {
	return i.ToVirtualPuppetRepositoryOutputWithContext(context.Background())
}

func (i *VirtualPuppetRepository) ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryOutput)
}

// VirtualPuppetRepositoryArrayInput is an input type that accepts VirtualPuppetRepositoryArray and VirtualPuppetRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualPuppetRepositoryArrayInput` via:
//
//          VirtualPuppetRepositoryArray{ VirtualPuppetRepositoryArgs{...} }
type VirtualPuppetRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput
	ToVirtualPuppetRepositoryArrayOutputWithContext(context.Context) VirtualPuppetRepositoryArrayOutput
}

type VirtualPuppetRepositoryArray []VirtualPuppetRepositoryInput

func (VirtualPuppetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPuppetRepository)(nil)).Elem()
}

func (i VirtualPuppetRepositoryArray) ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput {
	return i.ToVirtualPuppetRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualPuppetRepositoryArray) ToVirtualPuppetRepositoryArrayOutputWithContext(ctx context.Context) VirtualPuppetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryArrayOutput)
}

// VirtualPuppetRepositoryMapInput is an input type that accepts VirtualPuppetRepositoryMap and VirtualPuppetRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualPuppetRepositoryMapInput` via:
//
//          VirtualPuppetRepositoryMap{ "key": VirtualPuppetRepositoryArgs{...} }
type VirtualPuppetRepositoryMapInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput
	ToVirtualPuppetRepositoryMapOutputWithContext(context.Context) VirtualPuppetRepositoryMapOutput
}

type VirtualPuppetRepositoryMap map[string]VirtualPuppetRepositoryInput

func (VirtualPuppetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPuppetRepository)(nil)).Elem()
}

func (i VirtualPuppetRepositoryMap) ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput {
	return i.ToVirtualPuppetRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualPuppetRepositoryMap) ToVirtualPuppetRepositoryMapOutputWithContext(ctx context.Context) VirtualPuppetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryMapOutput)
}

type VirtualPuppetRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryOutput) ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput {
	return o
}

func (o VirtualPuppetRepositoryOutput) ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput {
	return o
}

type VirtualPuppetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryArrayOutput) ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput {
	return o
}

func (o VirtualPuppetRepositoryArrayOutput) ToVirtualPuppetRepositoryArrayOutputWithContext(ctx context.Context) VirtualPuppetRepositoryArrayOutput {
	return o
}

func (o VirtualPuppetRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualPuppetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualPuppetRepository {
		return vs[0].([]*VirtualPuppetRepository)[vs[1].(int)]
	}).(VirtualPuppetRepositoryOutput)
}

type VirtualPuppetRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryMapOutput) ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput {
	return o
}

func (o VirtualPuppetRepositoryMapOutput) ToVirtualPuppetRepositoryMapOutputWithContext(ctx context.Context) VirtualPuppetRepositoryMapOutput {
	return o
}

func (o VirtualPuppetRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualPuppetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualPuppetRepository {
		return vs[0].(map[string]*VirtualPuppetRepository)[vs[1].(string)]
	}).(VirtualPuppetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryInput)(nil)).Elem(), &VirtualPuppetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryArrayInput)(nil)).Elem(), VirtualPuppetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryMapInput)(nil)).Elem(), VirtualPuppetRepositoryMap{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryMapOutput{})
}
