// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Go Repository Resource
//
// Provides an Artifactory virtual repository resource with specific go lang features.
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
// 		_, err := artifactory.NewGoRepository(ctx, "baz-go", &artifactory.GoRepositoryArgs{
// 			Description:                 pulumi.String("A test virtual repo"),
// 			ExcludesPattern:             pulumi.String("com/google/**"),
// 			ExternalDependenciesEnabled: pulumi.Bool(true),
// 			ExternalDependenciesPatterns: pulumi.StringArray{
// 				pulumi.String("**/github.com/**"),
// 				pulumi.String("**/go.googlesource.com/**"),
// 			},
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("baz-go"),
// 			Notes:           pulumi.String("Internal description"),
// 			RepoLayoutRef:   pulumi.String("go-default"),
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
//  $ pulumi import artifactory:index/goRepository:GoRepository foo foo
// ```
type GoRepository struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list
	ExternalDependenciesEnabled pulumi.BoolPtrOutput `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayOutput `pulumi:"externalDependenciesPatterns"`
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

// NewGoRepository registers a new resource with the given unique name, arguments, and options.
func NewGoRepository(ctx *pulumi.Context,
	name string, args *GoRepositoryArgs, opts ...pulumi.ResourceOption) (*GoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource GoRepository
	err := ctx.RegisterResource("artifactory:index/goRepository:GoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGoRepository gets an existing GoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GoRepositoryState, opts ...pulumi.ResourceOption) (*GoRepository, error) {
	var resource GoRepository
	err := ctx.ReadResource("artifactory:index/goRepository:GoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GoRepository resources.
type goRepositoryState struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
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

type GoRepositoryState struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayInput
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

func (GoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*goRepositoryState)(nil)).Elem()
}

type goRepositoryArgs struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
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

// The set of arguments for constructing a GoRepository resource.
type GoRepositoryArgs struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayInput
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

func (GoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*goRepositoryArgs)(nil)).Elem()
}

type GoRepositoryInput interface {
	pulumi.Input

	ToGoRepositoryOutput() GoRepositoryOutput
	ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput
}

func (*GoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**GoRepository)(nil)).Elem()
}

func (i *GoRepository) ToGoRepositoryOutput() GoRepositoryOutput {
	return i.ToGoRepositoryOutputWithContext(context.Background())
}

func (i *GoRepository) ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryOutput)
}

// GoRepositoryArrayInput is an input type that accepts GoRepositoryArray and GoRepositoryArrayOutput values.
// You can construct a concrete instance of `GoRepositoryArrayInput` via:
//
//          GoRepositoryArray{ GoRepositoryArgs{...} }
type GoRepositoryArrayInput interface {
	pulumi.Input

	ToGoRepositoryArrayOutput() GoRepositoryArrayOutput
	ToGoRepositoryArrayOutputWithContext(context.Context) GoRepositoryArrayOutput
}

type GoRepositoryArray []GoRepositoryInput

func (GoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoRepository)(nil)).Elem()
}

func (i GoRepositoryArray) ToGoRepositoryArrayOutput() GoRepositoryArrayOutput {
	return i.ToGoRepositoryArrayOutputWithContext(context.Background())
}

func (i GoRepositoryArray) ToGoRepositoryArrayOutputWithContext(ctx context.Context) GoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryArrayOutput)
}

// GoRepositoryMapInput is an input type that accepts GoRepositoryMap and GoRepositoryMapOutput values.
// You can construct a concrete instance of `GoRepositoryMapInput` via:
//
//          GoRepositoryMap{ "key": GoRepositoryArgs{...} }
type GoRepositoryMapInput interface {
	pulumi.Input

	ToGoRepositoryMapOutput() GoRepositoryMapOutput
	ToGoRepositoryMapOutputWithContext(context.Context) GoRepositoryMapOutput
}

type GoRepositoryMap map[string]GoRepositoryInput

func (GoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoRepository)(nil)).Elem()
}

func (i GoRepositoryMap) ToGoRepositoryMapOutput() GoRepositoryMapOutput {
	return i.ToGoRepositoryMapOutputWithContext(context.Background())
}

func (i GoRepositoryMap) ToGoRepositoryMapOutputWithContext(ctx context.Context) GoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryMapOutput)
}

type GoRepositoryOutput struct{ *pulumi.OutputState }

func (GoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoRepository)(nil)).Elem()
}

func (o GoRepositoryOutput) ToGoRepositoryOutput() GoRepositoryOutput {
	return o
}

func (o GoRepositoryOutput) ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput {
	return o
}

type GoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (GoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoRepository)(nil)).Elem()
}

func (o GoRepositoryArrayOutput) ToGoRepositoryArrayOutput() GoRepositoryArrayOutput {
	return o
}

func (o GoRepositoryArrayOutput) ToGoRepositoryArrayOutputWithContext(ctx context.Context) GoRepositoryArrayOutput {
	return o
}

func (o GoRepositoryArrayOutput) Index(i pulumi.IntInput) GoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GoRepository {
		return vs[0].([]*GoRepository)[vs[1].(int)]
	}).(GoRepositoryOutput)
}

type GoRepositoryMapOutput struct{ *pulumi.OutputState }

func (GoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoRepository)(nil)).Elem()
}

func (o GoRepositoryMapOutput) ToGoRepositoryMapOutput() GoRepositoryMapOutput {
	return o
}

func (o GoRepositoryMapOutput) ToGoRepositoryMapOutputWithContext(ctx context.Context) GoRepositoryMapOutput {
	return o
}

func (o GoRepositoryMapOutput) MapIndex(k pulumi.StringInput) GoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GoRepository {
		return vs[0].(map[string]*GoRepository)[vs[1].(string)]
	}).(GoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryInput)(nil)).Elem(), &GoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryArrayInput)(nil)).Elem(), GoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryMapInput)(nil)).Elem(), GoRepositoryMap{})
	pulumi.RegisterOutputType(GoRepositoryOutput{})
	pulumi.RegisterOutputType(GoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(GoRepositoryMapOutput{})
}
