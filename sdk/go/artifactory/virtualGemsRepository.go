// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Gems Repository Resource
//
// Provides an Artifactory virtual repository resource, but with specific gems features. This should be preferred over the original
// one-size-fits-all `VirtualRepository`.
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
// 		_, err := artifactory.NewVirtualGemsRepository(ctx, "foo-gems", &artifactory.VirtualGemsRepositoryArgs{
// 			Description:     pulumi.String("A test virtual repo"),
// 			ExcludesPattern: pulumi.String("com/google/**"),
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("foo-gems"),
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
//  $ pulumi import artifactory:index/virtualGemsRepository:VirtualGemsRepository foo foo
// ```
type VirtualGemsRepository struct {
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

// NewVirtualGemsRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualGemsRepository(ctx *pulumi.Context,
	name string, args *VirtualGemsRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualGemsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualGemsRepository
	err := ctx.RegisterResource("artifactory:index/virtualGemsRepository:VirtualGemsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGemsRepository gets an existing VirtualGemsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGemsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGemsRepositoryState, opts ...pulumi.ResourceOption) (*VirtualGemsRepository, error) {
	var resource VirtualGemsRepository
	err := ctx.ReadResource("artifactory:index/virtualGemsRepository:VirtualGemsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGemsRepository resources.
type virtualGemsRepositoryState struct {
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

type VirtualGemsRepositoryState struct {
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

func (VirtualGemsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGemsRepositoryState)(nil)).Elem()
}

type virtualGemsRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualGemsRepository resource.
type VirtualGemsRepositoryArgs struct {
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

func (VirtualGemsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGemsRepositoryArgs)(nil)).Elem()
}

type VirtualGemsRepositoryInput interface {
	pulumi.Input

	ToVirtualGemsRepositoryOutput() VirtualGemsRepositoryOutput
	ToVirtualGemsRepositoryOutputWithContext(ctx context.Context) VirtualGemsRepositoryOutput
}

func (*VirtualGemsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGemsRepository)(nil)).Elem()
}

func (i *VirtualGemsRepository) ToVirtualGemsRepositoryOutput() VirtualGemsRepositoryOutput {
	return i.ToVirtualGemsRepositoryOutputWithContext(context.Background())
}

func (i *VirtualGemsRepository) ToVirtualGemsRepositoryOutputWithContext(ctx context.Context) VirtualGemsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGemsRepositoryOutput)
}

// VirtualGemsRepositoryArrayInput is an input type that accepts VirtualGemsRepositoryArray and VirtualGemsRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualGemsRepositoryArrayInput` via:
//
//          VirtualGemsRepositoryArray{ VirtualGemsRepositoryArgs{...} }
type VirtualGemsRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualGemsRepositoryArrayOutput() VirtualGemsRepositoryArrayOutput
	ToVirtualGemsRepositoryArrayOutputWithContext(context.Context) VirtualGemsRepositoryArrayOutput
}

type VirtualGemsRepositoryArray []VirtualGemsRepositoryInput

func (VirtualGemsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGemsRepository)(nil)).Elem()
}

func (i VirtualGemsRepositoryArray) ToVirtualGemsRepositoryArrayOutput() VirtualGemsRepositoryArrayOutput {
	return i.ToVirtualGemsRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualGemsRepositoryArray) ToVirtualGemsRepositoryArrayOutputWithContext(ctx context.Context) VirtualGemsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGemsRepositoryArrayOutput)
}

// VirtualGemsRepositoryMapInput is an input type that accepts VirtualGemsRepositoryMap and VirtualGemsRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualGemsRepositoryMapInput` via:
//
//          VirtualGemsRepositoryMap{ "key": VirtualGemsRepositoryArgs{...} }
type VirtualGemsRepositoryMapInput interface {
	pulumi.Input

	ToVirtualGemsRepositoryMapOutput() VirtualGemsRepositoryMapOutput
	ToVirtualGemsRepositoryMapOutputWithContext(context.Context) VirtualGemsRepositoryMapOutput
}

type VirtualGemsRepositoryMap map[string]VirtualGemsRepositoryInput

func (VirtualGemsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGemsRepository)(nil)).Elem()
}

func (i VirtualGemsRepositoryMap) ToVirtualGemsRepositoryMapOutput() VirtualGemsRepositoryMapOutput {
	return i.ToVirtualGemsRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualGemsRepositoryMap) ToVirtualGemsRepositoryMapOutputWithContext(ctx context.Context) VirtualGemsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGemsRepositoryMapOutput)
}

type VirtualGemsRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualGemsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGemsRepository)(nil)).Elem()
}

func (o VirtualGemsRepositoryOutput) ToVirtualGemsRepositoryOutput() VirtualGemsRepositoryOutput {
	return o
}

func (o VirtualGemsRepositoryOutput) ToVirtualGemsRepositoryOutputWithContext(ctx context.Context) VirtualGemsRepositoryOutput {
	return o
}

type VirtualGemsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualGemsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGemsRepository)(nil)).Elem()
}

func (o VirtualGemsRepositoryArrayOutput) ToVirtualGemsRepositoryArrayOutput() VirtualGemsRepositoryArrayOutput {
	return o
}

func (o VirtualGemsRepositoryArrayOutput) ToVirtualGemsRepositoryArrayOutputWithContext(ctx context.Context) VirtualGemsRepositoryArrayOutput {
	return o
}

func (o VirtualGemsRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualGemsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGemsRepository {
		return vs[0].([]*VirtualGemsRepository)[vs[1].(int)]
	}).(VirtualGemsRepositoryOutput)
}

type VirtualGemsRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualGemsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGemsRepository)(nil)).Elem()
}

func (o VirtualGemsRepositoryMapOutput) ToVirtualGemsRepositoryMapOutput() VirtualGemsRepositoryMapOutput {
	return o
}

func (o VirtualGemsRepositoryMapOutput) ToVirtualGemsRepositoryMapOutputWithContext(ctx context.Context) VirtualGemsRepositoryMapOutput {
	return o
}

func (o VirtualGemsRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualGemsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGemsRepository {
		return vs[0].(map[string]*VirtualGemsRepository)[vs[1].(string)]
	}).(VirtualGemsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGemsRepositoryInput)(nil)).Elem(), &VirtualGemsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGemsRepositoryArrayInput)(nil)).Elem(), VirtualGemsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGemsRepositoryMapInput)(nil)).Elem(), VirtualGemsRepositoryMap{})
	pulumi.RegisterOutputType(VirtualGemsRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualGemsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualGemsRepositoryMapOutput{})
}