// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Maven Repository Resource
//
// Creates a federated Maven repository
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
// 		_, err := artifactory.NewFederatedMavenRepository(ctx, "terraform-federated-test-maven-repo", &artifactory.FederatedMavenRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-maven-repo"),
// 			Members: FederatedMavenRepositoryMemberArray{
// 				&FederatedMavenRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-maven-repo"),
// 				},
// 				&FederatedMavenRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-maven-repo-2"),
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
type FederatedMavenRepository struct {
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
	Members     FederatedMavenRepositoryMemberArrayOutput `pulumi:"members"`
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

// NewFederatedMavenRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedMavenRepository(ctx *pulumi.Context,
	name string, args *FederatedMavenRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedMavenRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedMavenRepository
	err := ctx.RegisterResource("artifactory:index/federatedMavenRepository:FederatedMavenRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedMavenRepository gets an existing FederatedMavenRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedMavenRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedMavenRepositoryState, opts ...pulumi.ResourceOption) (*FederatedMavenRepository, error) {
	var resource FederatedMavenRepository
	err := ctx.ReadResource("artifactory:index/federatedMavenRepository:FederatedMavenRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedMavenRepository resources.
type federatedMavenRepositoryState struct {
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
	Members     []FederatedMavenRepositoryMember `pulumi:"members"`
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

type FederatedMavenRepositoryState struct {
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
	Members     FederatedMavenRepositoryMemberArrayInput
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

func (FederatedMavenRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedMavenRepositoryState)(nil)).Elem()
}

type federatedMavenRepositoryArgs struct {
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
	Members []FederatedMavenRepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedMavenRepository resource.
type FederatedMavenRepositoryArgs struct {
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
	Members FederatedMavenRepositoryMemberArrayInput
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

func (FederatedMavenRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedMavenRepositoryArgs)(nil)).Elem()
}

type FederatedMavenRepositoryInput interface {
	pulumi.Input

	ToFederatedMavenRepositoryOutput() FederatedMavenRepositoryOutput
	ToFederatedMavenRepositoryOutputWithContext(ctx context.Context) FederatedMavenRepositoryOutput
}

func (*FederatedMavenRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedMavenRepository)(nil)).Elem()
}

func (i *FederatedMavenRepository) ToFederatedMavenRepositoryOutput() FederatedMavenRepositoryOutput {
	return i.ToFederatedMavenRepositoryOutputWithContext(context.Background())
}

func (i *FederatedMavenRepository) ToFederatedMavenRepositoryOutputWithContext(ctx context.Context) FederatedMavenRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedMavenRepositoryOutput)
}

// FederatedMavenRepositoryArrayInput is an input type that accepts FederatedMavenRepositoryArray and FederatedMavenRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedMavenRepositoryArrayInput` via:
//
//          FederatedMavenRepositoryArray{ FederatedMavenRepositoryArgs{...} }
type FederatedMavenRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedMavenRepositoryArrayOutput() FederatedMavenRepositoryArrayOutput
	ToFederatedMavenRepositoryArrayOutputWithContext(context.Context) FederatedMavenRepositoryArrayOutput
}

type FederatedMavenRepositoryArray []FederatedMavenRepositoryInput

func (FederatedMavenRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedMavenRepository)(nil)).Elem()
}

func (i FederatedMavenRepositoryArray) ToFederatedMavenRepositoryArrayOutput() FederatedMavenRepositoryArrayOutput {
	return i.ToFederatedMavenRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedMavenRepositoryArray) ToFederatedMavenRepositoryArrayOutputWithContext(ctx context.Context) FederatedMavenRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedMavenRepositoryArrayOutput)
}

// FederatedMavenRepositoryMapInput is an input type that accepts FederatedMavenRepositoryMap and FederatedMavenRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedMavenRepositoryMapInput` via:
//
//          FederatedMavenRepositoryMap{ "key": FederatedMavenRepositoryArgs{...} }
type FederatedMavenRepositoryMapInput interface {
	pulumi.Input

	ToFederatedMavenRepositoryMapOutput() FederatedMavenRepositoryMapOutput
	ToFederatedMavenRepositoryMapOutputWithContext(context.Context) FederatedMavenRepositoryMapOutput
}

type FederatedMavenRepositoryMap map[string]FederatedMavenRepositoryInput

func (FederatedMavenRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedMavenRepository)(nil)).Elem()
}

func (i FederatedMavenRepositoryMap) ToFederatedMavenRepositoryMapOutput() FederatedMavenRepositoryMapOutput {
	return i.ToFederatedMavenRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedMavenRepositoryMap) ToFederatedMavenRepositoryMapOutputWithContext(ctx context.Context) FederatedMavenRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedMavenRepositoryMapOutput)
}

type FederatedMavenRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedMavenRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedMavenRepository)(nil)).Elem()
}

func (o FederatedMavenRepositoryOutput) ToFederatedMavenRepositoryOutput() FederatedMavenRepositoryOutput {
	return o
}

func (o FederatedMavenRepositoryOutput) ToFederatedMavenRepositoryOutputWithContext(ctx context.Context) FederatedMavenRepositoryOutput {
	return o
}

type FederatedMavenRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedMavenRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedMavenRepository)(nil)).Elem()
}

func (o FederatedMavenRepositoryArrayOutput) ToFederatedMavenRepositoryArrayOutput() FederatedMavenRepositoryArrayOutput {
	return o
}

func (o FederatedMavenRepositoryArrayOutput) ToFederatedMavenRepositoryArrayOutputWithContext(ctx context.Context) FederatedMavenRepositoryArrayOutput {
	return o
}

func (o FederatedMavenRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedMavenRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedMavenRepository {
		return vs[0].([]*FederatedMavenRepository)[vs[1].(int)]
	}).(FederatedMavenRepositoryOutput)
}

type FederatedMavenRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedMavenRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedMavenRepository)(nil)).Elem()
}

func (o FederatedMavenRepositoryMapOutput) ToFederatedMavenRepositoryMapOutput() FederatedMavenRepositoryMapOutput {
	return o
}

func (o FederatedMavenRepositoryMapOutput) ToFederatedMavenRepositoryMapOutputWithContext(ctx context.Context) FederatedMavenRepositoryMapOutput {
	return o
}

func (o FederatedMavenRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedMavenRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedMavenRepository {
		return vs[0].(map[string]*FederatedMavenRepository)[vs[1].(string)]
	}).(FederatedMavenRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedMavenRepositoryInput)(nil)).Elem(), &FederatedMavenRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedMavenRepositoryArrayInput)(nil)).Elem(), FederatedMavenRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedMavenRepositoryMapInput)(nil)).Elem(), FederatedMavenRepositoryMap{})
	pulumi.RegisterOutputType(FederatedMavenRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedMavenRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedMavenRepositoryMapOutput{})
}
