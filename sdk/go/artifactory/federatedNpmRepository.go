// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Npm Repository Resource
//
// Creates a federated Npm repository
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
// 		_, err := artifactory.NewFederatedNpmRepository(ctx, "terraform_federated_test_npm_repo", &artifactory.FederatedNpmRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-npm-repo"),
// 			Members: FederatedNpmRepositoryMemberArray{
// 				&FederatedNpmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-npm-repo"),
// 				},
// 				&FederatedNpmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-npm-repo-2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FederatedNpmRepository struct {
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
	Key pulumi.StringOutput `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedNpmRepository(ctx *pulumi.Context,
	name string, args *FederatedNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedNpmRepository
	err := ctx.RegisterResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedNpmRepository gets an existing FederatedNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedNpmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	var resource FederatedNpmRepository
	err := ctx.ReadResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedNpmRepository resources.
type federatedNpmRepositoryState struct {
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
	Key *string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedNpmRepositoryState struct {
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
	Key pulumi.StringPtrInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryState)(nil)).Elem()
}

type federatedNpmRepositoryArgs struct {
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
	Key string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedNpmRepository resource.
type FederatedNpmRepositoryArgs struct {
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
	Key pulumi.StringInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members FederatedNpmRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryArgs)(nil)).Elem()
}

type FederatedNpmRepositoryInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput
	ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput
}

func (*FederatedNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedNpmRepository)(nil))
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return i.ToFederatedNpmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryOutput)
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryPtrOutput() FederatedNpmRepositoryPtrOutput {
	return i.ToFederatedNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryPtrOutputWithContext(ctx context.Context) FederatedNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryPtrOutput)
}

type FederatedNpmRepositoryPtrInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryPtrOutput() FederatedNpmRepositoryPtrOutput
	ToFederatedNpmRepositoryPtrOutputWithContext(ctx context.Context) FederatedNpmRepositoryPtrOutput
}

type federatedNpmRepositoryPtrType FederatedNpmRepositoryArgs

func (*federatedNpmRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil))
}

func (i *federatedNpmRepositoryPtrType) ToFederatedNpmRepositoryPtrOutput() FederatedNpmRepositoryPtrOutput {
	return i.ToFederatedNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *federatedNpmRepositoryPtrType) ToFederatedNpmRepositoryPtrOutputWithContext(ctx context.Context) FederatedNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryPtrOutput)
}

// FederatedNpmRepositoryArrayInput is an input type that accepts FederatedNpmRepositoryArray and FederatedNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryArrayInput` via:
//
//          FederatedNpmRepositoryArray{ FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput
	ToFederatedNpmRepositoryArrayOutputWithContext(context.Context) FederatedNpmRepositoryArrayOutput
}

type FederatedNpmRepositoryArray []FederatedNpmRepositoryInput

func (FederatedNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return i.ToFederatedNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryArrayOutput)
}

// FederatedNpmRepositoryMapInput is an input type that accepts FederatedNpmRepositoryMap and FederatedNpmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryMapInput` via:
//
//          FederatedNpmRepositoryMap{ "key": FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput
	ToFederatedNpmRepositoryMapOutputWithContext(context.Context) FederatedNpmRepositoryMapOutput
}

type FederatedNpmRepositoryMap map[string]FederatedNpmRepositoryInput

func (FederatedNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return i.ToFederatedNpmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryMapOutput)
}

type FederatedNpmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedNpmRepository)(nil))
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return o
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return o
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryPtrOutput() FederatedNpmRepositoryPtrOutput {
	return o.ToFederatedNpmRepositoryPtrOutputWithContext(context.Background())
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryPtrOutputWithContext(ctx context.Context) FederatedNpmRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FederatedNpmRepository) *FederatedNpmRepository {
		return &v
	}).(FederatedNpmRepositoryPtrOutput)
}

type FederatedNpmRepositoryPtrOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil))
}

func (o FederatedNpmRepositoryPtrOutput) ToFederatedNpmRepositoryPtrOutput() FederatedNpmRepositoryPtrOutput {
	return o
}

func (o FederatedNpmRepositoryPtrOutput) ToFederatedNpmRepositoryPtrOutputWithContext(ctx context.Context) FederatedNpmRepositoryPtrOutput {
	return o
}

func (o FederatedNpmRepositoryPtrOutput) Elem() FederatedNpmRepositoryOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) FederatedNpmRepository {
		if v != nil {
			return *v
		}
		var ret FederatedNpmRepository
		return ret
	}).(FederatedNpmRepositoryOutput)
}

type FederatedNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FederatedNpmRepository)(nil))
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FederatedNpmRepository {
		return vs[0].([]FederatedNpmRepository)[vs[1].(int)]
	}).(FederatedNpmRepositoryOutput)
}

type FederatedNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FederatedNpmRepository)(nil))
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FederatedNpmRepository {
		return vs[0].(map[string]FederatedNpmRepository)[vs[1].(string)]
	}).(FederatedNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryInput)(nil)).Elem(), &FederatedNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryPtrInput)(nil)).Elem(), &FederatedNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryArrayInput)(nil)).Elem(), FederatedNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryMapInput)(nil)).Elem(), FederatedNpmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryPtrOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryMapOutput{})
}