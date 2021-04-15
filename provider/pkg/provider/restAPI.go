// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The set of arguments for creating a RestAPI component resource.
type RestAPIArgs struct {
	Routes RestAPIRouteArrayInput `pulumi:"routes"`
}

type RestAPIRouteArrayInput interface {
	pulumi.Input
	ToRestAPIRouteArrayOutput() RestAPIRouteArrayOutput
	ToRestAPIRouteArrayOutputWithContext(ctx context.Context) RestAPIRouteArrayOutput
}

type RestAPIRouteArrayOutput struct{ *pulumi.OutputState }

func (o RestAPIRouteArrayOutput) ToRestAPIRouteArrayOutput() RestAPIRouteArrayOutput {
	return o
}

func (o RestAPIRouteArrayOutput) ToRestAPIRouteArrayOutputWithContext(ctx context.Context) RestAPIRouteArrayOutput {
	return o
}

func (RestAPIRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestAPIRoute)(nil)).Elem()
}

type RestAPIRoute struct {
	// Path     string
	// Method   string
	// Function *lambda.Function
}

func init() {
	pulumi.RegisterOutputType(RestAPIRouteArrayOutput{})
}

// The RestAPI component resource.
type RestAPI struct {
	pulumi.ResourceState

	name   string
	ctx    *pulumi.Context
	parent pulumi.ResourceOption

	API   *apigateway.RestApi `pulumi:"api"`
	Stage *apigateway.Stage   `pulumi:"stage"`
	Url   pulumi.StringOutput `pulumi:"url"`
}

// NewRestAPI creates a new RestAPI component resource.
func NewRestAPI(ctx *pulumi.Context,
	name string, args *RestAPIArgs, opts ...pulumi.ResourceOption) (*RestAPI, error) {
	if args == nil {
		args = &RestAPIArgs{}
	}

	r := &RestAPI{
		ctx:  ctx,
		name: name,
	}
	r.parent = pulumi.Parent(r)
	err := ctx.RegisterComponentResource("apigateway:index:RestAPI", name, r, opts...)
	if err != nil {
		return nil, err
	}

	swagger, err := r.createSwaggerSpec(args.Routes)
	if err != nil {
		return nil, err
	}

	restApi, err := apigateway.NewRestApi(ctx, name, &apigateway.RestApiArgs{
		BinaryMediaTypes: pulumi.StringArray{pulumi.String("*/*")},
		Body:             swagger,
	}, pulumi.Parent(r))
	if err != nil {
		return nil, err
	}

	version := swagger.ToStringOutput().ApplyT(func(s string) (string, error) {
		h := sha1.New()
		h.Write([]byte(s))
		sha1_hash := hex.EncodeToString(h.Sum(nil))
		return sha1_hash[0:8], nil
	}).(pulumi.StringOutput)

	// Create a deployment of the Rest API.
	deployment, err := apigateway.NewDeployment(ctx, name, &apigateway.DeploymentArgs{
		RestApi: restApi,
		// Note: Set to empty to avoid creating an implicit stage, we'll create it explicitly below instead.
		StageName: pulumi.String(""),
		// Note: We set `variables` here because it forces recreation of the Deployment object
		// whenever the body hash changes.  Because we use a blank stage name above, there will
		// not actually be any stage created in AWS, and thus these variables will not actually
		// end up anywhere.  But this will still cause the right replacement of the Deployment
		// when needed.  The Stage allocated below will be the stable stage that always points
		// to the latest deployment of the API.
		Variables: pulumi.StringMap{"version": version},
	}, pulumi.Parent(r))
	if err != nil {
		return nil, err
	}

	// this.swaggerLambdas = swaggerLambdas || new Map();
	// const permissions = createLambdaPermissions(this, name, this.swaggerLambdas);

	// // Create a stage, which is an addressable instance of the Rest API. Set it to point at the latest deployment.
	stage, err := apigateway.NewStage(ctx, name, &apigateway.StageArgs{
		StageName:  pulumi.String("stage"),
		RestApi:    restApi,
		Deployment: deployment,
		// TODO: `dependsOn: permissions`
	}, pulumi.Parent(r))
	if err != nil {
		return nil, err
	}

	r.API = restApi
	r.Stage = stage
	r.Url = pulumi.Sprintf("%sstage/", deployment.InvokeUrl)

	if err := ctx.RegisterResourceOutputs(r, pulumi.Map{
		"url": r.Url,
	}); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *RestAPI) createSwaggerSpec(routes RestAPIRouteArrayInput) (pulumi.StringOutput, error) {
	return routes.ToRestAPIRouteArrayOutput().ApplyT(func(routes []RestAPIRoute) (string, error) {
		swagger := swaggerSpec{
			Swagger: "2.0",
			Info: swaggerInfo{
				Title:   r.name,
				Version: "1.0",
			},
			Paths: map[string]map[string]swaggerOperation{},
		}
		for route := range routes {
			r.ctx.Log.Info(fmt.Sprintf("%v", route), nil)
		}
		byts, err := json.Marshal(swagger)
		if err != nil {
			return "", err
		}
		return string(byts), nil
	}).(pulumi.StringOutput), nil

}
