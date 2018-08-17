![release-pre-alpha](https://rawgit.com/suite911/assets/master/shields/release-pre--alpha-red.svg)
![api-pre-alpha](https://rawgit.com/suite911/assets/master/shields/api-pre--alpha-red.svg)
![abi-pre-alpha](https://rawgit.com/suite911/assets/master/shields/abi-pre--alpha-red.svg)
[![Build Status](https://travis-ci.org/suite911/cloud911.svg?branch=master)](https://travis-ci.org/suite911/cloud911)
[![CC0-1.0](https://rawgit.com/suite911/assets/master/shields/license-cc0--1.0-efbfff.svg)](https://raw.githubusercontent.com/suite911/cloud911/master/LICENSE.txt)

# Cloud911&trade;

Cloud911&trade; is a fast and hackable cloud app framework for Google Go (golang).  It helps you make your app's backend and connect it to a very small built-in web server.

## Unencrypted API calls over HTTP
Cloud911&trade; is optimized for API calls over HTTP.  Payloads are not encrypted by TLS so encrypt any sensitive data yourself or use HTTPS.  Unencrypted API calls look like this: `http://example.com/api/your/custom/path`

## Encrypted API calls over HTTPS
Encrypting the API calls with TLS is trivial: simply use HTTPS for transport.  TLS-encrypted API calls look like this: `https://example.com/api/your/custom/path`

## Built-in web server
Cloud911&trade; includes a small built-in web server.  Visit a non-`/api` page in a web browser and get it served over HTTPS.
