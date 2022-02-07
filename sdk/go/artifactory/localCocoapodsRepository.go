// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Cocoapods Repository Resource
//
// Creates a local cocoapods repository.
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
// 		_, err := artifactory.NewLocalCocoapodsRepository(ctx, "terraform_local_test_cocoapods_repo", &artifactory.LocalCocoapodsRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-cocoapods-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalCocoapodsRepository struct {
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
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewLocalCocoapodsRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCocoapodsRepository(ctx *pulumi.Context,
	name string, args *LocalCocoapodsRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCocoapodsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalCocoapodsRepository
	err := ctx.RegisterResource("artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCocoapodsRepository gets an existing LocalCocoapodsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCocoapodsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCocoapodsRepositoryState, opts ...pulumi.ResourceOption) (*LocalCocoapodsRepository, error) {
	var resource LocalCocoapodsRepository
	err := ctx.ReadResource("artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCocoapodsRepository resources.
type localCocoapodsRepositoryState struct {
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
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type LocalCocoapodsRepositoryState struct {
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
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalCocoapodsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCocoapodsRepositoryState)(nil)).Elem()
}

type localCocoapodsRepositoryArgs struct {
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
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalCocoapodsRepository resource.
type LocalCocoapodsRepositoryArgs struct {
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
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalCocoapodsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCocoapodsRepositoryArgs)(nil)).Elem()
}

type LocalCocoapodsRepositoryInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput
	ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput
}

func (*LocalCocoapodsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCocoapodsRepository)(nil))
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput {
	return i.ToLocalCocoapodsRepositoryOutputWithContext(context.Background())
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryOutput)
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryPtrOutput() LocalCocoapodsRepositoryPtrOutput {
	return i.ToLocalCocoapodsRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryPtrOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryPtrOutput)
}

type LocalCocoapodsRepositoryPtrInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryPtrOutput() LocalCocoapodsRepositoryPtrOutput
	ToLocalCocoapodsRepositoryPtrOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryPtrOutput
}

type localCocoapodsRepositoryPtrType LocalCocoapodsRepositoryArgs

func (*localCocoapodsRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCocoapodsRepository)(nil))
}

func (i *localCocoapodsRepositoryPtrType) ToLocalCocoapodsRepositoryPtrOutput() LocalCocoapodsRepositoryPtrOutput {
	return i.ToLocalCocoapodsRepositoryPtrOutputWithContext(context.Background())
}

func (i *localCocoapodsRepositoryPtrType) ToLocalCocoapodsRepositoryPtrOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryPtrOutput)
}

// LocalCocoapodsRepositoryArrayInput is an input type that accepts LocalCocoapodsRepositoryArray and LocalCocoapodsRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCocoapodsRepositoryArrayInput` via:
//
//          LocalCocoapodsRepositoryArray{ LocalCocoapodsRepositoryArgs{...} }
type LocalCocoapodsRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput
	ToLocalCocoapodsRepositoryArrayOutputWithContext(context.Context) LocalCocoapodsRepositoryArrayOutput
}

type LocalCocoapodsRepositoryArray []LocalCocoapodsRepositoryInput

func (LocalCocoapodsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCocoapodsRepository)(nil)).Elem()
}

func (i LocalCocoapodsRepositoryArray) ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput {
	return i.ToLocalCocoapodsRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCocoapodsRepositoryArray) ToLocalCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryArrayOutput)
}

// LocalCocoapodsRepositoryMapInput is an input type that accepts LocalCocoapodsRepositoryMap and LocalCocoapodsRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCocoapodsRepositoryMapInput` via:
//
//          LocalCocoapodsRepositoryMap{ "key": LocalCocoapodsRepositoryArgs{...} }
type LocalCocoapodsRepositoryMapInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput
	ToLocalCocoapodsRepositoryMapOutputWithContext(context.Context) LocalCocoapodsRepositoryMapOutput
}

type LocalCocoapodsRepositoryMap map[string]LocalCocoapodsRepositoryInput

func (LocalCocoapodsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCocoapodsRepository)(nil)).Elem()
}

func (i LocalCocoapodsRepositoryMap) ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput {
	return i.ToLocalCocoapodsRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCocoapodsRepositoryMap) ToLocalCocoapodsRepositoryMapOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryMapOutput)
}

type LocalCocoapodsRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCocoapodsRepository)(nil))
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput {
	return o
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput {
	return o
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryPtrOutput() LocalCocoapodsRepositoryPtrOutput {
	return o.ToLocalCocoapodsRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryPtrOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalCocoapodsRepository) *LocalCocoapodsRepository {
		return &v
	}).(LocalCocoapodsRepositoryPtrOutput)
}

type LocalCocoapodsRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCocoapodsRepository)(nil))
}

func (o LocalCocoapodsRepositoryPtrOutput) ToLocalCocoapodsRepositoryPtrOutput() LocalCocoapodsRepositoryPtrOutput {
	return o
}

func (o LocalCocoapodsRepositoryPtrOutput) ToLocalCocoapodsRepositoryPtrOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryPtrOutput {
	return o
}

func (o LocalCocoapodsRepositoryPtrOutput) Elem() LocalCocoapodsRepositoryOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) LocalCocoapodsRepository {
		if v != nil {
			return *v
		}
		var ret LocalCocoapodsRepository
		return ret
	}).(LocalCocoapodsRepositoryOutput)
}

type LocalCocoapodsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalCocoapodsRepository)(nil))
}

func (o LocalCocoapodsRepositoryArrayOutput) ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput {
	return o
}

func (o LocalCocoapodsRepositoryArrayOutput) ToLocalCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryArrayOutput {
	return o
}

func (o LocalCocoapodsRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCocoapodsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalCocoapodsRepository {
		return vs[0].([]LocalCocoapodsRepository)[vs[1].(int)]
	}).(LocalCocoapodsRepositoryOutput)
}

type LocalCocoapodsRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalCocoapodsRepository)(nil))
}

func (o LocalCocoapodsRepositoryMapOutput) ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput {
	return o
}

func (o LocalCocoapodsRepositoryMapOutput) ToLocalCocoapodsRepositoryMapOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryMapOutput {
	return o
}

func (o LocalCocoapodsRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCocoapodsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalCocoapodsRepository {
		return vs[0].(map[string]LocalCocoapodsRepository)[vs[1].(string)]
	}).(LocalCocoapodsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryInput)(nil)).Elem(), &LocalCocoapodsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryPtrInput)(nil)).Elem(), &LocalCocoapodsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryArrayInput)(nil)).Elem(), LocalCocoapodsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryMapInput)(nil)).Elem(), LocalCocoapodsRepositoryMap{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryMapOutput{})
}
