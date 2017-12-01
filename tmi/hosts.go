package tmi

import (
	"encoding/json"

	"github.com/julusian/go-twitch-helix/twitch"
)

// HostsList is a structure representing a list of channels who are hosting others
// https://tmi.twitch.tv/hosts?include_logins=1&target=###
type HostsList struct {
	Hosts []HostsEntry `json:"hosts,omitempty"`
}

// HostsEntry is an element of HostsList
// {"host_id":42142659,"target_id":12943124,"host_login":"kiyoshi_shinji","target_login":"ltzonda","host_display_name":"Kiyoshi_Shinji","target_display_name":"Ltzonda"}
type HostsEntry struct {
	HostID            int    `json:"host_id,omitempty"`
	HostLogin         string `json:"host_login,omitempty"`
	HostDisplayName   string `json:"host_display_name,omitempty"`
	HostPartnered     bool   `json:"host_partnered,omitempty"`
	TargetID          int    `json:"target_id,omitempty"`
	TargetLogin       string `json:"target_login,omitempty"`
	TargetDisplayName string `json:"target_display_name,omitempty"`
}

func BuildGetHosts(channel int) twitch.IRequest {
	return newTmiRequest("hosts").
		WithParamInt("include_logins", 1).
		WithParamInt("target", channel).
		Get()
}

func GetHosts(client *twitch.ApiClient, channel int) (*HostsList, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(BuildGetHosts(channel))
	if err != nil {
		return nil, nil, err
	}

	r, err := UnmarshalHosts(res)
	return r, rate, err
}

func UnmarshalHosts(res []byte) (*HostsList, error) {
	r := &HostsList{}
	err := json.Unmarshal(res, r)
	return r, err
}
