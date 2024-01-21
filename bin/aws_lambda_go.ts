#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { AwsLambdaGoStack } from '../lib/aws_lambda_go-stack';

const app = new cdk.App();
new AwsLambdaGoStack(app, 'AwsLambdaGoStack', {});
