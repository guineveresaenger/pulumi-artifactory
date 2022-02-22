// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Debian Repository Resource
//
// Creates a local Debian repository and allows for the creation of a GPG key
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewKeypair(ctx, "some-keypairGPG1", &artifactory.KeypairArgs{
// 			PairName:   pulumi.String(fmt.Sprintf("%v%v", "some-keypair", random_id.Randid.Id)),
// 			PairType:   pulumi.String("GPG"),
// 			Alias:      pulumi.String("foo-alias1"),
// 			PrivateKey: readFileOrPanic("samples/gpg.priv"),
// 			PublicKey:  readFileOrPanic("samples/gpg.pub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewKeypair(ctx, "some-keypairGPG2", &artifactory.KeypairArgs{
// 			PairName:   pulumi.String(fmt.Sprintf("%v%v", "some-keypair4", random_id.Randid.Id)),
// 			PairType:   pulumi.String("GPG"),
// 			Alias:      pulumi.String("foo-alias2"),
// 			PrivateKey: readFileOrPanic("samples/gpg.priv"),
// 			PublicKey:  readFileOrPanic("samples/gpg.pub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewDebianRepository(ctx, "my-debian-repo", &artifactory.DebianRepositoryArgs{
// 			Key:                 pulumi.String("my-debian-repo"),
// 			PrimaryKeypairRef:   some_keypairGPG1.PairName,
// 			SecondaryKeypairRef: some_keypairGPG2.PairName,
// 			IndexCompressionFormats: pulumi.StringArray{
// 				pulumi.String("bz2"),
// 				pulumi.String("lzma"),
// 				pulumi.String("xz"),
// 			},
// 			TrivialLayout: pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			some_keypairGPG1,
// 			some_keypairGPG2,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DebianRepository struct {
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
	// - If you're creating this repo, then maybe you know?
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// - The RSA key to be used to sign packages
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// - Not really clear what this does
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// - Apparently this is a deprecated repo layout
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrOutput `pulumi:"trivialLayout"`
	XrayIndex     pulumi.BoolOutput    `pulumi:"xrayIndex"`
}

// NewDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewDebianRepository(ctx *pulumi.Context,
	name string, args *DebianRepositoryArgs, opts ...pulumi.ResourceOption) (*DebianRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource DebianRepository
	err := ctx.RegisterResource("artifactory:index/debianRepository:DebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDebianRepository gets an existing DebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DebianRepositoryState, opts ...pulumi.ResourceOption) (*DebianRepository, error) {
	var resource DebianRepository
	err := ctx.ReadResource("artifactory:index/debianRepository:DebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DebianRepository resources.
type debianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - If you're creating this repo, then maybe you know?
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// - The RSA key to be used to sign packages
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// - Not really clear what this does
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// - Apparently this is a deprecated repo layout
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

type DebianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - If you're creating this repo, then maybe you know?
	IndexCompressionFormats pulumi.StringArrayInput
	// - the identity key of the repo
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// - The RSA key to be used to sign packages
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// - Not really clear what this does
	SecondaryKeypairRef pulumi.StringPtrInput
	// - Apparently this is a deprecated repo layout
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (DebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*debianRepositoryState)(nil)).Elem()
}

type debianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - If you're creating this repo, then maybe you know?
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// - The RSA key to be used to sign packages
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// - Not really clear what this does
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// - Apparently this is a deprecated repo layout
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a DebianRepository resource.
type DebianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - If you're creating this repo, then maybe you know?
	IndexCompressionFormats pulumi.StringArrayInput
	// - the identity key of the repo
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// - The RSA key to be used to sign packages
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// - Not really clear what this does
	SecondaryKeypairRef pulumi.StringPtrInput
	// - Apparently this is a deprecated repo layout
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (DebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*debianRepositoryArgs)(nil)).Elem()
}

type DebianRepositoryInput interface {
	pulumi.Input

	ToDebianRepositoryOutput() DebianRepositoryOutput
	ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput
}

func (*DebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**DebianRepository)(nil)).Elem()
}

func (i *DebianRepository) ToDebianRepositoryOutput() DebianRepositoryOutput {
	return i.ToDebianRepositoryOutputWithContext(context.Background())
}

func (i *DebianRepository) ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryOutput)
}

// DebianRepositoryArrayInput is an input type that accepts DebianRepositoryArray and DebianRepositoryArrayOutput values.
// You can construct a concrete instance of `DebianRepositoryArrayInput` via:
//
//          DebianRepositoryArray{ DebianRepositoryArgs{...} }
type DebianRepositoryArrayInput interface {
	pulumi.Input

	ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput
	ToDebianRepositoryArrayOutputWithContext(context.Context) DebianRepositoryArrayOutput
}

type DebianRepositoryArray []DebianRepositoryInput

func (DebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DebianRepository)(nil)).Elem()
}

func (i DebianRepositoryArray) ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput {
	return i.ToDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i DebianRepositoryArray) ToDebianRepositoryArrayOutputWithContext(ctx context.Context) DebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryArrayOutput)
}

// DebianRepositoryMapInput is an input type that accepts DebianRepositoryMap and DebianRepositoryMapOutput values.
// You can construct a concrete instance of `DebianRepositoryMapInput` via:
//
//          DebianRepositoryMap{ "key": DebianRepositoryArgs{...} }
type DebianRepositoryMapInput interface {
	pulumi.Input

	ToDebianRepositoryMapOutput() DebianRepositoryMapOutput
	ToDebianRepositoryMapOutputWithContext(context.Context) DebianRepositoryMapOutput
}

type DebianRepositoryMap map[string]DebianRepositoryInput

func (DebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DebianRepository)(nil)).Elem()
}

func (i DebianRepositoryMap) ToDebianRepositoryMapOutput() DebianRepositoryMapOutput {
	return i.ToDebianRepositoryMapOutputWithContext(context.Background())
}

func (i DebianRepositoryMap) ToDebianRepositoryMapOutputWithContext(ctx context.Context) DebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryMapOutput)
}

type DebianRepositoryOutput struct{ *pulumi.OutputState }

func (DebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryOutput) ToDebianRepositoryOutput() DebianRepositoryOutput {
	return o
}

func (o DebianRepositoryOutput) ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput {
	return o
}

type DebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (DebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryArrayOutput) ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput {
	return o
}

func (o DebianRepositoryArrayOutput) ToDebianRepositoryArrayOutputWithContext(ctx context.Context) DebianRepositoryArrayOutput {
	return o
}

func (o DebianRepositoryArrayOutput) Index(i pulumi.IntInput) DebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DebianRepository {
		return vs[0].([]*DebianRepository)[vs[1].(int)]
	}).(DebianRepositoryOutput)
}

type DebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (DebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryMapOutput) ToDebianRepositoryMapOutput() DebianRepositoryMapOutput {
	return o
}

func (o DebianRepositoryMapOutput) ToDebianRepositoryMapOutputWithContext(ctx context.Context) DebianRepositoryMapOutput {
	return o
}

func (o DebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) DebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DebianRepository {
		return vs[0].(map[string]*DebianRepository)[vs[1].(string)]
	}).(DebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryInput)(nil)).Elem(), &DebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryArrayInput)(nil)).Elem(), DebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryMapInput)(nil)).Elem(), DebianRepositoryMap{})
	pulumi.RegisterOutputType(DebianRepositoryOutput{})
	pulumi.RegisterOutputType(DebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(DebianRepositoryMapOutput{})
}
