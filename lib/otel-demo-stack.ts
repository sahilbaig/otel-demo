import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";

export class OtelDemoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // Go Lambda
    const goFn = new GoFunction(this, "HelloLambda", {
      entry: "lambda-go",
    });

    const goUrl = goFn.addFunctionUrl({
      authType: lambda.FunctionUrlAuthType.NONE,
    });

    // JavaScript Lambda
    const jsFn = new lambda.Function(this, "JsLambda", {
      runtime: lambda.Runtime.NODEJS_22_X,
      handler: "index.handler",
      code: lambda.Code.fromAsset("lambda-js"),
    });

    const jsUrl = jsFn.addFunctionUrl({
      authType: lambda.FunctionUrlAuthType.NONE,
    });

    // Outputs
    new cdk.CfnOutput(this, "HelloLambdaUrl", {
      value: goUrl.url,
    });

    new cdk.CfnOutput(this, "JsLambdaUrl", {
      value: jsUrl.url,
    });
  }
}
