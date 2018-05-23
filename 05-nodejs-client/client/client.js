/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

var PROTO_PATH = __dirname + '/../proto/service.proto';

var async = require('async');
var fs = require('fs');
var parseArgs = require('minimist');
var path = require('path');
var _ = require('lodash');
var grpc = require('grpc');
var answerService = grpc.load(PROTO_PATH);

var client = new answerService.AnswerService('localhost:8000',
    grpc.credentials.createInsecure());

function runGetAnswer(callback) {
    var next = _.after(2, callback);
    function answerCallback(error, resp) {
        if (error) {
            console.log(error);
            return;
        }
        if (resp.name === '') {
            console.log('no answer found');
        } else {
            console.log('answer: ' + resp.answer);
        }
        next();
    }
    var question = "what is this?";
    console.log(question);
    client.answer({question: question}, answerCallback);
}

/**
 * Run all of the demos in order
 */
function main() {
    async.series([
        runGetAnswer,
    ]);
}

if (require.main === module) {
    main();
}

exports.runGetAnswer = runGetAnswer;
