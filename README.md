# gs-pack-calc

This is a basic project to create an api which takes in a list of pack sizes and a quantity and feeds out an array of packs needed to fulfill the order. Below is a bit more info on the project

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── pack-calculator             <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
│   └── main_test.go            <-- Unit tests
└── template.yaml
```

The url of the project is [https://28nmcxzugk.execute-api.eu-west-2.amazonaws.com/dev](https://28nmcxzugk.execute-api.eu-west-2.amazonaws.com/dev)
The methods available are:
POST - [/pack-calculator](https://28nmcxzugk.execute-api.eu-west-2.amazonaws.com/dev/pack-calculator)
The input required is:
```shell
{
    "packsarr": []int example: [ 250, 500, 1000, 2000, 5000],
    "quantity": int example: 501
}
```

The response is:
```shell
{
    PacksNeeded:[]int example: [ 500, 250 ]
}
```



## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

### Installing dependencies & building the target 

Before building your project cd into `pack-calculator` directory and run 

```bash
go mod init
```

then run

```bash
go mod tidy
```

In this example we use the built-in `sam build` to automatically download all the dependencies and package our build target.   

The `sam build` command is wrapped inside of the `Makefile`. To execute this simply run
 
```shell
make
```

### Local development

**Invoking function locally through local API Gateway**

```bash
sam local start-api
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/pack-calculator`

## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

```yaml
...
    GetPacksFunction:
        Type: AWS::Serverless::Function
        Properties:
            CodeUri: pack-calculator/
            ...
```

To deploy your application for the first time, make sure you have built your project then run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Parameter Stage**: This is the stage that will be created within API Gateway for your API
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: The template needs to create roles in order to run the api and function.
* **GetPacksFunction may not have authorization defined, Is this ok?**: This message appears because there is no protection on the api.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

### Testing

To run the tests navigate to the `pack-calculator` directory and run:

```shell
go test -v
```