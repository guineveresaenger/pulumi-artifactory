// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Bower Repository Resource
//
// Creates a federated Bower repository
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
// 		_, err := artifactory.NewFederatedBowerRepository(ctx, "terraform-federated-test-bower-repo", &artifactory.FederatedBowerRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-bower-repo"),
// 			Members: FederatedBowerRepositoryMemberArray{
// 				&FederatedBowerRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-bower-repo"),
// 				},
// 				&FederatedBowerRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-bower-repo-2"),
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
type FederatedBowerRepository struct {
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
	Members     FederatedBowerRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                    `pulumi:"notes"`
	PackageType pulumi.StringOutput                       `pulumi:"packageType"`
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

// NewFederatedBowerRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedBowerRepository(ctx *pulumi.Context,
	name string, args *FederatedBowerRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedBowerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedBowerRepository
	err := ctx.RegisterResource("artifactory:index/federatedBowerRepository:FederatedBowerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedBowerRepository gets an existing FederatedBowerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedBowerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedBowerRepositoryState, opts ...pulumi.ResourceOption) (*FederatedBowerRepository, error) {
	var resource FederatedBowerRepository
	err := ctx.ReadResource("artifactory:index/federatedBowerRepository:FederatedBowerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedBowerRepository resources.
type federatedBowerRepositoryState struct {
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
	Members     []FederatedBowerRepositoryMember `pulumi:"members"`
	Notes       *string                          `pulumi:"notes"`
	PackageType *string                          `pulumi:"packageType"`
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

type FederatedBowerRepositoryState struct {
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
	Members     FederatedBowerRepositoryMemberArrayInput
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

func (FederatedBowerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedBowerRepositoryState)(nil)).Elem()
}

type federatedBowerRepositoryArgs struct {
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
	Members []FederatedBowerRepositoryMember `pulumi:"members"`
	Notes   *string                          `pulumi:"notes"`
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

// The set of arguments for constructing a FederatedBowerRepository resource.
type FederatedBowerRepositoryArgs struct {
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
	Members FederatedBowerRepositoryMemberArrayInput
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
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedBowerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedBowerRepositoryArgs)(nil)).Elem()
}

type FederatedBowerRepositoryInput interface {
	pulumi.Input

	ToFederatedBowerRepositoryOutput() FederatedBowerRepositoryOutput
	ToFederatedBowerRepositoryOutputWithContext(ctx context.Context) FederatedBowerRepositoryOutput
}

func (*FederatedBowerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedBowerRepository)(nil)).Elem()
}

func (i *FederatedBowerRepository) ToFederatedBowerRepositoryOutput() FederatedBowerRepositoryOutput {
	return i.ToFederatedBowerRepositoryOutputWithContext(context.Background())
}

func (i *FederatedBowerRepository) ToFederatedBowerRepositoryOutputWithContext(ctx context.Context) FederatedBowerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedBowerRepositoryOutput)
}

// FederatedBowerRepositoryArrayInput is an input type that accepts FederatedBowerRepositoryArray and FederatedBowerRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedBowerRepositoryArrayInput` via:
//
//          FederatedBowerRepositoryArray{ FederatedBowerRepositoryArgs{...} }
type FederatedBowerRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedBowerRepositoryArrayOutput() FederatedBowerRepositoryArrayOutput
	ToFederatedBowerRepositoryArrayOutputWithContext(context.Context) FederatedBowerRepositoryArrayOutput
}

type FederatedBowerRepositoryArray []FederatedBowerRepositoryInput

func (FederatedBowerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedBowerRepository)(nil)).Elem()
}

func (i FederatedBowerRepositoryArray) ToFederatedBowerRepositoryArrayOutput() FederatedBowerRepositoryArrayOutput {
	return i.ToFederatedBowerRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedBowerRepositoryArray) ToFederatedBowerRepositoryArrayOutputWithContext(ctx context.Context) FederatedBowerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedBowerRepositoryArrayOutput)
}

// FederatedBowerRepositoryMapInput is an input type that accepts FederatedBowerRepositoryMap and FederatedBowerRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedBowerRepositoryMapInput` via:
//
//          FederatedBowerRepositoryMap{ "key": FederatedBowerRepositoryArgs{...} }
type FederatedBowerRepositoryMapInput interface {
	pulumi.Input

	ToFederatedBowerRepositoryMapOutput() FederatedBowerRepositoryMapOutput
	ToFederatedBowerRepositoryMapOutputWithContext(context.Context) FederatedBowerRepositoryMapOutput
}

type FederatedBowerRepositoryMap map[string]FederatedBowerRepositoryInput

func (FederatedBowerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedBowerRepository)(nil)).Elem()
}

func (i FederatedBowerRepositoryMap) ToFederatedBowerRepositoryMapOutput() FederatedBowerRepositoryMapOutput {
	return i.ToFederatedBowerRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedBowerRepositoryMap) ToFederatedBowerRepositoryMapOutputWithContext(ctx context.Context) FederatedBowerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedBowerRepositoryMapOutput)
}

type FederatedBowerRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedBowerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedBowerRepository)(nil)).Elem()
}

func (o FederatedBowerRepositoryOutput) ToFederatedBowerRepositoryOutput() FederatedBowerRepositoryOutput {
	return o
}

func (o FederatedBowerRepositoryOutput) ToFederatedBowerRepositoryOutputWithContext(ctx context.Context) FederatedBowerRepositoryOutput {
	return o
}

type FederatedBowerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedBowerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedBowerRepository)(nil)).Elem()
}

func (o FederatedBowerRepositoryArrayOutput) ToFederatedBowerRepositoryArrayOutput() FederatedBowerRepositoryArrayOutput {
	return o
}

func (o FederatedBowerRepositoryArrayOutput) ToFederatedBowerRepositoryArrayOutputWithContext(ctx context.Context) FederatedBowerRepositoryArrayOutput {
	return o
}

func (o FederatedBowerRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedBowerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedBowerRepository {
		return vs[0].([]*FederatedBowerRepository)[vs[1].(int)]
	}).(FederatedBowerRepositoryOutput)
}

type FederatedBowerRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedBowerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedBowerRepository)(nil)).Elem()
}

func (o FederatedBowerRepositoryMapOutput) ToFederatedBowerRepositoryMapOutput() FederatedBowerRepositoryMapOutput {
	return o
}

func (o FederatedBowerRepositoryMapOutput) ToFederatedBowerRepositoryMapOutputWithContext(ctx context.Context) FederatedBowerRepositoryMapOutput {
	return o
}

func (o FederatedBowerRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedBowerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedBowerRepository {
		return vs[0].(map[string]*FederatedBowerRepository)[vs[1].(string)]
	}).(FederatedBowerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedBowerRepositoryInput)(nil)).Elem(), &FederatedBowerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedBowerRepositoryArrayInput)(nil)).Elem(), FederatedBowerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedBowerRepositoryMapInput)(nil)).Elem(), FederatedBowerRepositoryMap{})
	pulumi.RegisterOutputType(FederatedBowerRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedBowerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedBowerRepositoryMapOutput{})
}
