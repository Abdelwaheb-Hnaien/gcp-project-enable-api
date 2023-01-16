package helpers

import(
  "google.golang.org/api/serviceusage/v1"
  osconfig "google.golang.org/api/osconfig/v1"
  "context"
  "fmt"
  "log"
)

func GetApiStatus(servicesService *serviceusage.ServicesService, name string) (status string)  {
  // ctx := context.Background()
  // serviceusageService, _ := serviceusage.NewService(ctx)
  // servicesService        := serviceusage.NewServicesService(serviceusageService)
  ServicesGetCall        := servicesService.Get(name)
  resp, err              := ServicesGetCall.Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  return resp.State

}
func EnableApi(servicesService *serviceusage.ServicesService, api string) (success bool)  {
  // ctx := context.Background()
  // serviceusageService, _ := serviceusage.NewService(ctx)
  // servicesService        := serviceusage.NewServicesService(serviceusageService)
  success = true
  var enableServiceRequest serviceusage.EnableServiceRequest
  ServicesGetCall        := servicesService.Enable(api,&enableServiceRequest)
  _, err              := ServicesGetCall.Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
    success = false
  }
  //print(resp.State)
  return success

}

func ListPatchDeployments(parent string) (PatchDeployment []*osconfig.PatchDeployment){
  ctx := context.Background()
  osconfigService, err := osconfig.NewService(ctx)
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  projectsPatchDeploymentsService := osconfig.NewProjectsPatchDeploymentsService(osconfigService)
  ProjectsPatchDeploymentsListCall := projectsPatchDeploymentsService.List(parent)

  resp, err := ProjectsPatchDeploymentsListCall.Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  return(resp.PatchDeployments)
}

func GetPatchDeployment(batchType string, projectId string, projectNumber int64) (*osconfig.PatchDeployment, error) {
  ctx := context.Background()
  osconfigService, err := osconfig.NewService(ctx)
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  projectsPatchDeploymentsService := osconfig.NewProjectsPatchDeploymentsService(osconfigService) // return *ProjectsPatchDeploymentsService
  PatchDeploymentId := fmt.Sprintf("projects/%v/patchDeployments/%s-%s",projectNumber, projectId, batchType)
  ProjectsPatchDeploymentsGetCall := projectsPatchDeploymentsService.Get(PatchDeploymentId)
  p, e := ProjectsPatchDeploymentsGetCall.Do()
  return p, e
}

func PreparePatchDeployment(t string, projectNumber string, projectID string ) (*osconfig.PatchDeployment) {
  ctx := context.Background()
  osconfigService, err := osconfig.NewService(ctx)
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  projectsPatchDeploymentsService := osconfig.NewProjectsPatchDeploymentsService(osconfigService) // return *ProjectsPatchDeploymentsService
  PatchDeploymentId := fmt.Sprintf("projects/%s/patchDeployments/%s-%s",projectNumber, projectID, t)
  ProjectsPatchDeploymentsGetCall := projectsPatchDeploymentsService.Get(PatchDeploymentId)
  p, e := ProjectsPatchDeploymentsGetCall.Do()
  if e != nil {
    log.Fatalf("Something is wrong ->  %v", e)
  }
  return p
}

func CreatePatchDeployment(parent string, patchDeployment *osconfig.PatchDeployment, patchDeploymentId string) (P *osconfig.PatchDeployment){
  ctx := context.Background()
  osconfigService, err := osconfig.NewService(ctx)
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  projectsPatchDeploymentsService := osconfig.NewProjectsPatchDeploymentsService(osconfigService)
  ProjectsPatchDeploymentsCreateCall := projectsPatchDeploymentsService.Create(parent, patchDeployment)
  resp, err := ProjectsPatchDeploymentsCreateCall.PatchDeploymentId(patchDeploymentId).Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  return resp
}

func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//`projects/840223767927/services/serviceusage.googleapis.com`
