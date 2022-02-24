// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Bower Repository Resource
//
// Creates a local bower repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewLocalBowerRepository(ctx, "terraform-local-test-bower-repo", &artifactory.LocalBowerRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-bower-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalBowerRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalBowerRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalBowerRepository(ctx *pulumi.Context,
	name string, args *LocalBowerRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalBowerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalBowerRepository
	err := ctx.RegisterResource("artifactory:index/localBowerRepository:LocalBowerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalBowerRepository gets an existing LocalBowerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalBowerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalBowerRepositoryState, opts ...pulumi.ResourceOption) (*LocalBowerRepository, error) {
	var resource LocalBowerRepository
	err := ctx.ReadResource("artifactory:index/localBowerRepository:LocalBowerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalBowerRepository resources.
type localBowerRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalBowerRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalBowerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localBowerRepositoryState)(nil)).Elem()
}

type localBowerRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalBowerRepository resource.
type LocalBowerRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalBowerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localBowerRepositoryArgs)(nil)).Elem()
}

type LocalBowerRepositoryInput interface {
	pulumi.Input

	ToLocalBowerRepositoryOutput() LocalBowerRepositoryOutput
	ToLocalBowerRepositoryOutputWithContext(ctx context.Context) LocalBowerRepositoryOutput
}

func (*LocalBowerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalBowerRepository)(nil)).Elem()
}

func (i *LocalBowerRepository) ToLocalBowerRepositoryOutput() LocalBowerRepositoryOutput {
	return i.ToLocalBowerRepositoryOutputWithContext(context.Background())
}

func (i *LocalBowerRepository) ToLocalBowerRepositoryOutputWithContext(ctx context.Context) LocalBowerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBowerRepositoryOutput)
}

// LocalBowerRepositoryArrayInput is an input type that accepts LocalBowerRepositoryArray and LocalBowerRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalBowerRepositoryArrayInput` via:
//
//          LocalBowerRepositoryArray{ LocalBowerRepositoryArgs{...} }
type LocalBowerRepositoryArrayInput interface {
	pulumi.Input

	ToLocalBowerRepositoryArrayOutput() LocalBowerRepositoryArrayOutput
	ToLocalBowerRepositoryArrayOutputWithContext(context.Context) LocalBowerRepositoryArrayOutput
}

type LocalBowerRepositoryArray []LocalBowerRepositoryInput

func (LocalBowerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalBowerRepository)(nil)).Elem()
}

func (i LocalBowerRepositoryArray) ToLocalBowerRepositoryArrayOutput() LocalBowerRepositoryArrayOutput {
	return i.ToLocalBowerRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalBowerRepositoryArray) ToLocalBowerRepositoryArrayOutputWithContext(ctx context.Context) LocalBowerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBowerRepositoryArrayOutput)
}

// LocalBowerRepositoryMapInput is an input type that accepts LocalBowerRepositoryMap and LocalBowerRepositoryMapOutput values.
// You can construct a concrete instance of `LocalBowerRepositoryMapInput` via:
//
//          LocalBowerRepositoryMap{ "key": LocalBowerRepositoryArgs{...} }
type LocalBowerRepositoryMapInput interface {
	pulumi.Input

	ToLocalBowerRepositoryMapOutput() LocalBowerRepositoryMapOutput
	ToLocalBowerRepositoryMapOutputWithContext(context.Context) LocalBowerRepositoryMapOutput
}

type LocalBowerRepositoryMap map[string]LocalBowerRepositoryInput

func (LocalBowerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalBowerRepository)(nil)).Elem()
}

func (i LocalBowerRepositoryMap) ToLocalBowerRepositoryMapOutput() LocalBowerRepositoryMapOutput {
	return i.ToLocalBowerRepositoryMapOutputWithContext(context.Background())
}

func (i LocalBowerRepositoryMap) ToLocalBowerRepositoryMapOutputWithContext(ctx context.Context) LocalBowerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalBowerRepositoryMapOutput)
}

type LocalBowerRepositoryOutput struct{ *pulumi.OutputState }

func (LocalBowerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalBowerRepository)(nil)).Elem()
}

func (o LocalBowerRepositoryOutput) ToLocalBowerRepositoryOutput() LocalBowerRepositoryOutput {
	return o
}

func (o LocalBowerRepositoryOutput) ToLocalBowerRepositoryOutputWithContext(ctx context.Context) LocalBowerRepositoryOutput {
	return o
}

type LocalBowerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalBowerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalBowerRepository)(nil)).Elem()
}

func (o LocalBowerRepositoryArrayOutput) ToLocalBowerRepositoryArrayOutput() LocalBowerRepositoryArrayOutput {
	return o
}

func (o LocalBowerRepositoryArrayOutput) ToLocalBowerRepositoryArrayOutputWithContext(ctx context.Context) LocalBowerRepositoryArrayOutput {
	return o
}

func (o LocalBowerRepositoryArrayOutput) Index(i pulumi.IntInput) LocalBowerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalBowerRepository {
		return vs[0].([]*LocalBowerRepository)[vs[1].(int)]
	}).(LocalBowerRepositoryOutput)
}

type LocalBowerRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalBowerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalBowerRepository)(nil)).Elem()
}

func (o LocalBowerRepositoryMapOutput) ToLocalBowerRepositoryMapOutput() LocalBowerRepositoryMapOutput {
	return o
}

func (o LocalBowerRepositoryMapOutput) ToLocalBowerRepositoryMapOutputWithContext(ctx context.Context) LocalBowerRepositoryMapOutput {
	return o
}

func (o LocalBowerRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalBowerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalBowerRepository {
		return vs[0].(map[string]*LocalBowerRepository)[vs[1].(string)]
	}).(LocalBowerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBowerRepositoryInput)(nil)).Elem(), &LocalBowerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBowerRepositoryArrayInput)(nil)).Elem(), LocalBowerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalBowerRepositoryMapInput)(nil)).Elem(), LocalBowerRepositoryMap{})
	pulumi.RegisterOutputType(LocalBowerRepositoryOutput{})
	pulumi.RegisterOutputType(LocalBowerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalBowerRepositoryMapOutput{})
}
