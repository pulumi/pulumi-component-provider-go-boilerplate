package provider

type swaggerSpec struct {
	Swagger string                                 `json:"swagger"`
	Info    swaggerInfo                            `json:"info"`
	Paths   map[string]map[string]swaggerOperation `json:"paths"`
	//     "x-amazon-apigateway-binary-media-types"?: string[];
	//     "x-amazon-apigateway-gateway-responses": Record<string, SwaggerGatewayResponse>;
	//     securityDefinitions?: { [securityDefinitionName: string]: SecurityDefinition };
	//     "x-amazon-apigateway-request-validators"?: {
	//         [validatorName: string]: {
	//             validateRequestBody: boolean;
	//             validateRequestParameters: boolean;
	//         };
	//     };
	//     "x-amazon-apigateway-request-validator"?: RequestValidator;
	//     "x-amazon-apigateway-api-key-source"?: APIKeySource;
}

type swaggerInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type swaggerOperation struct {
	Parameters   []swaggerParameter         `json:"parameters,omitempty"`
	Responses    map[string]swaggerResponse `json:"responses,omitempty"`
	XIntegration apigatewayIntegration      `json:"x-amazon-apigateway-integration"`
	//     "x-amazon-apigateway-request-validator"?: RequestValidator;
	//     "x-amazon-apigateway-auth"?: ApigatewayAuth;
	//     /**
	//      * security a list of objects whose keys are the names of the authorizer. Each authorizer name
	//      * refers to a SecurityDefinition, defined at the top level of the swagger definition, by
	//      * matching a Security Definition's name property. For Cognito User Pool Authorizers, the value
	//      * of these object can be left as an empty array or used to define the resource servers and
	//      * custom scopes (e.g. "resource-server/scope"). For lambda authorizers, the value of the
	//      * objects is an empty array.
	//      */
	//     security?: Record<string, string[]>[];
}

type swaggerParameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Type     string `json:"type,omitempty"`
}

type swaggerResponse struct {
	Description string                   `json:"description"`
	Schema      *swaggerSchema           `json:"schema,omitempty"`
	Headers     map[string]swaggerHeader `json:"headers,omitempty"`
}

type swaggerSchema struct {
	Type string `json:"type"`
}

type swaggerHeaderType string

const (
	swaggerHeaderTypeString  swaggerHeaderType = "string"
	swaggerHeaderTypeNumber  swaggerHeaderType = "number"
	swaggerHeaderTypeInteger swaggerHeaderType = "integer"
	swaggerHeaderTypeBoolean swaggerHeaderType = "boolean"
	swaggerHeaderTypeArray   swaggerHeaderType = "array"
)

type swaggerHeader struct {
	Type  swaggerHeaderType `json:"type"`
	Items *swaggerHeader    `json:"items,omitempty"`
}

// export type IntegrationPassthroughBehavior = "when_no_match" | "when_no_templates" | "never";
type apigatewayIntegrationPassthroughBehavior string

// export type Method = "ANY" | "GET" | "PUT" | "POST" | "DELETE" | "PATCH" | "OPTIONS";
type apigatewayMethod string

// export type IntegrationType = "aws" | "aws_proxy" | "http" | "http_proxy" | "mock";
type apigatewayIntegrationType string

// export type IntegrationConnectionType = "INTERNET" | "VPC_LINK";
type apigatewayIntegrationConnectionType string

type apigatewayIntegration struct {
	RequestParameters   interface{}                              `json:"requestParameters,omitempty"`
	PassthroughBehavior apigatewayIntegrationPassthroughBehavior `json:"passthroughBehavior,omitempty"`
	HttpMethod          apigatewayMethod                         `json:"httpMethod"`
	Type                apigatewayIntegrationType                `json:"type"`
	Reponses            map[string]apigatewayIntegrationResponse `json:"responses,omitempty"`
	Uri                 string                                   `json:"uri"`
	ConnectionType      apigatewayIntegrationConnectionType      `json:"connectionType,omitempty"`
	ConnectionId        map[string]swaggerHeader                 `json:"connectionId,omitempty"`
	Credentials         map[string]swaggerHeader                 `json:"credentials,omitempty"`
}

type apigatewayIntegrationResponse struct {
	StatusCode         string            `json:"statusCode"`
	ResponseParameters map[string]string `json:"responseParameters,omitempty"`
}

// export type RequestValidator = "ALL" | "PARAMS_ONLY" | "BODY_ONLY";
// export type APIKeySource = "HEADER" | "AUTHORIZER";

// export interface SwaggerGatewayResponse {
//     statusCode: number;
//     responseTemplates: {
//         "application/json": string,
//     };
//     responseParameters?: {
//         [parameter: string]: string,
//     };
// }

// export interface SecurityDefinition {
//     type: "apiKey";
//     name: string;
//     in: "header" | "query";
//     "x-amazon-apigateway-authtype"?: string;
//     "x-amazon-apigateway-authorizer"?: SwaggerLambdaAuthorizer | SwaggerCognitoAuthorizer;
// }

// export interface SwaggerLambdaAuthorizer {
//     type: "token" | "request";
//     authorizerUri: pulumi.Input<string>;
//     authorizerCredentials: pulumi.Input<string>;
//     identitySource?: string;
//     identityValidationExpression?: string;
//     authorizerResultTtlInSeconds?: number;
// }

// export interface SwaggerCognitoAuthorizer {
//     type: "cognito_user_pools";
//     identitySource: string;
//     providerARNs: pulumi.Input<string>[];
//     authorizerResultTtlInSeconds?: number;
// }

// export interface ApigatewayAuth {
//     type: string;
// }
