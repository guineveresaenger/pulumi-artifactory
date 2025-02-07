// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Ivy Repository Resource
//
// Provides an Artifactory virtual repository resource with specific ivy features.
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
// 		_, err := artifactory.NewVirtualIvyRepository(ctx, "foo-ivy", &artifactory.VirtualIvyRepositoryArgs{
// 			Description:                          pulumi.String("A test virtual repo"),
// 			ExcludesPattern:                      pulumi.String("com/google/**"),
// 			IncludesPattern:                      pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:                                  pulumi.String("foo-ivy"),
// 			Notes:                                pulumi.String("Internal description"),
// 			PomRepositoryReferencesCleanupPolicy: pulumi.String("discard_active_reference"),
// 			Repositories:                         pulumi.StringArray{},
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
//  $ pulumi import artifactory:index/virtualIvyRepository:VirtualIvyRepository foo foo
// ```
type VirtualIvyRepository struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolOutput `pulumi:"forceMavenAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringOutput `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringOutput `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

// NewVirtualIvyRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualIvyRepository(ctx *pulumi.Context,
	name string, args *VirtualIvyRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualIvyRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualIvyRepository
	err := ctx.RegisterResource("artifactory:index/virtualIvyRepository:VirtualIvyRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualIvyRepository gets an existing VirtualIvyRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualIvyRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualIvyRepositoryState, opts ...pulumi.ResourceOption) (*VirtualIvyRepository, error) {
	var resource VirtualIvyRepository
	err := ctx.ReadResource("artifactory:index/virtualIvyRepository:VirtualIvyRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualIvyRepository resources.
type virtualIvyRepositoryState struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key *string `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

type VirtualIvyRepositoryState struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringPtrInput
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
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

func (VirtualIvyRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualIvyRepositoryState)(nil)).Elem()
}

type virtualIvyRepositoryArgs struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key string `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

// The set of arguments for constructing a VirtualIvyRepository resource.
type VirtualIvyRepositoryArgs struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringInput
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
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

func (VirtualIvyRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualIvyRepositoryArgs)(nil)).Elem()
}

type VirtualIvyRepositoryInput interface {
	pulumi.Input

	ToVirtualIvyRepositoryOutput() VirtualIvyRepositoryOutput
	ToVirtualIvyRepositoryOutputWithContext(ctx context.Context) VirtualIvyRepositoryOutput
}

func (*VirtualIvyRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualIvyRepository)(nil)).Elem()
}

func (i *VirtualIvyRepository) ToVirtualIvyRepositoryOutput() VirtualIvyRepositoryOutput {
	return i.ToVirtualIvyRepositoryOutputWithContext(context.Background())
}

func (i *VirtualIvyRepository) ToVirtualIvyRepositoryOutputWithContext(ctx context.Context) VirtualIvyRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIvyRepositoryOutput)
}

// VirtualIvyRepositoryArrayInput is an input type that accepts VirtualIvyRepositoryArray and VirtualIvyRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualIvyRepositoryArrayInput` via:
//
//          VirtualIvyRepositoryArray{ VirtualIvyRepositoryArgs{...} }
type VirtualIvyRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualIvyRepositoryArrayOutput() VirtualIvyRepositoryArrayOutput
	ToVirtualIvyRepositoryArrayOutputWithContext(context.Context) VirtualIvyRepositoryArrayOutput
}

type VirtualIvyRepositoryArray []VirtualIvyRepositoryInput

func (VirtualIvyRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualIvyRepository)(nil)).Elem()
}

func (i VirtualIvyRepositoryArray) ToVirtualIvyRepositoryArrayOutput() VirtualIvyRepositoryArrayOutput {
	return i.ToVirtualIvyRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualIvyRepositoryArray) ToVirtualIvyRepositoryArrayOutputWithContext(ctx context.Context) VirtualIvyRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIvyRepositoryArrayOutput)
}

// VirtualIvyRepositoryMapInput is an input type that accepts VirtualIvyRepositoryMap and VirtualIvyRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualIvyRepositoryMapInput` via:
//
//          VirtualIvyRepositoryMap{ "key": VirtualIvyRepositoryArgs{...} }
type VirtualIvyRepositoryMapInput interface {
	pulumi.Input

	ToVirtualIvyRepositoryMapOutput() VirtualIvyRepositoryMapOutput
	ToVirtualIvyRepositoryMapOutputWithContext(context.Context) VirtualIvyRepositoryMapOutput
}

type VirtualIvyRepositoryMap map[string]VirtualIvyRepositoryInput

func (VirtualIvyRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualIvyRepository)(nil)).Elem()
}

func (i VirtualIvyRepositoryMap) ToVirtualIvyRepositoryMapOutput() VirtualIvyRepositoryMapOutput {
	return i.ToVirtualIvyRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualIvyRepositoryMap) ToVirtualIvyRepositoryMapOutputWithContext(ctx context.Context) VirtualIvyRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIvyRepositoryMapOutput)
}

type VirtualIvyRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualIvyRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualIvyRepository)(nil)).Elem()
}

func (o VirtualIvyRepositoryOutput) ToVirtualIvyRepositoryOutput() VirtualIvyRepositoryOutput {
	return o
}

func (o VirtualIvyRepositoryOutput) ToVirtualIvyRepositoryOutputWithContext(ctx context.Context) VirtualIvyRepositoryOutput {
	return o
}

type VirtualIvyRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualIvyRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualIvyRepository)(nil)).Elem()
}

func (o VirtualIvyRepositoryArrayOutput) ToVirtualIvyRepositoryArrayOutput() VirtualIvyRepositoryArrayOutput {
	return o
}

func (o VirtualIvyRepositoryArrayOutput) ToVirtualIvyRepositoryArrayOutputWithContext(ctx context.Context) VirtualIvyRepositoryArrayOutput {
	return o
}

func (o VirtualIvyRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualIvyRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualIvyRepository {
		return vs[0].([]*VirtualIvyRepository)[vs[1].(int)]
	}).(VirtualIvyRepositoryOutput)
}

type VirtualIvyRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualIvyRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualIvyRepository)(nil)).Elem()
}

func (o VirtualIvyRepositoryMapOutput) ToVirtualIvyRepositoryMapOutput() VirtualIvyRepositoryMapOutput {
	return o
}

func (o VirtualIvyRepositoryMapOutput) ToVirtualIvyRepositoryMapOutputWithContext(ctx context.Context) VirtualIvyRepositoryMapOutput {
	return o
}

func (o VirtualIvyRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualIvyRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualIvyRepository {
		return vs[0].(map[string]*VirtualIvyRepository)[vs[1].(string)]
	}).(VirtualIvyRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualIvyRepositoryInput)(nil)).Elem(), &VirtualIvyRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualIvyRepositoryArrayInput)(nil)).Elem(), VirtualIvyRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualIvyRepositoryMapInput)(nil)).Elem(), VirtualIvyRepositoryMap{})
	pulumi.RegisterOutputType(VirtualIvyRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualIvyRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualIvyRepositoryMapOutput{})
}
