// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Docker Repository Resource
//
// Provides an Artifactory virtual repository resource with specific docker features.
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
// 		_, err := artifactory.NewVirtualDockerRepository(ctx, "foo-docker", &artifactory.VirtualDockerRepositoryArgs{
// 			Description:     pulumi.String("A test virtual repo"),
// 			ExcludesPattern: pulumi.String("com/google/**"),
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("foo-docker"),
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
//  $ pulumi import artifactory:index/virtualDockerRepository:VirtualDockerRepository foo foo
// ```
type VirtualDockerRepository struct {
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

// NewVirtualDockerRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualDockerRepository(ctx *pulumi.Context,
	name string, args *VirtualDockerRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualDockerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualDockerRepository
	err := ctx.RegisterResource("artifactory:index/virtualDockerRepository:VirtualDockerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualDockerRepository gets an existing VirtualDockerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDockerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualDockerRepositoryState, opts ...pulumi.ResourceOption) (*VirtualDockerRepository, error) {
	var resource VirtualDockerRepository
	err := ctx.ReadResource("artifactory:index/virtualDockerRepository:VirtualDockerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualDockerRepository resources.
type virtualDockerRepositoryState struct {
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

type VirtualDockerRepositoryState struct {
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

func (VirtualDockerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDockerRepositoryState)(nil)).Elem()
}

type virtualDockerRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualDockerRepository resource.
type VirtualDockerRepositoryArgs struct {
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

func (VirtualDockerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDockerRepositoryArgs)(nil)).Elem()
}

type VirtualDockerRepositoryInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput
	ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput
}

func (*VirtualDockerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDockerRepository)(nil)).Elem()
}

func (i *VirtualDockerRepository) ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput {
	return i.ToVirtualDockerRepositoryOutputWithContext(context.Background())
}

func (i *VirtualDockerRepository) ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryOutput)
}

// VirtualDockerRepositoryArrayInput is an input type that accepts VirtualDockerRepositoryArray and VirtualDockerRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualDockerRepositoryArrayInput` via:
//
//          VirtualDockerRepositoryArray{ VirtualDockerRepositoryArgs{...} }
type VirtualDockerRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput
	ToVirtualDockerRepositoryArrayOutputWithContext(context.Context) VirtualDockerRepositoryArrayOutput
}

type VirtualDockerRepositoryArray []VirtualDockerRepositoryInput

func (VirtualDockerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDockerRepository)(nil)).Elem()
}

func (i VirtualDockerRepositoryArray) ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput {
	return i.ToVirtualDockerRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualDockerRepositoryArray) ToVirtualDockerRepositoryArrayOutputWithContext(ctx context.Context) VirtualDockerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryArrayOutput)
}

// VirtualDockerRepositoryMapInput is an input type that accepts VirtualDockerRepositoryMap and VirtualDockerRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualDockerRepositoryMapInput` via:
//
//          VirtualDockerRepositoryMap{ "key": VirtualDockerRepositoryArgs{...} }
type VirtualDockerRepositoryMapInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput
	ToVirtualDockerRepositoryMapOutputWithContext(context.Context) VirtualDockerRepositoryMapOutput
}

type VirtualDockerRepositoryMap map[string]VirtualDockerRepositoryInput

func (VirtualDockerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDockerRepository)(nil)).Elem()
}

func (i VirtualDockerRepositoryMap) ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput {
	return i.ToVirtualDockerRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualDockerRepositoryMap) ToVirtualDockerRepositoryMapOutputWithContext(ctx context.Context) VirtualDockerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryMapOutput)
}

type VirtualDockerRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryOutput) ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput {
	return o
}

func (o VirtualDockerRepositoryOutput) ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput {
	return o
}

type VirtualDockerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryArrayOutput) ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput {
	return o
}

func (o VirtualDockerRepositoryArrayOutput) ToVirtualDockerRepositoryArrayOutputWithContext(ctx context.Context) VirtualDockerRepositoryArrayOutput {
	return o
}

func (o VirtualDockerRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualDockerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualDockerRepository {
		return vs[0].([]*VirtualDockerRepository)[vs[1].(int)]
	}).(VirtualDockerRepositoryOutput)
}

type VirtualDockerRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryMapOutput) ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput {
	return o
}

func (o VirtualDockerRepositoryMapOutput) ToVirtualDockerRepositoryMapOutputWithContext(ctx context.Context) VirtualDockerRepositoryMapOutput {
	return o
}

func (o VirtualDockerRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualDockerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualDockerRepository {
		return vs[0].(map[string]*VirtualDockerRepository)[vs[1].(string)]
	}).(VirtualDockerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryInput)(nil)).Elem(), &VirtualDockerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryArrayInput)(nil)).Elem(), VirtualDockerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryMapInput)(nil)).Elem(), VirtualDockerRepositoryMap{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryMapOutput{})
}
