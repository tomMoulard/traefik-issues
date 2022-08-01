package org.example;

import com.amazonaws.auth.AWSCredentials;
import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.client.builder.AwsClientBuilder;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.ObjectMetadata;
import com.amazonaws.services.s3.model.PutObjectRequest;

import java.io.ByteArrayInputStream;

public class Test {

    public static void main(String[] args) {
        String bucketName = "test";
        String fileName = "emptyFile.txt";
        AWSCredentials credentials = new BasicAWSCredentials("admin", "admin123");

        final AmazonS3 s3 = AmazonS3ClientBuilder.standard()
                .withEndpointConfiguration(
                        new AwsClientBuilder.EndpointConfiguration(
                                "http://localhost:9100", Regions.DEFAULT_REGION.getName()))
                .enablePathStyleAccess()
                .withCredentials(new AWSStaticCredentialsProvider(credentials))
                .build();


        if (!s3.doesBucketExistV2(bucketName)) {
            s3.createBucket(bucketName);
        }

        ObjectMetadata metadata = new ObjectMetadata();
        metadata.setContentLength(0);

        PutObjectRequest putObjectRequest =
                new PutObjectRequest(bucketName, fileName, new ByteArrayInputStream(new byte[0]), metadata);

        s3.putObject(putObjectRequest);
    }
}
