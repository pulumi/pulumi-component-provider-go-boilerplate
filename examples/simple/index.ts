import * as apigateway from "@pulumi/apigateway";
import * as aws from "@pulumi/aws";

const f = new aws.lambda.Function("f", {

});

const api = new apigateway.RestAPI("api", {
    routes: [{
        path: "/",
        method: "GET",
        function: f,
    }],
});

export const url = api.url;
