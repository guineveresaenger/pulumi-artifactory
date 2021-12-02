// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Gradle Repository Resource
//
// Creates a local Maven repository and allows for the creation of a
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
// 		_, err := artifactory.NewLocalMavenRepository(ctx, "terraform_local_test_maven_repo_basic", &artifactory.LocalMavenRepositoryArgs{
// 			ChecksumPolicyType:           pulumi.String("client-checksums"),
// 			HandleReleases:               pulumi.Bool(true),
// 			HandleSnapshots:              pulumi.Bool(true),
// 			Key:                          pulumi.String("terraform-local-test-maven-repo-basic"),
// 			MaxUniqueSnapshots:           pulumi.Int(10),
// 			SnapshotVersionBehavior:      pulumi.String("unique"),
// 			SuppressPomConsistencyChecks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalMavenRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// - Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	//   "resource is missing or conflicts with the locally calculated checksum (bad checksum). For more details,
	//   "please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	ChecksumPolicyType pulumi.StringPtrOutput `pulumi:"checksumPolicyType"`
	Description        pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect     pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern    pulumi.StringOutput    `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrOutput `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrOutput `pulumi:"handleSnapshots"`
	IncludesPattern pulumi.StringOutput  `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrOutput      `pulumi:"maxUniqueSnapshots"`
	Notes              pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType        pulumi.StringOutput      `pulumi:"packageType"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// Unique: Version number is based on a time-stamp (default)
	// Non-unique: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// Deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrOutput `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository
	SuppressPomConsistencyChecks pulumi.BoolPtrOutput `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolOutput    `pulumi:"xrayIndex"`
}

// NewLocalMavenRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalMavenRepository(ctx *pulumi.Context,
	name string, args *LocalMavenRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalMavenRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalMavenRepository
	err := ctx.RegisterResource("artifactory:index/localMavenRepository:LocalMavenRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalMavenRepository gets an existing LocalMavenRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalMavenRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalMavenRepositoryState, opts ...pulumi.ResourceOption) (*LocalMavenRepository, error) {
	var resource LocalMavenRepository
	err := ctx.ReadResource("artifactory:index/localMavenRepository:LocalMavenRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalMavenRepository resources.
type localMavenRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// - Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	//   "resource is missing or conflicts with the locally calculated checksum (bad checksum). For more details,
	//   "please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int     `pulumi:"maxUniqueSnapshots"`
	Notes              *string  `pulumi:"notes"`
	PackageType        *string  `pulumi:"packageType"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// Unique: Version number is based on a time-stamp (default)
	// Non-unique: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// Deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool `pulumi:"xrayIndex"`
}

type LocalMavenRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	// - Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	//   "resource is missing or conflicts with the locally calculated checksum (bad checksum). For more details,
	//   "please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	ChecksumPolicyType pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	DownloadDirect     pulumi.BoolPtrInput
	ExcludesPattern    pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput
	IncludesPattern pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	PackageType        pulumi.StringPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// Unique: Version number is based on a time-stamp (default)
	// Non-unique: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// Deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	XrayIndex                    pulumi.BoolPtrInput
}

func (LocalMavenRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localMavenRepositoryState)(nil)).Elem()
}

type localMavenRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// - Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	//   "resource is missing or conflicts with the locally calculated checksum (bad checksum). For more details,
	//   "please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int     `pulumi:"maxUniqueSnapshots"`
	Notes              *string  `pulumi:"notes"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// Unique: Version number is based on a time-stamp (default)
	// Non-unique: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// Deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalMavenRepository resource.
type LocalMavenRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	// - Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	//   "resource is missing or conflicts with the locally calculated checksum (bad checksum). For more details,
	//   "please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	ChecksumPolicyType pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	DownloadDirect     pulumi.BoolPtrInput
	ExcludesPattern    pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput
	IncludesPattern pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// Unique: Version number is based on a time-stamp (default)
	// Non-unique: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// Deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	XrayIndex                    pulumi.BoolPtrInput
}

func (LocalMavenRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localMavenRepositoryArgs)(nil)).Elem()
}

type LocalMavenRepositoryInput interface {
	pulumi.Input

	ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput
	ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput
}

func (*LocalMavenRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalMavenRepository)(nil))
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput {
	return i.ToLocalMavenRepositoryOutputWithContext(context.Background())
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryOutput)
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryPtrOutput() LocalMavenRepositoryPtrOutput {
	return i.ToLocalMavenRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryPtrOutputWithContext(ctx context.Context) LocalMavenRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryPtrOutput)
}

type LocalMavenRepositoryPtrInput interface {
	pulumi.Input

	ToLocalMavenRepositoryPtrOutput() LocalMavenRepositoryPtrOutput
	ToLocalMavenRepositoryPtrOutputWithContext(ctx context.Context) LocalMavenRepositoryPtrOutput
}

type localMavenRepositoryPtrType LocalMavenRepositoryArgs

func (*localMavenRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalMavenRepository)(nil))
}

func (i *localMavenRepositoryPtrType) ToLocalMavenRepositoryPtrOutput() LocalMavenRepositoryPtrOutput {
	return i.ToLocalMavenRepositoryPtrOutputWithContext(context.Background())
}

func (i *localMavenRepositoryPtrType) ToLocalMavenRepositoryPtrOutputWithContext(ctx context.Context) LocalMavenRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryPtrOutput)
}

// LocalMavenRepositoryArrayInput is an input type that accepts LocalMavenRepositoryArray and LocalMavenRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalMavenRepositoryArrayInput` via:
//
//          LocalMavenRepositoryArray{ LocalMavenRepositoryArgs{...} }
type LocalMavenRepositoryArrayInput interface {
	pulumi.Input

	ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput
	ToLocalMavenRepositoryArrayOutputWithContext(context.Context) LocalMavenRepositoryArrayOutput
}

type LocalMavenRepositoryArray []LocalMavenRepositoryInput

func (LocalMavenRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalMavenRepository)(nil)).Elem()
}

func (i LocalMavenRepositoryArray) ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput {
	return i.ToLocalMavenRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalMavenRepositoryArray) ToLocalMavenRepositoryArrayOutputWithContext(ctx context.Context) LocalMavenRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryArrayOutput)
}

// LocalMavenRepositoryMapInput is an input type that accepts LocalMavenRepositoryMap and LocalMavenRepositoryMapOutput values.
// You can construct a concrete instance of `LocalMavenRepositoryMapInput` via:
//
//          LocalMavenRepositoryMap{ "key": LocalMavenRepositoryArgs{...} }
type LocalMavenRepositoryMapInput interface {
	pulumi.Input

	ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput
	ToLocalMavenRepositoryMapOutputWithContext(context.Context) LocalMavenRepositoryMapOutput
}

type LocalMavenRepositoryMap map[string]LocalMavenRepositoryInput

func (LocalMavenRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalMavenRepository)(nil)).Elem()
}

func (i LocalMavenRepositoryMap) ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput {
	return i.ToLocalMavenRepositoryMapOutputWithContext(context.Background())
}

func (i LocalMavenRepositoryMap) ToLocalMavenRepositoryMapOutputWithContext(ctx context.Context) LocalMavenRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryMapOutput)
}

type LocalMavenRepositoryOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalMavenRepository)(nil))
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput {
	return o
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput {
	return o
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryPtrOutput() LocalMavenRepositoryPtrOutput {
	return o.ToLocalMavenRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryPtrOutputWithContext(ctx context.Context) LocalMavenRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalMavenRepository) *LocalMavenRepository {
		return &v
	}).(LocalMavenRepositoryPtrOutput)
}

type LocalMavenRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalMavenRepository)(nil))
}

func (o LocalMavenRepositoryPtrOutput) ToLocalMavenRepositoryPtrOutput() LocalMavenRepositoryPtrOutput {
	return o
}

func (o LocalMavenRepositoryPtrOutput) ToLocalMavenRepositoryPtrOutputWithContext(ctx context.Context) LocalMavenRepositoryPtrOutput {
	return o
}

func (o LocalMavenRepositoryPtrOutput) Elem() LocalMavenRepositoryOutput {
	return o.ApplyT(func(v *LocalMavenRepository) LocalMavenRepository {
		if v != nil {
			return *v
		}
		var ret LocalMavenRepository
		return ret
	}).(LocalMavenRepositoryOutput)
}

type LocalMavenRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalMavenRepository)(nil))
}

func (o LocalMavenRepositoryArrayOutput) ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput {
	return o
}

func (o LocalMavenRepositoryArrayOutput) ToLocalMavenRepositoryArrayOutputWithContext(ctx context.Context) LocalMavenRepositoryArrayOutput {
	return o
}

func (o LocalMavenRepositoryArrayOutput) Index(i pulumi.IntInput) LocalMavenRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalMavenRepository {
		return vs[0].([]LocalMavenRepository)[vs[1].(int)]
	}).(LocalMavenRepositoryOutput)
}

type LocalMavenRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalMavenRepository)(nil))
}

func (o LocalMavenRepositoryMapOutput) ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput {
	return o
}

func (o LocalMavenRepositoryMapOutput) ToLocalMavenRepositoryMapOutputWithContext(ctx context.Context) LocalMavenRepositoryMapOutput {
	return o
}

func (o LocalMavenRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalMavenRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalMavenRepository {
		return vs[0].(map[string]LocalMavenRepository)[vs[1].(string)]
	}).(LocalMavenRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryInput)(nil)).Elem(), &LocalMavenRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryPtrInput)(nil)).Elem(), &LocalMavenRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryArrayInput)(nil)).Elem(), LocalMavenRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryMapInput)(nil)).Elem(), LocalMavenRepositoryMap{})
	pulumi.RegisterOutputType(LocalMavenRepositoryOutput{})
	pulumi.RegisterOutputType(LocalMavenRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalMavenRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalMavenRepositoryMapOutput{})
}