# Script to enable API for GCP projects

## What is it about ?
If you are working on Google Cloud Platform, you may want to enable an API for all projects you maintain for a particular reason (security, governance, etc.), this script does it for you.

## What does it do ?
The script uses the application default credentials to list all projects that has access to, then it will enable the API you define in the var bloc for each project. The script will skip projects in which the API is already enabled.

## How to run the script ?
1. Provide credentials for Application Default Credentials.
```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/sa-json.key"
```
More information [here](https://cloud.google.com/docs/authentication/provide-credentials-adc).  
Make sure the service account has at least the following permissions on the organization level:
  - resourcemanager.projects.list
  - serviceusage.services.enable


2. Define the API you want to enable:
```go
...
var (
    API = "recommender.googleapis.com"
)
...
```

3. Initialize the module and install dependencies
```bash
go mod init
go mod tidy
```

4. Run the script:
```bash
go run enable_api_v1.go
```

**Notice** : In this script, I tried to skip some projects using the following condition :
```go
if project.Parent != nil && strings.Contains(project.ProjectId,"some-prefix")
```
`project.Parent != nil` means make sure to consider only projects within the organization.
the other condition is to consider only project ids which starts with "some-prefix".

This is optional, if you want to enable the API for all projects, comment line 45 and 60 before running the script.
