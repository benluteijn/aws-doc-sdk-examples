<?php
/**
 * Copyright 2010-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License. A copy of
 * the License is located at
 *
 * http://aws.amazon.com/apache2.0/
 *
 * This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 *  ABOUT THIS PHP SAMPLE: This sample is part of the
 *
 */

require 'vendor/autoload.php';

use Aws\Firehose\FirehoseClient;
use Aws\Exception\AwsException;

/**
 * Add a data blog to an existing Amazon Kinesis Firehose Delivery Stream.
 *
 * This code expects that you have AWS credentials set up per:
 * https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/guide_credentials.html
 */

//Create a KinesisClient
$FirehoseClient = new Aws\Firehose\FirehoseClient([
    'profile' => 'default',
    'version' => '2015-08-04',
    'region' => 'us-east-2'
]);

$name = "my_stream_name";
$content = '{"ticker_symbol":"QXZ", "sector":"HEALTHCARE", "change":-0.05, "price":84.51}';

try {
    $result = $FirehoseClient->putRecord([
        'DeliveryStreamName' => $name,
        'Record' => [
            'Data' => $content,
        ],
    ]);
    var_dump($result);
} catch (AwsException $e) {
    // output error message if fails
    echo $e->getMessage();
    echo "\n";
}
 

//snippet-sourcedescription:[PutRecordtoDeliveryStream.php demonstrates how to send data through Firehose Delivery Streams to your specified destination. After you create a Delivery Stream, use DescribeDeliveryStream to see if the Delivery Stream is active before sending data to it.]
//snippet-keyword:[PHP]
//snippet-keyword:[AWS SDK for PHP v3]
//snippet-keyword:[Code Sample]
//snippet-keyword:[Amazon Kinesis Data Firehose]
//snippet-service:[firehose]
//snippet-sourcetype:[full-example]
//snippet-sourcedate:[2018-09-20]
//snippet-sourceauthor:[jschwarzwalder]

