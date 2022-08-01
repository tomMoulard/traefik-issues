# Traefik S3 Java Minio TLS Problem

## Description

The cominbation of the Java SDK for Amazon S3 with the S3 compatible data storage server MinIO, which is behind Traefik and the MinIO server is secured with TLS,
uploads of empty files against Traefik end in a SIGSEGV: segmentation violation error.
The error does not occur, if non-empty files are uploaded. It does also not occur if MinIO server is not secured with TLS and an empty file is uploaded.

Regardless of the segmentation fault, the empty file is uploaded correctly to MinIO and after a restart of Traefik can be properly viewed in the MinIO UI. 

This example is to reproduce this problem in a minimal setup.

## Requirements

To execute the sample on your host the following applications need to be installed:

* docker
* docker-compose
* Java >= 11 with set JAVA_HOME environment variable

## Content

This folder contains the following content:

* `compose-with-tls`: Docker compose setup for Traefik and two MinIO Instances which are secured with TLS.
* `java-s3-sample-upload`: Gradle Java Project - can be used with Gradle Wrapper so Gradle must not be installed, but Java must be available on host. (Gradle Wrapper uses environment variable JAVA_HOME to find Java)

## How to run

### Start the Docker Compose with TLS

Open a terminal in the folder and switch to the folder `compose-with-tls` and start it:

```bash
cd compose-with-tls
docker-compose up
```

When finished it should be possible to access `http://localhost:9100/`. MinIO should ask for credentials - 
these are `admin` for the user and `admin123` for the password. The MinIO UI should now be visible, but not contain much content, which is expected, as no structure and files are available yet.

### Run the Java Sample

Open another terminal in the folder and switch to the folder `java-s3-sample-upload` and start it:

```bash
cd java-s3-sample-upload
./gradlew run
```

This executes the only contained Java class `Test`, which uses the S3 Amazon Java SDK to upload the empty file. This is done by creating a folder (called bucket in MinIO) with the name `test` and uploading an empty byte stream with the file name `emptyFile.txt`. If you are familiar with Java, feel welcome to make changes before executing.

The terminal should state, that the build was successful.
