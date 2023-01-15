package main

import(
  crm "google.golang.org/api/cloudresourcemanager/v1"
  "google.golang.org/api/googleapi"
  "google.golang.org/api/serviceusage/v1"
  "os_config_w/helpers"
  "context"
  "fmt"
  "strings"
)
var (
  API = "recommender.googleapis.com"
)

func main()  {

  ctx := context.Background()
  crmService, err  := crm.NewService(ctx)
  Projects         := make([]*crm.Project, 0)
  if err != nil {
    fmt.Errorf("Can't establish cloudresourcemanagerService service %v", err)
  }
  projectsService := crm.NewProjectsService(crmService)
  projectsListCall := projectsService.List()

  serviceusageService, _ := serviceusage.NewService(ctx)
  servicesService        := serviceusage.NewServicesService(serviceusageService)

  listProjectsResponse, err := projectsListCall.Do()
  if err != nil {
    fmt.Errorf("Something is wrong %v", err)
  }

  for next := true; next; next=listProjectsResponse.NextPageToken != "" {
    Projects = append(Projects, listProjectsResponse.Projects...)
    listProjectsResponse, err = projectsListCall.Do(googleapi.QueryParameter("page_token", listProjectsResponse.NextPageToken))
    if err != nil {
      fmt.Errorf("Something is wrong %v", err)
    }
  }


  for index, project := range Projects {
    if project.Parent != nil && strings.Contains(project.ProjectId,"some-prefix")   {
      url_api := fmt.Sprintf("projects/%v/services/%v",project.ProjectNumber, API)
      status  := helpers.GetApiStatus(servicesService, url_api)

      if status != "ENABLED" {
        msg := fmt.Sprintf(" --> %v : Enabling %v.. ", project.ProjectId, API)
        fmt.Println(index, msg)
        opearationSuccess  := helpers.EnableApi(servicesService, url_api)
        if opearationSuccess {
          fmt.Println("Success")
        }
      } else {
        msg := fmt.Sprintf(" --> %v : Skip, %v is already enabled", project.ProjectId, API)
        fmt.Println(index, msg)
      }
    }
  }
}
