package cloudstack

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CreateMpxLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *CreateMpxLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["loadbalanceroption"]; found {
		u.Set("loadbalanceroption", v.(string))
	}
	if v, found := p.p["serviceip"]; found {
		u.Set("serviceip", v.(string))
	}
	if v, found := p.p["serviceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("serviceport", vv)
	}
	if v, found := p.p["servicetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("servicetype", vv)
	}
	if v, found := p.p["healthchecktype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthchecktype", vv)
	}
	if v, found := p.p["healthcheckurl"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthcheckurl", vv)
	}

	return u
}

func (p *CreateMpxLoadBalancerParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v

	return
}

func (p *CreateMpxLoadBalancerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetLoadbalanceroption(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalanceroption"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetServiceip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceip"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetServiceport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceport"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetServicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicetype"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetHealthchecktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthchecktype"] = v
	return
}

func (p *CreateMpxLoadBalancerParams) SetHealthcheckurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthcheckurl"] = v
	return
}

func (s *MpxLoadBalancerService) NewCreateMpxLoadBalancerParams(zoneid string, name string, loadbalanceroption string, serviceip string, serviceport int, servicetype string, helthchecktype string, healthcheckurl string) *CreateMpxLoadBalancerParams {
	p := &CreateMpxLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	p.p["name"] = name
	p.p["loadbalanceroption"] = loadbalanceroption
	p.p["serviceip"] = serviceip
	p.p["serviceport"] = serviceport
	p.p["servicetype"] = servicetype
	p.p["helthchecktype"] = helthchecktype
	p.p["healthcheckurl"] = healthcheckurl
	return p
}

// Creates a MPX load balancer
func (s *MpxLoadBalancerService) CreateMpxLoadBalancer(p *CreateMpxLoadBalancerParams) (*CreateMpxLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateMpxLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateMpxLoadBalancerResponse struct {
	Zoneid             string `json:"zoneid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Loadbalacerid      string `json:"loadbalancerid,omitempty"`
	Name               string `json:"name,omitempty"`
	Loadbalanceroption string `json:"loadbalanceroption,omitempty"`
	Serviceip          string `json:"serviceip,omitempty"`
	Serviceport        string `json:"serviceport,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Healthchecktype    string `json:"healthchecktype,omitempty"`
	Healthcheckurl     string `json:"healthcheckurl,omitempty"`
}

type DeleteMpxLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteMpxLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}

	return u
}

func (p *DeleteMpxLoadBalancerParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}

func (s *MpxLoadBalancerService) NewDeleteMpxLoadBalancerParams(loadbalancerid string) *DeleteMpxLoadBalancerParams {
	p := &DeleteMpxLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["loadbalancerid"] = loadbalancerid
	return p
}

func (s *MpxLoadBalancerService) DeleteMpxLoadBalancer(p *DeleteMpxLoadBalancerParams) (*DeleteMpxLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteMpxLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteMpxLoadBalancerResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type UpdateMpxLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *UpdateMpxLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	if v, found := p.p["loadbalanceroption"]; found {
		u.Set("loadbalanceroption", v.(string))
	}
	if v, found := p.p["servicetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("servicetype", vv)
	}
	if v, found := p.p["healthchecktype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthchecktype", vv)
	}
	if v, found := p.p["healthcheckurl"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthcheckurl", vv)
	}
	return u
}
func (p *UpdateMpxLoadBalancerParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}
func (p *UpdateMpxLoadBalancerParams) SetLoadbalanceroption(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalanceroption"] = v
	return
}
func (p *UpdateMpxLoadBalancerParams) SetServicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicetype"] = v
	return
}
func (p *UpdateMpxLoadBalancerParams) SetHealthchecktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthchecktype"] = v
	return
}
func (p *UpdateMpxLoadBalancerParams) SetHealthcheckurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthcheckurl"] = v
	return
}
func (s *MpxLoadBalancerService) UpdateLoadBalancer(p *UpdateMpxLoadBalancerParams) (*UpdateMpxLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateMpxLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateMpxLoadBalancerResponse struct {
	Zoneid             string `json:"zoneid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Loadbalacerid      string `json:"loadbalancerid,omitempty"`
	Name               string `json:"name,omitempty"`
	Loadbalanceroption string `json:"loadbalanceroption,omitempty"`
	Serviceip          string `json:"serviceip,omitempty"`
	Serviceport        string `json:"serviceport,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Healthchecktype    string `json:"healthchecktype,omitempty"`
	Healthcheckurl     string `json:"healthcheckurl,omitempty"`
}

type ListMpxLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListMpxLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		println("p.p == nil임")
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["serviceip"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("serviceip", vv)
	}
	if v, found := p.p["memid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memid", vv)
	}
	return u
}
func (p *ListMpxLoadBalancersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}
func (p *ListMpxLoadBalancersParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}
func (p *ListMpxLoadBalancersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}
func (p *ListMpxLoadBalancersParams) SetServiceip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceip"] = v
	return
}
func (p *ListMpxLoadBalancersParams) SetMemid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["memid"] = v
	return
}
func (s *MpxLoadBalancerService) NewListMpxLoadBalancerParams() *ListMpxLoadBalancersParams {
	p := &ListMpxLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *MpxLoadBalancerService) ListMpxLoadBalancers(p *ListMpxLoadBalancersParams) (*ListMpxLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListMpxLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListMpxLoadBalancersResponse struct {
	Count            int                `json:"count"`
	MpxLoadBalancers []*MpxLoadBalancer `json:"loadbalancer"`
}
type MpxLoadBalancer struct {
	Zoneid             string `json:"zoneid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Loadbalacerid      int    `json:"loadbalancerid"`
	Name               string `json:"name"`
	Loadbalanceroption string `json:"loadbalanceroption,omitempty"`
	Serviceip          string `json:"serviceip,omitempty"`
	Serviceport        string `json:"serviceport,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Healthchecktype    string `json:"healthchecktype,omitempty"`
	Healthcheckurl     string `json:"healthcheckurl,omitempty"`
	Certificatename    string `json:"certificatename,omitempty"`
	CipherGroupName    string `json:"cipherGroupName,omitempty"`
	ClientIpYn         string `json:"clientIpYn,omitempty"`
	Networkid          string `json:"networkid,omitempty"`
	Requestsrate       string `json:"requestsrate,omitempty"`
	Sslv2              string `json:"sslv2,omitempty"`
	Sslv3              string `json:"sslv3,omitempty"`
	State              string `json:"state,omitempty"`
	Tag                string `json:"tag,omitempty"`
	Tlsv1              string `json:"tlsv1,omitempty"`
	Tlsv11             string `json:"tlsv11,omitempty"`
	Tlsv12             string `json:"tlsv12,omitempty"`
}

type CheckMpxLoadBalancerNameParams struct {
	p map[string]interface{}
}

func (p *CheckMpxLoadBalancerNameParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}
func (p *CheckMpxLoadBalancerNameParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}
func (p *CheckMpxLoadBalancerNameParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

func (s *MpxLoadBalancerService) NewCheckMpxLoadBalancerNameParams(name string, zoneid string) *CheckMpxLoadBalancerNameParams {
	p := &CheckMpxLoadBalancerNameParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

func (s *MpxLoadBalancerService) CheckMpxLoadBalancerName(p *CheckMpxLoadBalancerNameParams) (*CheckMpxLoadBalancerNameResponse, error) {
	resp, err := s.cs.newRequest("checkLoadBalancerName", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CheckMpxLoadBalancerNameResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CheckMpxLoadBalancerNameResponse struct {
	Success string `json:"success,omitempty"`
	Text    string `json:"text,omitempty"`
}

// Mpx LoadBalancer Service Resource 관리 API
type AddMpxLoadBlancerWebServerParams struct {
	p map[string]interface{}
}

func (p *AddMpxLoadBlancerWebServerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		u.Set("publicport", v.(string))
	}
	return u
}
func (p *AddMpxLoadBlancerWebServerParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}
func (p *AddMpxLoadBlancerWebServerParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}
func (p *AddMpxLoadBlancerWebServerParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}
func (p *AddMpxLoadBlancerWebServerParams) SetPublicport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (s *MpxLoadBalancerService) NewAddMpxLoadBlancerWebServerParams(loadbalancerid int, virtualmachineid string, ipaddress string, publicport int) *AddMpxLoadBlancerWebServerParams {
	p := &AddMpxLoadBlancerWebServerParams{}
	p.p = make(map[string]interface{})
	p.p["loadbalancerid"] = loadbalancerid
	p.p["virtualmachineid"] = virtualmachineid
	p.p["ipaddress"] = ipaddress
	p.p["publicport"] = publicport
	return p
}

func (s *MpxLoadBalancerService) AddMpxLoadBlancerWebServer(p *AddMpxLoadBlancerWebServerParams) (*AddMpxLoadBlancerWebServerResponse, error) {
	resp, err := s.cs.newRequest("addLoadBlancerWebServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddMpxLoadBlancerWebServerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddMpxLoadBlancerWebServerResponse struct {
	Serviceid        int    `json:"serviceid,omitempty"`
	Loadbalancerid   string `json:"loadbalancerid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Publicport       int    `json:"publicport,omitempty"`
}

type RemoveMpxLoadBlancerWebServerParams struct {
	p map[string]interface{}
}

func (p *RemoveMpxLoadBlancerWebServerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["serviceid"]; found {
		u.Set("serviceid", v.(string))
	}
	return u
}
func (p *RemoveMpxLoadBlancerWebServerParams) SetServiceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceid"] = v
	return
}

func (s *MpxLoadBalancerService) NewRemoveMpxLoadBlancerWebServerParams(serviceid string) *RemoveMpxLoadBlancerWebServerParams {
	p := &RemoveMpxLoadBlancerWebServerParams{}
	p.p = make(map[string]interface{})
	p.p["serviceid"] = serviceid
	return p
}

func (s *MpxLoadBalancerService) RemoveLoadBalancerWebServer(p *RemoveMpxLoadBlancerWebServerParams) (*RemoveMpxLoadBlancerWebServerResponse, error) {
	resp, err := s.cs.newRequest("removeLoadBalancerWebServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveMpxLoadBlancerWebServerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveMpxLoadBlancerWebServerResponse struct {
	Success     bool   `json:"success"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListMpxLoadBlancerWebServersParams struct {
	p map[string]interface{}
}

func (p *ListMpxLoadBlancerWebServersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	return u
}
func (p *ListMpxLoadBlancerWebServersParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}

func (s *MpxLoadBalancerService) NewListMpxLoadBlancerWebServersParams(loadbalancerid string) *ListMpxLoadBlancerWebServersParams {
	p := &ListMpxLoadBlancerWebServersParams{}
	p.p = make(map[string]interface{})
	p.p["loadbalancerid"] = loadbalancerid
	return p
}

// func (s *MpxLoadBalancerService) NewListMpxLoadBlancerWebServersParams(lbname string) *ListMpxLoadBlancerWebServersParams {
// 	p := &ListMpxLoadBlancerWebServersParams{}
// 	p.p = make(map[string]interface{})
// 	p.p["loadbalancerid"] = loadbalancerid
// 	return p
// }
func (s *MpxLoadBalancerService) ListMpxLoadBlancerWebServers(p *ListMpxLoadBlancerWebServersParams) (*ListMpxLoadBlancerWebServersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerWebServers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListMpxLoadBlancerWebServersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListMpxLoadBlancerWebServersResponse struct {
	Count                    int                         `json:"count"`
	MpxLoadbalancerWebserver []*MpxLoadbalancerWebserver `json:"loadbalancerwebserver"`
}
type MpxLoadbalancerWebserver struct {
	Loadbalancerid     int    `json:"loadbalancerid,omitempty"`
	Serviceid          int    `json:"serviceid,omitempty"`
	Virtualmachineid   string `json:"virtualmachineid,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Publicport         string `json:"publicport,omitempty"`
	Cursrvrconnections string `json:"cursrvrconnections,omitempty"`
	State              string `json:"state,omitempty"`
	Throughputrate     int    `json:"throughputrate ,omitempty"`
	Avgsvrttfb         int    `json:"avgsvrttfb ,omitempty"`
	Requestsrate       int    `json:"requestsrate ,omitempty"`
}

type CreateTagParams struct {
	p map[string]interface{}
}

func (p *CreateTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}
func (p *CreateTagParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}
func (p *CreateTagParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
	return
}
func (s *MpxLoadBalancerService) CreateTag(p *CreateTagParams) (*CreateTagResponse, error) {
	resp, err := s.cs.newRequest("createTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTagResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateTagResponse struct {
	Success     bool   `json:success,omitempth`
	Displaytext string `json:displaytext,omitempth`
}

type DeleteTagParams struct {
	p map[string]interface{}
}

func (p *DeleteTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["loadbalancerid"]; found {
		u.Set("loadbalancerid", v.(string))
	}
	return u
}
func (p *DeleteTagParams) SetLoadbalancerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerid"] = v
	return
}
func (s *MpxLoadBalancerService) DeleteTag(p *DeleteTagParams) (*DeleteTagResponse, error) {
	resp, err := s.cs.newRequest("deleteTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTagResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteTagResponse struct {
	Success     bool   `json:success,omitempth`
	Displaytext string `json:displaytext,omitempth`
}

// Certificate API
type CreateCertificateParams struct {
	p map[string]interface{}
}

func (p *CreateCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cert"]; found {
		u.Set("cert", v.(string))
	}
	if v, found := p.p["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := p.p["interCert"]; found {
		u.Set("interCert", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}
func (p *CreateCertificateParams) SetCert(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cert"] = v
	return
}
func (p *CreateCertificateParams) SetKey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["key"] = v
	return
}
func (p *CreateCertificateParams) SetInterCert(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intercert"] = v
	return
}
func (p *CreateCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}
func (p *CreateCertificateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}
func (s *MpxLoadBalancerService) CreateCertificate(p *CreateCertificateParams) (*CreateCertificateResponse, error) {
	resp, err := s.cs.newRequest("createCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateCertificateResponse struct {
	Success bool `json:success,omitempth`
}

type DeleteCertificateParams struct {
	p map[string]interface{}
}

func (p *DeleteCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}
func (p *DeleteCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}
func (s *MpxLoadBalancerService) DeleteCertificate(p *DeleteCertificateParams) (*DeleteCertificateResponse, error) {
	resp, err := s.cs.newRequest("deleteCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteCertificateResponse struct {
	Success bool `json:success,omitempth`
}

type ManageCertificateParams struct {
	p map[string]interface{}
}

func (p *ManageCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["lbname"]; found {
		u.Set("lbname", v.(string))
	}
	if v, found := p.p["bindingyn"]; found {
		u.Set("bindingyn", v.(string))
	}
	return u
}
func (p *ManageCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}
func (p *ManageCertificateParams) SetLbname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbname"] = v
	return
}
func (p *ManageCertificateParams) SetBindingyn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bindingyn"] = v
	return
}
func (s *MpxLoadBalancerService) ManageCertificate(p *ManageCertificateParams) (*ManageCertificateResponse, error) {
	resp, err := s.cs.newRequest("deleteCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ManageCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ManageCertificateResponse struct {
	Success bool `json:success,omitempth`
}
