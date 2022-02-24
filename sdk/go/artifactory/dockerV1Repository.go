// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Docker V1 Repository Resource
//
// Creates a local Docker v1 repository - By choosing a V1 repository, you don't really have many options
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
// 		_, err := artifactory.NewDockerV2Repository(ctx, "foo", &artifactory.DockerV2RepositoryArgs{
// 			Key: pulumi.String("foo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DockerV1Repository struct {
	pulumi.CustomResourceState

	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolOutput      `pulumi:"blockPushingSchema1"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect      pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern     pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern     pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntOutput       `pulumi:"maxUniqueTags"`
	Notes         pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType   pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention pulumi.IntOutput `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewDockerV1Repository registers a new resource with the given unique name, arguments, and options.
func NewDockerV1Repository(ctx *pulumi.Context,
	name string, args *DockerV1RepositoryArgs, opts ...pulumi.ResourceOption) (*DockerV1Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource DockerV1Repository
	err := ctx.RegisterResource("artifactory:index/dockerV1Repository:DockerV1Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerV1Repository gets an existing DockerV1Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerV1Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerV1RepositoryState, opts ...pulumi.ResourceOption) (*DockerV1Repository, error) {
	var resource DockerV1Repository
	err := ctx.ReadResource("artifactory:index/dockerV1Repository:DockerV1Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerV1Repository resources.
type dockerV1RepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion *string `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 *bool   `pulumi:"blockPushingSchema1"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	IncludesPattern     *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
	Notes         *string `pulumi:"notes"`
	PackageType   *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type DockerV1RepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
	Description         pulumi.StringPtrInput
	DownloadDirect      pulumi.BoolPtrInput
	ExcludesPattern     pulumi.StringPtrInput
	IncludesPattern     pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	Notes         pulumi.StringPtrInput
	PackageType   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DockerV1RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV1RepositoryState)(nil)).Elem()
}

type dockerV1RepositoryArgs struct {
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
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
	Notes         *string `pulumi:"notes"`
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

// The set of arguments for constructing a DockerV1Repository resource.
type DockerV1RepositoryArgs struct {
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
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	Notes         pulumi.StringPtrInput
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

func (DockerV1RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV1RepositoryArgs)(nil)).Elem()
}

type DockerV1RepositoryInput interface {
	pulumi.Input

	ToDockerV1RepositoryOutput() DockerV1RepositoryOutput
	ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput
}

func (*DockerV1Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV1Repository)(nil)).Elem()
}

func (i *DockerV1Repository) ToDockerV1RepositoryOutput() DockerV1RepositoryOutput {
	return i.ToDockerV1RepositoryOutputWithContext(context.Background())
}

func (i *DockerV1Repository) ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryOutput)
}

// DockerV1RepositoryArrayInput is an input type that accepts DockerV1RepositoryArray and DockerV1RepositoryArrayOutput values.
// You can construct a concrete instance of `DockerV1RepositoryArrayInput` via:
//
//          DockerV1RepositoryArray{ DockerV1RepositoryArgs{...} }
type DockerV1RepositoryArrayInput interface {
	pulumi.Input

	ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput
	ToDockerV1RepositoryArrayOutputWithContext(context.Context) DockerV1RepositoryArrayOutput
}

type DockerV1RepositoryArray []DockerV1RepositoryInput

func (DockerV1RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV1Repository)(nil)).Elem()
}

func (i DockerV1RepositoryArray) ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput {
	return i.ToDockerV1RepositoryArrayOutputWithContext(context.Background())
}

func (i DockerV1RepositoryArray) ToDockerV1RepositoryArrayOutputWithContext(ctx context.Context) DockerV1RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryArrayOutput)
}

// DockerV1RepositoryMapInput is an input type that accepts DockerV1RepositoryMap and DockerV1RepositoryMapOutput values.
// You can construct a concrete instance of `DockerV1RepositoryMapInput` via:
//
//          DockerV1RepositoryMap{ "key": DockerV1RepositoryArgs{...} }
type DockerV1RepositoryMapInput interface {
	pulumi.Input

	ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput
	ToDockerV1RepositoryMapOutputWithContext(context.Context) DockerV1RepositoryMapOutput
}

type DockerV1RepositoryMap map[string]DockerV1RepositoryInput

func (DockerV1RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV1Repository)(nil)).Elem()
}

func (i DockerV1RepositoryMap) ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput {
	return i.ToDockerV1RepositoryMapOutputWithContext(context.Background())
}

func (i DockerV1RepositoryMap) ToDockerV1RepositoryMapOutputWithContext(ctx context.Context) DockerV1RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryMapOutput)
}

type DockerV1RepositoryOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryOutput) ToDockerV1RepositoryOutput() DockerV1RepositoryOutput {
	return o
}

func (o DockerV1RepositoryOutput) ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput {
	return o
}

type DockerV1RepositoryArrayOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryArrayOutput) ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput {
	return o
}

func (o DockerV1RepositoryArrayOutput) ToDockerV1RepositoryArrayOutputWithContext(ctx context.Context) DockerV1RepositoryArrayOutput {
	return o
}

func (o DockerV1RepositoryArrayOutput) Index(i pulumi.IntInput) DockerV1RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerV1Repository {
		return vs[0].([]*DockerV1Repository)[vs[1].(int)]
	}).(DockerV1RepositoryOutput)
}

type DockerV1RepositoryMapOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryMapOutput) ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput {
	return o
}

func (o DockerV1RepositoryMapOutput) ToDockerV1RepositoryMapOutputWithContext(ctx context.Context) DockerV1RepositoryMapOutput {
	return o
}

func (o DockerV1RepositoryMapOutput) MapIndex(k pulumi.StringInput) DockerV1RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerV1Repository {
		return vs[0].(map[string]*DockerV1Repository)[vs[1].(string)]
	}).(DockerV1RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryInput)(nil)).Elem(), &DockerV1Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryArrayInput)(nil)).Elem(), DockerV1RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryMapInput)(nil)).Elem(), DockerV1RepositoryMap{})
	pulumi.RegisterOutputType(DockerV1RepositoryOutput{})
	pulumi.RegisterOutputType(DockerV1RepositoryArrayOutput{})
	pulumi.RegisterOutputType(DockerV1RepositoryMapOutput{})
}
