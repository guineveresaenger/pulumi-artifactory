// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Release Bundle Webhook Resource
//
// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// .
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewReleaseBundleWebhook(ctx, "release-bundle-webhook", &artifactory.ReleaseBundleWebhookArgs{
// 			Criteria: &ReleaseBundleWebhookCriteriaArgs{
// 				AnyReleaseBundle: pulumi.Bool(false),
// 				ExcludePatterns: pulumi.StringArray{
// 					pulumi.String("bar/**"),
// 				},
// 				IncludePatterns: pulumi.StringArray{
// 					pulumi.String("foo/**"),
// 				},
// 				RegisteredReleaseBundleNames: pulumi.StringArray{
// 					pulumi.String("bundle-name"),
// 				},
// 			},
// 			CustomHttpHeaders: pulumi.StringMap{
// 				"header-1": pulumi.String("value-1"),
// 				"header-2": pulumi.String("value-2"),
// 			},
// 			EventTypes: pulumi.StringArray{
// 				pulumi.String("created"),
// 				pulumi.String("signed"),
// 				pulumi.String("deleted"),
// 			},
// 			Key:    pulumi.String("release-bundle-webhook"),
// 			Proxy:  pulumi.String("proxy-key"),
// 			Secret: pulumi.String("some-secret"),
// 			Url:    pulumi.String("http://tempurl.org/webhook"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ArtifactoryReleaseBundleWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleWebhookCriteriaOutput `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapOutput `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewArtifactoryReleaseBundleWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactoryReleaseBundleWebhook(ctx *pulumi.Context,
	name string, args *ArtifactoryReleaseBundleWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactoryReleaseBundleWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource ArtifactoryReleaseBundleWebhook
	err := ctx.RegisterResource("artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactoryReleaseBundleWebhook gets an existing ArtifactoryReleaseBundleWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactoryReleaseBundleWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactoryReleaseBundleWebhookState, opts ...pulumi.ResourceOption) (*ArtifactoryReleaseBundleWebhook, error) {
	var resource ArtifactoryReleaseBundleWebhook
	err := ctx.ReadResource("artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactoryReleaseBundleWebhook resources.
type artifactoryReleaseBundleWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ArtifactoryReleaseBundleWebhookCriteria `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url *string `pulumi:"url"`
}

type ArtifactoryReleaseBundleWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleWebhookCriteriaPtrInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringPtrInput
}

func (ArtifactoryReleaseBundleWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactoryReleaseBundleWebhookState)(nil)).Elem()
}

type artifactoryReleaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleWebhookCriteria `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ArtifactoryReleaseBundleWebhook resource.
type ArtifactoryReleaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleWebhookCriteriaInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringInput
}

func (ArtifactoryReleaseBundleWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactoryReleaseBundleWebhookArgs)(nil)).Elem()
}

type ArtifactoryReleaseBundleWebhookInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleWebhookOutput() ArtifactoryReleaseBundleWebhookOutput
	ToArtifactoryReleaseBundleWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookOutput
}

func (*ArtifactoryReleaseBundleWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (i *ArtifactoryReleaseBundleWebhook) ToArtifactoryReleaseBundleWebhookOutput() ArtifactoryReleaseBundleWebhookOutput {
	return i.ToArtifactoryReleaseBundleWebhookOutputWithContext(context.Background())
}

func (i *ArtifactoryReleaseBundleWebhook) ToArtifactoryReleaseBundleWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleWebhookOutput)
}

// ArtifactoryReleaseBundleWebhookArrayInput is an input type that accepts ArtifactoryReleaseBundleWebhookArray and ArtifactoryReleaseBundleWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactoryReleaseBundleWebhookArrayInput` via:
//
//          ArtifactoryReleaseBundleWebhookArray{ ArtifactoryReleaseBundleWebhookArgs{...} }
type ArtifactoryReleaseBundleWebhookArrayInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleWebhookArrayOutput() ArtifactoryReleaseBundleWebhookArrayOutput
	ToArtifactoryReleaseBundleWebhookArrayOutputWithContext(context.Context) ArtifactoryReleaseBundleWebhookArrayOutput
}

type ArtifactoryReleaseBundleWebhookArray []ArtifactoryReleaseBundleWebhookInput

func (ArtifactoryReleaseBundleWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (i ArtifactoryReleaseBundleWebhookArray) ToArtifactoryReleaseBundleWebhookArrayOutput() ArtifactoryReleaseBundleWebhookArrayOutput {
	return i.ToArtifactoryReleaseBundleWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactoryReleaseBundleWebhookArray) ToArtifactoryReleaseBundleWebhookArrayOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleWebhookArrayOutput)
}

// ArtifactoryReleaseBundleWebhookMapInput is an input type that accepts ArtifactoryReleaseBundleWebhookMap and ArtifactoryReleaseBundleWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactoryReleaseBundleWebhookMapInput` via:
//
//          ArtifactoryReleaseBundleWebhookMap{ "key": ArtifactoryReleaseBundleWebhookArgs{...} }
type ArtifactoryReleaseBundleWebhookMapInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleWebhookMapOutput() ArtifactoryReleaseBundleWebhookMapOutput
	ToArtifactoryReleaseBundleWebhookMapOutputWithContext(context.Context) ArtifactoryReleaseBundleWebhookMapOutput
}

type ArtifactoryReleaseBundleWebhookMap map[string]ArtifactoryReleaseBundleWebhookInput

func (ArtifactoryReleaseBundleWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (i ArtifactoryReleaseBundleWebhookMap) ToArtifactoryReleaseBundleWebhookMapOutput() ArtifactoryReleaseBundleWebhookMapOutput {
	return i.ToArtifactoryReleaseBundleWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactoryReleaseBundleWebhookMap) ToArtifactoryReleaseBundleWebhookMapOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleWebhookMapOutput)
}

type ArtifactoryReleaseBundleWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleWebhookOutput) ToArtifactoryReleaseBundleWebhookOutput() ArtifactoryReleaseBundleWebhookOutput {
	return o
}

func (o ArtifactoryReleaseBundleWebhookOutput) ToArtifactoryReleaseBundleWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookOutput {
	return o
}

type ArtifactoryReleaseBundleWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleWebhookArrayOutput) ToArtifactoryReleaseBundleWebhookArrayOutput() ArtifactoryReleaseBundleWebhookArrayOutput {
	return o
}

func (o ArtifactoryReleaseBundleWebhookArrayOutput) ToArtifactoryReleaseBundleWebhookArrayOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookArrayOutput {
	return o
}

func (o ArtifactoryReleaseBundleWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactoryReleaseBundleWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactoryReleaseBundleWebhook {
		return vs[0].([]*ArtifactoryReleaseBundleWebhook)[vs[1].(int)]
	}).(ArtifactoryReleaseBundleWebhookOutput)
}

type ArtifactoryReleaseBundleWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactoryReleaseBundleWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleWebhookMapOutput) ToArtifactoryReleaseBundleWebhookMapOutput() ArtifactoryReleaseBundleWebhookMapOutput {
	return o
}

func (o ArtifactoryReleaseBundleWebhookMapOutput) ToArtifactoryReleaseBundleWebhookMapOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleWebhookMapOutput {
	return o
}

func (o ArtifactoryReleaseBundleWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactoryReleaseBundleWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactoryReleaseBundleWebhook {
		return vs[0].(map[string]*ArtifactoryReleaseBundleWebhook)[vs[1].(string)]
	}).(ArtifactoryReleaseBundleWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleWebhookInput)(nil)).Elem(), &ArtifactoryReleaseBundleWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleWebhookArrayInput)(nil)).Elem(), ArtifactoryReleaseBundleWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleWebhookMapInput)(nil)).Elem(), ArtifactoryReleaseBundleWebhookMap{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleWebhookMapOutput{})
}
