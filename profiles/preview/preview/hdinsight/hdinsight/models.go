// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package hdinsight

import original "github.com/Azure/azure-sdk-for-go/services/preview/hdinsight/2018-11-01-preview/hdinsight"

type ApplicationState = original.ApplicationState

const (
	ACCEPTED  ApplicationState = original.ACCEPTED
	FAILED    ApplicationState = original.FAILED
	FINISHED  ApplicationState = original.FINISHED
	FINISHING ApplicationState = original.FINISHING
	KILLED    ApplicationState = original.KILLED
	NEW       ApplicationState = original.NEW
	NEWSAVING ApplicationState = original.NEWSAVING
	RUNNING   ApplicationState = original.RUNNING
	SUBMITTED ApplicationState = original.SUBMITTED
)

type SessionJobKind = original.SessionJobKind

const (
	Pyspark SessionJobKind = original.Pyspark
	Spark   SessionJobKind = original.Spark
	Sparkr  SessionJobKind = original.Sparkr
	SQL     SessionJobKind = original.SQL
)

type AppState = original.AppState
type BaseClient = original.BaseClient
type JobClient = original.JobClient
type JobDetailRootJSONObject = original.JobDetailRootJSONObject
type JobID = original.JobID
type JobListJSONObject = original.JobListJSONObject
type JobOperationsErrorResponse = original.JobOperationsErrorResponse
type JobSubmissionJSONResponse = original.JobSubmissionJSONResponse
type ListJobListJSONObject = original.ListJobListJSONObject
type Profile = original.Profile
type SparkBatchJob = original.SparkBatchJob
type SparkBatchJobCollection = original.SparkBatchJobCollection
type SparkBatchJobRequest = original.SparkBatchJobRequest
type SparkJobDeletedResult = original.SparkJobDeletedResult
type SparkJobLog = original.SparkJobLog
type SparkJobState = original.SparkJobState
type SparkSessionCollection = original.SparkSessionCollection
type SparkSessionJob = original.SparkSessionJob
type SparkSessionJobRequest = original.SparkSessionJobRequest
type SparkStatement = original.SparkStatement
type SparkStatementCancellationResult = original.SparkStatementCancellationResult
type SparkStatementCollection = original.SparkStatementCollection
type SparkStatementOutput = original.SparkStatementOutput
type SparkStatementRequest = original.SparkStatementRequest
type Status = original.Status
type Userargs = original.Userargs

func New(endpoint string, userName string) BaseClient {
	return original.New(endpoint, userName)
}
func NewJobClient(endpoint string, userName string) JobClient {
	return original.NewJobClient(endpoint, userName)
}
func NewWithoutDefaults(endpoint string, userName string) BaseClient {
	return original.NewWithoutDefaults(endpoint, userName)
}
func PossibleApplicationStateValues() []ApplicationState {
	return original.PossibleApplicationStateValues()
}
func PossibleSessionJobKindValues() []SessionJobKind {
	return original.PossibleSessionJobKindValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
