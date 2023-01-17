package helpers

import(
  "google.golang.org/api/serviceusage/v1"
  "log"
)

func GetApiStatus(servicesService *serviceusage.ServicesService, name string) (status string)  {
  ServicesGetCall        := servicesService.Get(name)
  resp, err              := ServicesGetCall.Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
  }
  return resp.State
}

func EnableApi(servicesService *serviceusage.ServicesService, api string) (success bool)  {
  success = true
  var enableServiceRequest serviceusage.EnableServiceRequest
  ServicesGetCall        := servicesService.Enable(api,&enableServiceRequest)
  _, err              := ServicesGetCall.Do()
  if err != nil {
    log.Fatalf("Something is wrong ->  %v", err)
    success = false
  }
  return success
}
