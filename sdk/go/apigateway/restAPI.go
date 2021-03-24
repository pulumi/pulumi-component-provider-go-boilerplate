// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RestAPI struct {
	pulumi.ResourceState

	RestAPI apigateway.RestApiOutput `pulumi:"restAPI"`
	Url     pulumi.StringOutput      `pulumi:"url"`
}

// NewRestAPI registers a new resource with the given unique name, arguments, and options.
func NewRestAPI(ctx *pulumi.Context,
	name string, args *RestAPIArgs, opts ...pulumi.ResourceOption) (*RestAPI, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Routes == nil {
		return nil, errors.New("invalid value for required argument 'Routes'")
	}
	var resource RestAPI
	err := ctx.RegisterRemoteComponentResource("apigateway:index:RestAPI", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type restAPIArgs struct {
	Routes []EventHandlerRoute `pulumi:"routes"`
}

// The set of arguments for constructing a RestAPI resource.
type RestAPIArgs struct {
	Routes EventHandlerRouteArrayInput
}

func (RestAPIArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restAPIArgs)(nil)).Elem()
}

type RestAPIInput interface {
	pulumi.Input

	ToRestAPIOutput() RestAPIOutput
	ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput
}

func (*RestAPI) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAPI)(nil))
}

func (i *RestAPI) ToRestAPIOutput() RestAPIOutput {
	return i.ToRestAPIOutputWithContext(context.Background())
}

func (i *RestAPI) ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIOutput)
}

func (i *RestAPI) ToRestAPIPtrOutput() RestAPIPtrOutput {
	return i.ToRestAPIPtrOutputWithContext(context.Background())
}

func (i *RestAPI) ToRestAPIPtrOutputWithContext(ctx context.Context) RestAPIPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIPtrOutput)
}

type RestAPIPtrInput interface {
	pulumi.Input

	ToRestAPIPtrOutput() RestAPIPtrOutput
	ToRestAPIPtrOutputWithContext(ctx context.Context) RestAPIPtrOutput
}

type restAPIPtrType RestAPIArgs

func (*restAPIPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAPI)(nil))
}

func (i *restAPIPtrType) ToRestAPIPtrOutput() RestAPIPtrOutput {
	return i.ToRestAPIPtrOutputWithContext(context.Background())
}

func (i *restAPIPtrType) ToRestAPIPtrOutputWithContext(ctx context.Context) RestAPIPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIPtrOutput)
}

// RestAPIArrayInput is an input type that accepts RestAPIArray and RestAPIArrayOutput values.
// You can construct a concrete instance of `RestAPIArrayInput` via:
//
//          RestAPIArray{ RestAPIArgs{...} }
type RestAPIArrayInput interface {
	pulumi.Input

	ToRestAPIArrayOutput() RestAPIArrayOutput
	ToRestAPIArrayOutputWithContext(context.Context) RestAPIArrayOutput
}

type RestAPIArray []RestAPIInput

func (RestAPIArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RestAPI)(nil))
}

func (i RestAPIArray) ToRestAPIArrayOutput() RestAPIArrayOutput {
	return i.ToRestAPIArrayOutputWithContext(context.Background())
}

func (i RestAPIArray) ToRestAPIArrayOutputWithContext(ctx context.Context) RestAPIArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIArrayOutput)
}

// RestAPIMapInput is an input type that accepts RestAPIMap and RestAPIMapOutput values.
// You can construct a concrete instance of `RestAPIMapInput` via:
//
//          RestAPIMap{ "key": RestAPIArgs{...} }
type RestAPIMapInput interface {
	pulumi.Input

	ToRestAPIMapOutput() RestAPIMapOutput
	ToRestAPIMapOutputWithContext(context.Context) RestAPIMapOutput
}

type RestAPIMap map[string]RestAPIInput

func (RestAPIMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RestAPI)(nil))
}

func (i RestAPIMap) ToRestAPIMapOutput() RestAPIMapOutput {
	return i.ToRestAPIMapOutputWithContext(context.Background())
}

func (i RestAPIMap) ToRestAPIMapOutputWithContext(ctx context.Context) RestAPIMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIMapOutput)
}

type RestAPIOutput struct {
	*pulumi.OutputState
}

func (RestAPIOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestAPI)(nil))
}

func (o RestAPIOutput) ToRestAPIOutput() RestAPIOutput {
	return o
}

func (o RestAPIOutput) ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput {
	return o
}

func (o RestAPIOutput) ToRestAPIPtrOutput() RestAPIPtrOutput {
	return o.ToRestAPIPtrOutputWithContext(context.Background())
}

func (o RestAPIOutput) ToRestAPIPtrOutputWithContext(ctx context.Context) RestAPIPtrOutput {
	return o.ApplyT(func(v RestAPI) *RestAPI {
		return &v
	}).(RestAPIPtrOutput)
}

type RestAPIPtrOutput struct {
	*pulumi.OutputState
}

func (RestAPIPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAPI)(nil))
}

func (o RestAPIPtrOutput) ToRestAPIPtrOutput() RestAPIPtrOutput {
	return o
}

func (o RestAPIPtrOutput) ToRestAPIPtrOutputWithContext(ctx context.Context) RestAPIPtrOutput {
	return o
}

type RestAPIArrayOutput struct{ *pulumi.OutputState }

func (RestAPIArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestAPI)(nil))
}

func (o RestAPIArrayOutput) ToRestAPIArrayOutput() RestAPIArrayOutput {
	return o
}

func (o RestAPIArrayOutput) ToRestAPIArrayOutputWithContext(ctx context.Context) RestAPIArrayOutput {
	return o
}

func (o RestAPIArrayOutput) Index(i pulumi.IntInput) RestAPIOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestAPI {
		return vs[0].([]RestAPI)[vs[1].(int)]
	}).(RestAPIOutput)
}

type RestAPIMapOutput struct{ *pulumi.OutputState }

func (RestAPIMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RestAPI)(nil))
}

func (o RestAPIMapOutput) ToRestAPIMapOutput() RestAPIMapOutput {
	return o
}

func (o RestAPIMapOutput) ToRestAPIMapOutputWithContext(ctx context.Context) RestAPIMapOutput {
	return o
}

func (o RestAPIMapOutput) MapIndex(k pulumi.StringInput) RestAPIOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RestAPI {
		return vs[0].(map[string]RestAPI)[vs[1].(string)]
	}).(RestAPIOutput)
}

func init() {
	pulumi.RegisterOutputType(RestAPIOutput{})
	pulumi.RegisterOutputType(RestAPIPtrOutput{})
	pulumi.RegisterOutputType(RestAPIArrayOutput{})
	pulumi.RegisterOutputType(RestAPIMapOutput{})
}
