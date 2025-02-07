// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Push Replication Resource
//
// Provides an Artifactory push replication resource. This can be used to create and manage Artifactory push replications.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		providerTestSource, err := artifactory.NewLocalMavenRepository(ctx, "providerTestSource", &artifactory.LocalMavenRepositoryArgs{
// 			Key: pulumi.String("provider_test_source"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewLocalMavenRepository(ctx, "providerTestDest", &artifactory.LocalMavenRepositoryArgs{
// 			Key: pulumi.String("provider_test_dest"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewPushReplication(ctx, "foo-rep", &artifactory.PushReplicationArgs{
// 			CronExp:                pulumi.String("0 0 * * * ?"),
// 			EnableEventReplication: pulumi.Bool(true),
// 			Replications: PushReplicationReplicationArray{
// 				&PushReplicationReplicationArgs{
// 					Password: pulumi.String(fmt.Sprintf("%v%v", "$", "var.artifactory_password")),
// 					Url:      pulumi.String(fmt.Sprintf("%v%v", "$", "var.artifactory_url")),
// 					Username: pulumi.String(fmt.Sprintf("%v%v", "$", "var.artifactory_username")),
// 				},
// 			},
// 			RepoKey: providerTestSource.Key,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Push replication configs can be imported using their repo key, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/pushReplication:PushReplication foo-rep provider_test_source
// ```
type PushReplication struct {
	pulumi.CustomResourceState

	CronExp                pulumi.StringOutput                   `pulumi:"cronExp"`
	EnableEventReplication pulumi.BoolOutput                     `pulumi:"enableEventReplication"`
	Replications           PushReplicationReplicationArrayOutput `pulumi:"replications"`
	RepoKey                pulumi.StringOutput                   `pulumi:"repoKey"`
}

// NewPushReplication registers a new resource with the given unique name, arguments, and options.
func NewPushReplication(ctx *pulumi.Context,
	name string, args *PushReplicationArgs, opts ...pulumi.ResourceOption) (*PushReplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExp == nil {
		return nil, errors.New("invalid value for required argument 'CronExp'")
	}
	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	var resource PushReplication
	err := ctx.RegisterResource("artifactory:index/pushReplication:PushReplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPushReplication gets an existing PushReplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPushReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PushReplicationState, opts ...pulumi.ResourceOption) (*PushReplication, error) {
	var resource PushReplication
	err := ctx.ReadResource("artifactory:index/pushReplication:PushReplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PushReplication resources.
type pushReplicationState struct {
	CronExp                *string                      `pulumi:"cronExp"`
	EnableEventReplication *bool                        `pulumi:"enableEventReplication"`
	Replications           []PushReplicationReplication `pulumi:"replications"`
	RepoKey                *string                      `pulumi:"repoKey"`
}

type PushReplicationState struct {
	CronExp                pulumi.StringPtrInput
	EnableEventReplication pulumi.BoolPtrInput
	Replications           PushReplicationReplicationArrayInput
	RepoKey                pulumi.StringPtrInput
}

func (PushReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*pushReplicationState)(nil)).Elem()
}

type pushReplicationArgs struct {
	CronExp                string                       `pulumi:"cronExp"`
	EnableEventReplication *bool                        `pulumi:"enableEventReplication"`
	Replications           []PushReplicationReplication `pulumi:"replications"`
	RepoKey                string                       `pulumi:"repoKey"`
}

// The set of arguments for constructing a PushReplication resource.
type PushReplicationArgs struct {
	CronExp                pulumi.StringInput
	EnableEventReplication pulumi.BoolPtrInput
	Replications           PushReplicationReplicationArrayInput
	RepoKey                pulumi.StringInput
}

func (PushReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pushReplicationArgs)(nil)).Elem()
}

type PushReplicationInput interface {
	pulumi.Input

	ToPushReplicationOutput() PushReplicationOutput
	ToPushReplicationOutputWithContext(ctx context.Context) PushReplicationOutput
}

func (*PushReplication) ElementType() reflect.Type {
	return reflect.TypeOf((**PushReplication)(nil)).Elem()
}

func (i *PushReplication) ToPushReplicationOutput() PushReplicationOutput {
	return i.ToPushReplicationOutputWithContext(context.Background())
}

func (i *PushReplication) ToPushReplicationOutputWithContext(ctx context.Context) PushReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushReplicationOutput)
}

// PushReplicationArrayInput is an input type that accepts PushReplicationArray and PushReplicationArrayOutput values.
// You can construct a concrete instance of `PushReplicationArrayInput` via:
//
//          PushReplicationArray{ PushReplicationArgs{...} }
type PushReplicationArrayInput interface {
	pulumi.Input

	ToPushReplicationArrayOutput() PushReplicationArrayOutput
	ToPushReplicationArrayOutputWithContext(context.Context) PushReplicationArrayOutput
}

type PushReplicationArray []PushReplicationInput

func (PushReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushReplication)(nil)).Elem()
}

func (i PushReplicationArray) ToPushReplicationArrayOutput() PushReplicationArrayOutput {
	return i.ToPushReplicationArrayOutputWithContext(context.Background())
}

func (i PushReplicationArray) ToPushReplicationArrayOutputWithContext(ctx context.Context) PushReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushReplicationArrayOutput)
}

// PushReplicationMapInput is an input type that accepts PushReplicationMap and PushReplicationMapOutput values.
// You can construct a concrete instance of `PushReplicationMapInput` via:
//
//          PushReplicationMap{ "key": PushReplicationArgs{...} }
type PushReplicationMapInput interface {
	pulumi.Input

	ToPushReplicationMapOutput() PushReplicationMapOutput
	ToPushReplicationMapOutputWithContext(context.Context) PushReplicationMapOutput
}

type PushReplicationMap map[string]PushReplicationInput

func (PushReplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushReplication)(nil)).Elem()
}

func (i PushReplicationMap) ToPushReplicationMapOutput() PushReplicationMapOutput {
	return i.ToPushReplicationMapOutputWithContext(context.Background())
}

func (i PushReplicationMap) ToPushReplicationMapOutputWithContext(ctx context.Context) PushReplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushReplicationMapOutput)
}

type PushReplicationOutput struct{ *pulumi.OutputState }

func (PushReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushReplication)(nil)).Elem()
}

func (o PushReplicationOutput) ToPushReplicationOutput() PushReplicationOutput {
	return o
}

func (o PushReplicationOutput) ToPushReplicationOutputWithContext(ctx context.Context) PushReplicationOutput {
	return o
}

type PushReplicationArrayOutput struct{ *pulumi.OutputState }

func (PushReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PushReplication)(nil)).Elem()
}

func (o PushReplicationArrayOutput) ToPushReplicationArrayOutput() PushReplicationArrayOutput {
	return o
}

func (o PushReplicationArrayOutput) ToPushReplicationArrayOutputWithContext(ctx context.Context) PushReplicationArrayOutput {
	return o
}

func (o PushReplicationArrayOutput) Index(i pulumi.IntInput) PushReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PushReplication {
		return vs[0].([]*PushReplication)[vs[1].(int)]
	}).(PushReplicationOutput)
}

type PushReplicationMapOutput struct{ *pulumi.OutputState }

func (PushReplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PushReplication)(nil)).Elem()
}

func (o PushReplicationMapOutput) ToPushReplicationMapOutput() PushReplicationMapOutput {
	return o
}

func (o PushReplicationMapOutput) ToPushReplicationMapOutputWithContext(ctx context.Context) PushReplicationMapOutput {
	return o
}

func (o PushReplicationMapOutput) MapIndex(k pulumi.StringInput) PushReplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PushReplication {
		return vs[0].(map[string]*PushReplication)[vs[1].(string)]
	}).(PushReplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PushReplicationInput)(nil)).Elem(), &PushReplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushReplicationArrayInput)(nil)).Elem(), PushReplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PushReplicationMapInput)(nil)).Elem(), PushReplicationMap{})
	pulumi.RegisterOutputType(PushReplicationOutput{})
	pulumi.RegisterOutputType(PushReplicationArrayOutput{})
	pulumi.RegisterOutputType(PushReplicationMapOutput{})
}
