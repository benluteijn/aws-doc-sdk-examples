#snippet-sourcedescription:[<<FILENAME>> demonstrates how to ...]
#snippet-keyword:[Ruby]
#snippet-keyword:[Code Sample]
#snippet-keyword:[Amazon S3]
#snippet-service:[s3]
#snippet-sourcetype:[<<snippet or full-example>>]
#snippet-sourcedate:[]
#snippet-sourceauthor:[AWS]
# Copyright 2010-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# This file is licensed under the Apache License, Version 2.0 (the "License").
# You may not use this file except in compliance with the License. A copy of the
# License is located at
#
# http://aws.amazon.com/apache2.0/
#
# This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
# OF ANY KIND, either express or implied. See the License for the specific
# language governing permissions and limitations under the License.

require 'aws-sdk-s3' # v2: require 'aws-sdk'
require 'openssl'

bucket = 'my_bucket'
item = 'my_item'
key_file = 'public_key.pem'

# Get file content as string
contents = File.read(item)
public_key = File.read(key_file)

key = OpenSSL::PKey::RSA.new(public_key)

begin
  # encryption client
  enc_client = Aws::S3::Encryption::Client.new(encryption_key: key)

  # Add encrypted item to bucket
  enc_client.put_object(
    body: contents,
    bucket: bucket,
    key: item_name
  )

  puts 'Added ' + item_name + ' to bucket ' + bucket + ' using key from ' + key_file
rescue StandardError => err
  puts 'Could not add item'
  puts 'Error:'
  puts ex.message
end
