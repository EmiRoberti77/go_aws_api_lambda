import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { RestApi, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';

export class AwsLambdaGoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myFunction = new lambda.Function(this, 'go_lambda', {
      code: lambda.Code.fromAsset('lambdas'),
      handler: 'main',
      runtime: lambda.Runtime.GO_1_X,
    });

    const gateway = new RestApi(this, 'go_api', {
      defaultCorsPreflightOptions: {
        allowOrigins: ['*'],
        allowMethods: ['GET', 'POST', 'OPTIONS', 'PUT', 'DELETE'],
      },
    });

    const integration = new LambdaIntegration(myFunction);
    const resource = gateway.root.addResource('sync');
    resource.addMethod('GET', integration);
    resource.addMethod('POST', integration);
  }
}
