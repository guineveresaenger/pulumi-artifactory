// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated SBT Repository Resource
//
// Creates a federated SBT repository
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
// 		_, err := artifactory.NewFederatedSbtRepository(ctx, "terraform-federated-test-sbt-repo", &artifactory.FederatedSbtRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-sbt-repo"),
// 			Members: FederatedSbtRepositoryMemberArray{
// 				&FederatedSbtRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-sbt-repo"),
// 				},
// 				&FederatedSbtRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-sbt-repo-2"),
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
type FederatedSbtRepository struct {
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
	Members     FederatedSbtRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedSbtRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedSbtRepository(ctx *pulumi.Context,
	name string, args *FederatedSbtRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedSbtRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedSbtRepository
	err := ctx.RegisterResource("artifactory:index/federatedSbtRepository:FederatedSbtRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedSbtRepository gets an existing FederatedSbtRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedSbtRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedSbtRepositoryState, opts ...pulumi.ResourceOption) (*FederatedSbtRepository, error) {
	var resource FederatedSbtRepository
	err := ctx.ReadResource("artifactory:index/federatedSbtRepository:FederatedSbtRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedSbtRepository resources.
type federatedSbtRepositoryState struct {
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
	Members     []FederatedSbtRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

type FederatedSbtRepositoryState struct {
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
	Members     FederatedSbtRepositoryMemberArrayInput
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
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedSbtRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedSbtRepositoryState)(nil)).Elem()
}

type federatedSbtRepositoryArgs struct {
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
	Members []FederatedSbtRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedSbtRepository resource.
type FederatedSbtRepositoryArgs struct {
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
	Members FederatedSbtRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedSbtRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedSbtRepositoryArgs)(nil)).Elem()
}

type FederatedSbtRepositoryInput interface {
	pulumi.Input

	ToFederatedSbtRepositoryOutput() FederatedSbtRepositoryOutput
	ToFederatedSbtRepositoryOutputWithContext(ctx context.Context) FederatedSbtRepositoryOutput
}

func (*FederatedSbtRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedSbtRepository)(nil)).Elem()
}

func (i *FederatedSbtRepository) ToFederatedSbtRepositoryOutput() FederatedSbtRepositoryOutput {
	return i.ToFederatedSbtRepositoryOutputWithContext(context.Background())
}

func (i *FederatedSbtRepository) ToFederatedSbtRepositoryOutputWithContext(ctx context.Context) FederatedSbtRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSbtRepositoryOutput)
}

// FederatedSbtRepositoryArrayInput is an input type that accepts FederatedSbtRepositoryArray and FederatedSbtRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedSbtRepositoryArrayInput` via:
//
//          FederatedSbtRepositoryArray{ FederatedSbtRepositoryArgs{...} }
type FederatedSbtRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedSbtRepositoryArrayOutput() FederatedSbtRepositoryArrayOutput
	ToFederatedSbtRepositoryArrayOutputWithContext(context.Context) FederatedSbtRepositoryArrayOutput
}

type FederatedSbtRepositoryArray []FederatedSbtRepositoryInput

func (FederatedSbtRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedSbtRepository)(nil)).Elem()
}

func (i FederatedSbtRepositoryArray) ToFederatedSbtRepositoryArrayOutput() FederatedSbtRepositoryArrayOutput {
	return i.ToFederatedSbtRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedSbtRepositoryArray) ToFederatedSbtRepositoryArrayOutputWithContext(ctx context.Context) FederatedSbtRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSbtRepositoryArrayOutput)
}

// FederatedSbtRepositoryMapInput is an input type that accepts FederatedSbtRepositoryMap and FederatedSbtRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedSbtRepositoryMapInput` via:
//
//          FederatedSbtRepositoryMap{ "key": FederatedSbtRepositoryArgs{...} }
type FederatedSbtRepositoryMapInput interface {
	pulumi.Input

	ToFederatedSbtRepositoryMapOutput() FederatedSbtRepositoryMapOutput
	ToFederatedSbtRepositoryMapOutputWithContext(context.Context) FederatedSbtRepositoryMapOutput
}

type FederatedSbtRepositoryMap map[string]FederatedSbtRepositoryInput

func (FederatedSbtRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedSbtRepository)(nil)).Elem()
}

func (i FederatedSbtRepositoryMap) ToFederatedSbtRepositoryMapOutput() FederatedSbtRepositoryMapOutput {
	return i.ToFederatedSbtRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedSbtRepositoryMap) ToFederatedSbtRepositoryMapOutputWithContext(ctx context.Context) FederatedSbtRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedSbtRepositoryMapOutput)
}

type FederatedSbtRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedSbtRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedSbtRepository)(nil)).Elem()
}

func (o FederatedSbtRepositoryOutput) ToFederatedSbtRepositoryOutput() FederatedSbtRepositoryOutput {
	return o
}

func (o FederatedSbtRepositoryOutput) ToFederatedSbtRepositoryOutputWithContext(ctx context.Context) FederatedSbtRepositoryOutput {
	return o
}

type FederatedSbtRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedSbtRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedSbtRepository)(nil)).Elem()
}

func (o FederatedSbtRepositoryArrayOutput) ToFederatedSbtRepositoryArrayOutput() FederatedSbtRepositoryArrayOutput {
	return o
}

func (o FederatedSbtRepositoryArrayOutput) ToFederatedSbtRepositoryArrayOutputWithContext(ctx context.Context) FederatedSbtRepositoryArrayOutput {
	return o
}

func (o FederatedSbtRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedSbtRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedSbtRepository {
		return vs[0].([]*FederatedSbtRepository)[vs[1].(int)]
	}).(FederatedSbtRepositoryOutput)
}

type FederatedSbtRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedSbtRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedSbtRepository)(nil)).Elem()
}

func (o FederatedSbtRepositoryMapOutput) ToFederatedSbtRepositoryMapOutput() FederatedSbtRepositoryMapOutput {
	return o
}

func (o FederatedSbtRepositoryMapOutput) ToFederatedSbtRepositoryMapOutputWithContext(ctx context.Context) FederatedSbtRepositoryMapOutput {
	return o
}

func (o FederatedSbtRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedSbtRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedSbtRepository {
		return vs[0].(map[string]*FederatedSbtRepository)[vs[1].(string)]
	}).(FederatedSbtRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSbtRepositoryInput)(nil)).Elem(), &FederatedSbtRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSbtRepositoryArrayInput)(nil)).Elem(), FederatedSbtRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedSbtRepositoryMapInput)(nil)).Elem(), FederatedSbtRepositoryMap{})
	pulumi.RegisterOutputType(FederatedSbtRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedSbtRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedSbtRepositoryMapOutput{})
}
