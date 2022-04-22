// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Certificate Resource
//
// Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
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
// 		_, err := artifactory.NewCertificate(ctx, "my-cert", &artifactory.CertificateArgs{
// 			Alias:   pulumi.String("my-cert"),
// 			Content: readFileOrPanic("/path/to/bundle.pem"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewRemoteMavenRepository(ctx, "my-remote", &artifactory.RemoteMavenRepositoryArgs{
// 			ClientTlsCertificate: my_cert.Alias,
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
// Certificates can be imported using their alias, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// Name of certificate.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// PEM-encoded client certificate and private key.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	File    pulumi.StringPtrOutput `pulumi:"file"`
	// SHA256 fingerprint of the certificate.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Name of the certificate authority that issued the certificate.
	IssuedBy pulumi.StringOutput `pulumi:"issuedBy"`
	// The time & date when the certificate is valid from.
	IssuedOn pulumi.StringOutput `pulumi:"issuedOn"`
	// Name of whom the certificate has been issued to.
	IssuedTo pulumi.StringOutput `pulumi:"issuedTo"`
	// The time & date when the certificate expires.
	ValidUntil pulumi.StringOutput `pulumi:"validUntil"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	var resource Certificate
	err := ctx.RegisterResource("artifactory:index/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("artifactory:index/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Name of certificate.
	Alias *string `pulumi:"alias"`
	// PEM-encoded client certificate and private key.
	Content *string `pulumi:"content"`
	File    *string `pulumi:"file"`
	// SHA256 fingerprint of the certificate.
	Fingerprint *string `pulumi:"fingerprint"`
	// Name of the certificate authority that issued the certificate.
	IssuedBy *string `pulumi:"issuedBy"`
	// The time & date when the certificate is valid from.
	IssuedOn *string `pulumi:"issuedOn"`
	// Name of whom the certificate has been issued to.
	IssuedTo *string `pulumi:"issuedTo"`
	// The time & date when the certificate expires.
	ValidUntil *string `pulumi:"validUntil"`
}

type CertificateState struct {
	// Name of certificate.
	Alias pulumi.StringPtrInput
	// PEM-encoded client certificate and private key.
	Content pulumi.StringPtrInput
	File    pulumi.StringPtrInput
	// SHA256 fingerprint of the certificate.
	Fingerprint pulumi.StringPtrInput
	// Name of the certificate authority that issued the certificate.
	IssuedBy pulumi.StringPtrInput
	// The time & date when the certificate is valid from.
	IssuedOn pulumi.StringPtrInput
	// Name of whom the certificate has been issued to.
	IssuedTo pulumi.StringPtrInput
	// The time & date when the certificate expires.
	ValidUntil pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Name of certificate.
	Alias string `pulumi:"alias"`
	// PEM-encoded client certificate and private key.
	Content *string `pulumi:"content"`
	File    *string `pulumi:"file"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Name of certificate.
	Alias pulumi.StringInput
	// PEM-encoded client certificate and private key.
	Content pulumi.StringPtrInput
	File    pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//          CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//          CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
