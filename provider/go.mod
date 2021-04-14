module github.com/pulumi/pulumi-xyz

go 1.16

require (
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-aws/sdk/v4 v4.0.0-beta.2
	github.com/pulumi/pulumi/pkg/v3 v3.0.0-beta.2
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-rc.1
)

replace github.com/pulumi/pulumi/pkg/v3 => ../../pulumi/pkg

replace github.com/pulumi/pulumi/sdk/v3 => ../../pulumi/sdk
