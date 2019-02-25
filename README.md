# DPS - Dataset pipeline server

[![Build Status](https://travis-ci.com/rafaelcn/dataset-pipeline-server.svg?token=agoYobPmasJqPwFp7s8p&branch=master)](https://travis-ci.com/rafaelcn/dataset-pipeline-server)

This is the tool which serves a HTTP service that accepts json file uploads and
stores it on the specified folder, currently being uploads. This tool also does
save every data sent to the server to a sqlite database which can then be
recovered for easy inspection.
