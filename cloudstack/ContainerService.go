package cloudstack

import (
	"encoding/json"
	"net/url"
)

// Cluster api block.
// listClusters api block.
type ListCsClustersParams struct {
	p map[string]interface{}
}

func (p *ListCsClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterIds"]; found {
		u.Set("clusterIds", v.(string))
	}
	if v, found := p.p["clusterName"]; found {
		u.Set("clusterName", v.(string))
	}
	return u
}

func (p *ListCsClustersParams) SetClusterIds(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterIds"] = v
	return
}

func (p *ListCsClustersParams) SetClusterName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterName"] = v
	return
}
func (s *ContainerService) NewListClusterParams() *ListCsClustersParams {
	p := &ListCsClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *ContainerService) ListCluseters(p *ListCsClustersParams) (*ListCsClusetersResponse, error) {
	resp, err := s.cs.newRequest("listClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCsClusetersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCsClusetersResponse struct {
	JobId    string       `json:"jobId"`
	Count    int          `json:"count"`
	Clusters []*CsCluster `json:"clusterList"`
}

type CsCluster struct {
	ClusterId       string `json:"clusterId,omitempty"`
	ClusterName     string `json:"clusterName,omitempty"`
	OrchestrateType string `json:"orchestrateType,omitempty"`
	ClusterStatus   string `json:"clusterStatus,omitempty"`
	// HaEnabled               string `json:"haEnabled,omitempty"`
	AutoscalingGroupName    string `json:"autoscalingGroupName,omitempty"`
	LaunchConfigurationName string `json:"launchConfigurationName,omitempty"`
	ZoneId                  string `json:"zoneId,omitempty"`
}

// createClusters api block.

// ** Instance api block.
// listInstances api block.
type ListCsInstancesParams struct {
	p map[string]interface{}
}

func (p *ListCsInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterIds"]; found {
		u.Set("clusterIds", v.(string))
	}
	if v, found := p.p["clusterName"]; found {
		u.Set("clusterName", v.(string))
	}
	return u
}

func (p *ListCsInstancesParams) SetClusterIds(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterIds"] = v
	return
}

func (p *ListCsInstancesParams) SetClusterName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterName"] = v
	return
}

func (s *ContainerService) NewListInstancesParams() *ListCsInstancesParams {
	p := &ListCsInstancesParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *ContainerService) ListInstances(p *ListCsInstancesParams) (*ListCsInstancesResponse, error) {
	resp, err := s.cs.newRequest("listInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCsInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCsInstancesResponse struct {
	JobId     string        `json:"jobId"`
	Count     int           `json:"count"`
	Instances []*CsInstance `json:"instanceList"`
}

type CsInstance struct {
	ClusterId           string `json:"clusterId,omitempty"`
	ClusterName         string `json:"clusterName,omitempty"`
	OrchestrateType     string `json:"orchestrateType,omitempty"`
	ClusterStatus       string `json:"clusterStatus,omitempty"`
	InstanceId          string `json:"instanceId,omitempty"`
	InstanceName        string `json:"instanceName,omitempty"`
	InstanceType        string `json:"instanceType,omitempty"`
	InstanceStatus      string `json:"instanceStatus,omitempty"`
	CreationDt          string `json:"creationDt,omitempty"`
	UpdateDt            string `json:"updateDt,omitempty"`
	HostName            string `json:"hostName,omitempty"`
	OstypeName          string `json:"ostypeName,omitempty"`
	ServiceofferingId   string `json:"serviceofferingId,omitempty"`
	ServiceofferingName string `json:"serviceofferingName,omitempty"`
	ZoneId              string `json:"zoneId,omitempty"`
	ZoneName            string `json:"zoneName,omitempty"`
	TemplateId          string `json:"templateId,omitempty"`
	TemplateName        string `json:"templateName,omitempty"`
	CpuNum              string `json:"cpuNum,omitempty"`
	MemorySize          string `json:"memorySize,omitempty"`
	PrivateIpAddr       string `privateIpAddr:"zoneId,omitempty"`
}
