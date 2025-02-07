// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Debian Repository Resource
//
// Provides an Artifactory virtual repository resource with specific debian features.
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
// 		_, err := artifactory.NewVirtualDebianRepository(ctx, "foo-debian", &artifactory.VirtualDebianRepositoryArgs{
// 			DebianDefaultArchitectures: pulumi.String("amd64,i386"),
// 			Description:                pulumi.String("A test virtual repo"),
// 			ExcludesPattern:            pulumi.String("com/google/**"),
// 			IncludesPattern:            pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:                        pulumi.String("foo-debian"),
// 			Notes:                      pulumi.String("Internal description"),
// 			OptionalIndexCompressionFormats: pulumi.StringArray{
// 				pulumi.String("bz2"),
// 				pulumi.String("xz"),
// 			},
// 			Repositories: pulumi.StringArray{},
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
//  $ pulumi import artifactory:index/virtualDebianRepository:VirtualDebianRepository foo foo
// ```
type VirtualDebianRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures pulumi.StringPtrOutput `pulumi:"debianDefaultArchitectures"`
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
	// Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.
	OptionalIndexCompressionFormats pulumi.StringArrayOutput `pulumi:"optionalIndexCompressionFormats"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
}

// NewVirtualDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualDebianRepository(ctx *pulumi.Context,
	name string, args *VirtualDebianRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualDebianRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualDebianRepository
	err := ctx.RegisterResource("artifactory:index/virtualDebianRepository:VirtualDebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualDebianRepository gets an existing VirtualDebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualDebianRepositoryState, opts ...pulumi.ResourceOption) (*VirtualDebianRepository, error) {
	var resource VirtualDebianRepository
	err := ctx.ReadResource("artifactory:index/virtualDebianRepository:VirtualDebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualDebianRepository resources.
type virtualDebianRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures *string `pulumi:"debianDefaultArchitectures"`
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
	// Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.
	OptionalIndexCompressionFormats []string `pulumi:"optionalIndexCompressionFormats"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

type VirtualDebianRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures pulumi.StringPtrInput
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
	// Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.
	OptionalIndexCompressionFormats pulumi.StringArrayInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualDebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDebianRepositoryState)(nil)).Elem()
}

type virtualDebianRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures *string `pulumi:"debianDefaultArchitectures"`
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
	// Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.
	OptionalIndexCompressionFormats []string `pulumi:"optionalIndexCompressionFormats"`
	// Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

// The set of arguments for constructing a VirtualDebianRepository resource.
type VirtualDebianRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures pulumi.StringPtrInput
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
	// Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.
	OptionalIndexCompressionFormats pulumi.StringArrayInput
	// Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualDebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDebianRepositoryArgs)(nil)).Elem()
}

type VirtualDebianRepositoryInput interface {
	pulumi.Input

	ToVirtualDebianRepositoryOutput() VirtualDebianRepositoryOutput
	ToVirtualDebianRepositoryOutputWithContext(ctx context.Context) VirtualDebianRepositoryOutput
}

func (*VirtualDebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDebianRepository)(nil)).Elem()
}

func (i *VirtualDebianRepository) ToVirtualDebianRepositoryOutput() VirtualDebianRepositoryOutput {
	return i.ToVirtualDebianRepositoryOutputWithContext(context.Background())
}

func (i *VirtualDebianRepository) ToVirtualDebianRepositoryOutputWithContext(ctx context.Context) VirtualDebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDebianRepositoryOutput)
}

// VirtualDebianRepositoryArrayInput is an input type that accepts VirtualDebianRepositoryArray and VirtualDebianRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualDebianRepositoryArrayInput` via:
//
//          VirtualDebianRepositoryArray{ VirtualDebianRepositoryArgs{...} }
type VirtualDebianRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualDebianRepositoryArrayOutput() VirtualDebianRepositoryArrayOutput
	ToVirtualDebianRepositoryArrayOutputWithContext(context.Context) VirtualDebianRepositoryArrayOutput
}

type VirtualDebianRepositoryArray []VirtualDebianRepositoryInput

func (VirtualDebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDebianRepository)(nil)).Elem()
}

func (i VirtualDebianRepositoryArray) ToVirtualDebianRepositoryArrayOutput() VirtualDebianRepositoryArrayOutput {
	return i.ToVirtualDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualDebianRepositoryArray) ToVirtualDebianRepositoryArrayOutputWithContext(ctx context.Context) VirtualDebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDebianRepositoryArrayOutput)
}

// VirtualDebianRepositoryMapInput is an input type that accepts VirtualDebianRepositoryMap and VirtualDebianRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualDebianRepositoryMapInput` via:
//
//          VirtualDebianRepositoryMap{ "key": VirtualDebianRepositoryArgs{...} }
type VirtualDebianRepositoryMapInput interface {
	pulumi.Input

	ToVirtualDebianRepositoryMapOutput() VirtualDebianRepositoryMapOutput
	ToVirtualDebianRepositoryMapOutputWithContext(context.Context) VirtualDebianRepositoryMapOutput
}

type VirtualDebianRepositoryMap map[string]VirtualDebianRepositoryInput

func (VirtualDebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDebianRepository)(nil)).Elem()
}

func (i VirtualDebianRepositoryMap) ToVirtualDebianRepositoryMapOutput() VirtualDebianRepositoryMapOutput {
	return i.ToVirtualDebianRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualDebianRepositoryMap) ToVirtualDebianRepositoryMapOutputWithContext(ctx context.Context) VirtualDebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDebianRepositoryMapOutput)
}

type VirtualDebianRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualDebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDebianRepository)(nil)).Elem()
}

func (o VirtualDebianRepositoryOutput) ToVirtualDebianRepositoryOutput() VirtualDebianRepositoryOutput {
	return o
}

func (o VirtualDebianRepositoryOutput) ToVirtualDebianRepositoryOutputWithContext(ctx context.Context) VirtualDebianRepositoryOutput {
	return o
}

type VirtualDebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDebianRepository)(nil)).Elem()
}

func (o VirtualDebianRepositoryArrayOutput) ToVirtualDebianRepositoryArrayOutput() VirtualDebianRepositoryArrayOutput {
	return o
}

func (o VirtualDebianRepositoryArrayOutput) ToVirtualDebianRepositoryArrayOutputWithContext(ctx context.Context) VirtualDebianRepositoryArrayOutput {
	return o
}

func (o VirtualDebianRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualDebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualDebianRepository {
		return vs[0].([]*VirtualDebianRepository)[vs[1].(int)]
	}).(VirtualDebianRepositoryOutput)
}

type VirtualDebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualDebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDebianRepository)(nil)).Elem()
}

func (o VirtualDebianRepositoryMapOutput) ToVirtualDebianRepositoryMapOutput() VirtualDebianRepositoryMapOutput {
	return o
}

func (o VirtualDebianRepositoryMapOutput) ToVirtualDebianRepositoryMapOutputWithContext(ctx context.Context) VirtualDebianRepositoryMapOutput {
	return o
}

func (o VirtualDebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualDebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualDebianRepository {
		return vs[0].(map[string]*VirtualDebianRepository)[vs[1].(string)]
	}).(VirtualDebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDebianRepositoryInput)(nil)).Elem(), &VirtualDebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDebianRepositoryArrayInput)(nil)).Elem(), VirtualDebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDebianRepositoryMapInput)(nil)).Elem(), VirtualDebianRepositoryMap{})
	pulumi.RegisterOutputType(VirtualDebianRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualDebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDebianRepositoryMapOutput{})
}
