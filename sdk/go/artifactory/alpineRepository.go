// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Alpine Repository Resource
//
// Creates a local Alpine repository and allows for the creation of a
//
// ## Example Usage
//
// ```go
// package main
//
// import (
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
// 		_, err := artifactory.NewKeypair(ctx, "some-keypairRSA", &artifactory.KeypairArgs{
// 			PairName:   pulumi.String("some-keypair"),
// 			PairType:   pulumi.String("RSA"),
// 			Alias:      pulumi.String("foo-alias"),
// 			PrivateKey: readFileOrPanic("samples/rsa.priv"),
// 			PublicKey:  readFileOrPanic("samples/rsa.pub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewAlpineRepository(ctx, "terraform-local-test-alpine-repo-basic", &artifactory.AlpineRepositoryArgs{
// 			Key:               pulumi.String("terraform-local-test-alpine-repo-basic"),
// 			PrimaryKeypairRef: some_keypairRSA.PairName,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			some_keypairRSA,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AlpineRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled  pulumi.BoolPtrOutput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              pulumi.BoolPtrOutput     `pulumi:"blackedOut"`
	Description             pulumi.StringPtrOutput   `pulumi:"description"`
	DownloadDirect          pulumi.BoolPtrOutput     `pulumi:"downloadDirect"`
	ExcludesPattern         pulumi.StringOutput      `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringOutput      `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// - The RSA key to be used to sign alpine indecies
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
	XrayIndex     pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewAlpineRepository registers a new resource with the given unique name, arguments, and options.
func NewAlpineRepository(ctx *pulumi.Context,
	name string, args *AlpineRepositoryArgs, opts ...pulumi.ResourceOption) (*AlpineRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource AlpineRepository
	err := ctx.RegisterResource("artifactory:index/alpineRepository:AlpineRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlpineRepository gets an existing AlpineRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlpineRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlpineRepositoryState, opts ...pulumi.ResourceOption) (*AlpineRepository, error) {
	var resource AlpineRepository
	err := ctx.ReadResource("artifactory:index/alpineRepository:AlpineRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlpineRepository resources.
type alpineRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled  *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool    `pulumi:"blackedOut"`
	Description             *string  `pulumi:"description"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// - The RSA key to be used to sign alpine indecies
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
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

type AlpineRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled  pulumi.BoolPtrInput
	BlackedOut              pulumi.BoolPtrInput
	Description             pulumi.StringPtrInput
	DownloadDirect          pulumi.BoolPtrInput
	ExcludesPattern         pulumi.StringPtrInput
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// - the identity key of the repo
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// - The RSA key to be used to sign alpine indecies
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
	XrayIndex     pulumi.BoolPtrInput
}

func (AlpineRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*alpineRepositoryState)(nil)).Elem()
}

type alpineRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled  *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool    `pulumi:"blackedOut"`
	Description             *string  `pulumi:"description"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// - the identity key of the repo
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// - The RSA key to be used to sign alpine indecies
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
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a AlpineRepository resource.
type AlpineRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled  pulumi.BoolPtrInput
	BlackedOut              pulumi.BoolPtrInput
	Description             pulumi.StringPtrInput
	DownloadDirect          pulumi.BoolPtrInput
	ExcludesPattern         pulumi.StringPtrInput
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// - the identity key of the repo
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// - The RSA key to be used to sign alpine indecies
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
	XrayIndex     pulumi.BoolPtrInput
}

func (AlpineRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alpineRepositoryArgs)(nil)).Elem()
}

type AlpineRepositoryInput interface {
	pulumi.Input

	ToAlpineRepositoryOutput() AlpineRepositoryOutput
	ToAlpineRepositoryOutputWithContext(ctx context.Context) AlpineRepositoryOutput
}

func (*AlpineRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**AlpineRepository)(nil)).Elem()
}

func (i *AlpineRepository) ToAlpineRepositoryOutput() AlpineRepositoryOutput {
	return i.ToAlpineRepositoryOutputWithContext(context.Background())
}

func (i *AlpineRepository) ToAlpineRepositoryOutputWithContext(ctx context.Context) AlpineRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlpineRepositoryOutput)
}

// AlpineRepositoryArrayInput is an input type that accepts AlpineRepositoryArray and AlpineRepositoryArrayOutput values.
// You can construct a concrete instance of `AlpineRepositoryArrayInput` via:
//
//          AlpineRepositoryArray{ AlpineRepositoryArgs{...} }
type AlpineRepositoryArrayInput interface {
	pulumi.Input

	ToAlpineRepositoryArrayOutput() AlpineRepositoryArrayOutput
	ToAlpineRepositoryArrayOutputWithContext(context.Context) AlpineRepositoryArrayOutput
}

type AlpineRepositoryArray []AlpineRepositoryInput

func (AlpineRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlpineRepository)(nil)).Elem()
}

func (i AlpineRepositoryArray) ToAlpineRepositoryArrayOutput() AlpineRepositoryArrayOutput {
	return i.ToAlpineRepositoryArrayOutputWithContext(context.Background())
}

func (i AlpineRepositoryArray) ToAlpineRepositoryArrayOutputWithContext(ctx context.Context) AlpineRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlpineRepositoryArrayOutput)
}

// AlpineRepositoryMapInput is an input type that accepts AlpineRepositoryMap and AlpineRepositoryMapOutput values.
// You can construct a concrete instance of `AlpineRepositoryMapInput` via:
//
//          AlpineRepositoryMap{ "key": AlpineRepositoryArgs{...} }
type AlpineRepositoryMapInput interface {
	pulumi.Input

	ToAlpineRepositoryMapOutput() AlpineRepositoryMapOutput
	ToAlpineRepositoryMapOutputWithContext(context.Context) AlpineRepositoryMapOutput
}

type AlpineRepositoryMap map[string]AlpineRepositoryInput

func (AlpineRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlpineRepository)(nil)).Elem()
}

func (i AlpineRepositoryMap) ToAlpineRepositoryMapOutput() AlpineRepositoryMapOutput {
	return i.ToAlpineRepositoryMapOutputWithContext(context.Background())
}

func (i AlpineRepositoryMap) ToAlpineRepositoryMapOutputWithContext(ctx context.Context) AlpineRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlpineRepositoryMapOutput)
}

type AlpineRepositoryOutput struct{ *pulumi.OutputState }

func (AlpineRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlpineRepository)(nil)).Elem()
}

func (o AlpineRepositoryOutput) ToAlpineRepositoryOutput() AlpineRepositoryOutput {
	return o
}

func (o AlpineRepositoryOutput) ToAlpineRepositoryOutputWithContext(ctx context.Context) AlpineRepositoryOutput {
	return o
}

type AlpineRepositoryArrayOutput struct{ *pulumi.OutputState }

func (AlpineRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlpineRepository)(nil)).Elem()
}

func (o AlpineRepositoryArrayOutput) ToAlpineRepositoryArrayOutput() AlpineRepositoryArrayOutput {
	return o
}

func (o AlpineRepositoryArrayOutput) ToAlpineRepositoryArrayOutputWithContext(ctx context.Context) AlpineRepositoryArrayOutput {
	return o
}

func (o AlpineRepositoryArrayOutput) Index(i pulumi.IntInput) AlpineRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlpineRepository {
		return vs[0].([]*AlpineRepository)[vs[1].(int)]
	}).(AlpineRepositoryOutput)
}

type AlpineRepositoryMapOutput struct{ *pulumi.OutputState }

func (AlpineRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlpineRepository)(nil)).Elem()
}

func (o AlpineRepositoryMapOutput) ToAlpineRepositoryMapOutput() AlpineRepositoryMapOutput {
	return o
}

func (o AlpineRepositoryMapOutput) ToAlpineRepositoryMapOutputWithContext(ctx context.Context) AlpineRepositoryMapOutput {
	return o
}

func (o AlpineRepositoryMapOutput) MapIndex(k pulumi.StringInput) AlpineRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlpineRepository {
		return vs[0].(map[string]*AlpineRepository)[vs[1].(string)]
	}).(AlpineRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlpineRepositoryInput)(nil)).Elem(), &AlpineRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlpineRepositoryArrayInput)(nil)).Elem(), AlpineRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlpineRepositoryMapInput)(nil)).Elem(), AlpineRepositoryMap{})
	pulumi.RegisterOutputType(AlpineRepositoryOutput{})
	pulumi.RegisterOutputType(AlpineRepositoryArrayOutput{})
	pulumi.RegisterOutputType(AlpineRepositoryMapOutput{})
}
