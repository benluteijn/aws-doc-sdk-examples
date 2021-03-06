 
//snippet-sourcedescription:[<<FILENAME>> demonstrates how to ...]
//snippet-keyword:[Go]
//snippet-keyword:[Code Sample]
//snippet-service:[<<ADD SERVICE>>]
//snippet-sourcetype:[<<snippet or full-example>>]
//snippet-sourcedate:[]
//snippet-sourceauthor:[AWS]


/*
   Copyright 2010-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"

    "flag"
    "fmt"
    "os"
    "encoding/json"
)

func main() {
    bucket := "myBucket"

    // Initialize a session that the SDK uses to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and the region from the shard configuration file ~/.aws/config.
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := s3.New(sess)

    PolicyDoc := map[string]interface{}{
        "Version": "2012-10-17",
        "Statement": []map[string]interface{}{
            {
                "Sid": "DenyIncorrectEncryptionHeader",
                "Effect": "Deny",
                "Principal": "*",
                "Action": "s3:PutObject",
                "Resource": "arn:aws:s3:::" + bucket + "/*",
                "Condition": map[string]interface{}{
                    "StringNotEquals": map[string]interface{}{
                        "s3:x-amz-server-side-encryption": "aws:kms",
                    },
                },
            },
            {
                "Sid": "DenyUnEncryptedObjectUploads",
                "Effect": "Deny",
                "Principal": "*",
                "Action": "s3:PutObject",
                "Resource": "arn:aws:s3:::" + bucket + "/*",
                "Condition": map[string]interface{}{
                    "Null": map[string]interface{}{
                        "s3:x-amz-server-side-encryption": "true",
                    },
                },
            },
        },
    }

    // Marshal the policy into a JSON value so that it can be sent to S3.
    policy, err := json.Marshal(PolicyDoc)
    if err != nil {
        fmt.Println("Error marshalling policy:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    input := &s3.PutBucketPolicyInput{
        Bucket: bucketPtr,
        Policy: aws.String(string(policy)),
    }

    _, err = svc.PutBucketPolicy(input)
    if err != nil {
        fmt.Println("Got an error adding policy to bucket " + bucket + ":")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println("Set policy for " + bucket)
}
