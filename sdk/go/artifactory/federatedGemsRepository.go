// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Gem Repository Resource
//
// Creates a federated Gem repository
type FederatedGemsRepository struct {
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
	Members     FederatedGemsRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                   `pulumi:"notes"`
	PackageType pulumi.StringOutput                      `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewFederatedGemsRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedGemsRepository(ctx *pulumi.Context,
	name string, args *FederatedGemsRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedGemsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedGemsRepository
	err := ctx.RegisterResource("artifactory:index/federatedGemsRepository:FederatedGemsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedGemsRepository gets an existing FederatedGemsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedGemsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedGemsRepositoryState, opts ...pulumi.ResourceOption) (*FederatedGemsRepository, error) {
	var resource FederatedGemsRepository
	err := ctx.ReadResource("artifactory:index/federatedGemsRepository:FederatedGemsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedGemsRepository resources.
type federatedGemsRepositoryState struct {
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
	Members     []FederatedGemsRepositoryMember `pulumi:"members"`
	Notes       *string                         `pulumi:"notes"`
	PackageType *string                         `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   *string  `pulumi:"projectKey"`
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

type FederatedGemsRepositoryState struct {
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
	Members     FederatedGemsRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrInput
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedGemsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGemsRepositoryState)(nil)).Elem()
}

type federatedGemsRepositoryArgs struct {
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
	Members []FederatedGemsRepositoryMember `pulumi:"members"`
	Notes   *string                         `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   *string  `pulumi:"projectKey"`
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedGemsRepository resource.
type FederatedGemsRepositoryArgs struct {
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
	Members FederatedGemsRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrInput
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedGemsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGemsRepositoryArgs)(nil)).Elem()
}

type FederatedGemsRepositoryInput interface {
	pulumi.Input

	ToFederatedGemsRepositoryOutput() FederatedGemsRepositoryOutput
	ToFederatedGemsRepositoryOutputWithContext(ctx context.Context) FederatedGemsRepositoryOutput
}

func (*FederatedGemsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGemsRepository)(nil)).Elem()
}

func (i *FederatedGemsRepository) ToFederatedGemsRepositoryOutput() FederatedGemsRepositoryOutput {
	return i.ToFederatedGemsRepositoryOutputWithContext(context.Background())
}

func (i *FederatedGemsRepository) ToFederatedGemsRepositoryOutputWithContext(ctx context.Context) FederatedGemsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGemsRepositoryOutput)
}

// FederatedGemsRepositoryArrayInput is an input type that accepts FederatedGemsRepositoryArray and FederatedGemsRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedGemsRepositoryArrayInput` via:
//
//          FederatedGemsRepositoryArray{ FederatedGemsRepositoryArgs{...} }
type FederatedGemsRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedGemsRepositoryArrayOutput() FederatedGemsRepositoryArrayOutput
	ToFederatedGemsRepositoryArrayOutputWithContext(context.Context) FederatedGemsRepositoryArrayOutput
}

type FederatedGemsRepositoryArray []FederatedGemsRepositoryInput

func (FederatedGemsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGemsRepository)(nil)).Elem()
}

func (i FederatedGemsRepositoryArray) ToFederatedGemsRepositoryArrayOutput() FederatedGemsRepositoryArrayOutput {
	return i.ToFederatedGemsRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedGemsRepositoryArray) ToFederatedGemsRepositoryArrayOutputWithContext(ctx context.Context) FederatedGemsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGemsRepositoryArrayOutput)
}

// FederatedGemsRepositoryMapInput is an input type that accepts FederatedGemsRepositoryMap and FederatedGemsRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedGemsRepositoryMapInput` via:
//
//          FederatedGemsRepositoryMap{ "key": FederatedGemsRepositoryArgs{...} }
type FederatedGemsRepositoryMapInput interface {
	pulumi.Input

	ToFederatedGemsRepositoryMapOutput() FederatedGemsRepositoryMapOutput
	ToFederatedGemsRepositoryMapOutputWithContext(context.Context) FederatedGemsRepositoryMapOutput
}

type FederatedGemsRepositoryMap map[string]FederatedGemsRepositoryInput

func (FederatedGemsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGemsRepository)(nil)).Elem()
}

func (i FederatedGemsRepositoryMap) ToFederatedGemsRepositoryMapOutput() FederatedGemsRepositoryMapOutput {
	return i.ToFederatedGemsRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedGemsRepositoryMap) ToFederatedGemsRepositoryMapOutputWithContext(ctx context.Context) FederatedGemsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGemsRepositoryMapOutput)
}

type FederatedGemsRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedGemsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGemsRepository)(nil)).Elem()
}

func (o FederatedGemsRepositoryOutput) ToFederatedGemsRepositoryOutput() FederatedGemsRepositoryOutput {
	return o
}

func (o FederatedGemsRepositoryOutput) ToFederatedGemsRepositoryOutputWithContext(ctx context.Context) FederatedGemsRepositoryOutput {
	return o
}

type FederatedGemsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedGemsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGemsRepository)(nil)).Elem()
}

func (o FederatedGemsRepositoryArrayOutput) ToFederatedGemsRepositoryArrayOutput() FederatedGemsRepositoryArrayOutput {
	return o
}

func (o FederatedGemsRepositoryArrayOutput) ToFederatedGemsRepositoryArrayOutputWithContext(ctx context.Context) FederatedGemsRepositoryArrayOutput {
	return o
}

func (o FederatedGemsRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedGemsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedGemsRepository {
		return vs[0].([]*FederatedGemsRepository)[vs[1].(int)]
	}).(FederatedGemsRepositoryOutput)
}

type FederatedGemsRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedGemsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGemsRepository)(nil)).Elem()
}

func (o FederatedGemsRepositoryMapOutput) ToFederatedGemsRepositoryMapOutput() FederatedGemsRepositoryMapOutput {
	return o
}

func (o FederatedGemsRepositoryMapOutput) ToFederatedGemsRepositoryMapOutputWithContext(ctx context.Context) FederatedGemsRepositoryMapOutput {
	return o
}

func (o FederatedGemsRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedGemsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedGemsRepository {
		return vs[0].(map[string]*FederatedGemsRepository)[vs[1].(string)]
	}).(FederatedGemsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGemsRepositoryInput)(nil)).Elem(), &FederatedGemsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGemsRepositoryArrayInput)(nil)).Elem(), FederatedGemsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGemsRepositoryMapInput)(nil)).Elem(), FederatedGemsRepositoryMap{})
	pulumi.RegisterOutputType(FederatedGemsRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedGemsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedGemsRepositoryMapOutput{})
}
