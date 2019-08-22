package services

import (
	"encoding/json"

	"github.com/iqbalnurhakim/bitmovin-go/bitmovin"
	"github.com/iqbalnurhakim/bitmovin-go/models"
)

type HLSManifestService struct {
	RestService *RestService
}

const (
	HLSManifestEndpoint string = "encoding/manifests/hls"
)

func NewHLSManifestService(bitmovin *bitmovin.Bitmovin) *HLSManifestService {
	r := NewRestService(bitmovin)

	return &HLSManifestService{RestService: r}
}

func (s *HLSManifestService) Create(a *models.HLSManifest) (*models.HLSManifestResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}

	o, err := s.RestService.Create(HLSManifestEndpoint, b)
	if err != nil {
		return nil, err
	}
	var r models.HLSManifestResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) Retrieve(id string) (*models.HLSManifestResponse, error) {
	path := HLSManifestEndpoint + "/" + id
	o, err := s.RestService.Retrieve(path)
	if err != nil {
		return nil, err
	}
	var r models.HLSManifestResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) Delete(id string) (*models.HLSManifestResponse, error) {
	path := HLSManifestEndpoint + "/" + id
	o, err := s.RestService.Delete(path)
	if err != nil {
		return nil, err
	}
	var r models.HLSManifestResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) AddMediaInfo(manifestID string, a *models.MediaInfo) (*models.MediaInfoResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}
	path := HLSManifestEndpoint + "/" + manifestID + "/media"
	o, err := s.RestService.Create(path, b)
	if err != nil {
		return nil, err
	}
	var r models.MediaInfoResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) AddStreamInfo(manifestID string, a *models.StreamInfo) (*models.StreamInfoResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}
	path := HLSManifestEndpoint + "/" + manifestID + "/streams"
	o, err := s.RestService.Create(path, b)
	if err != nil {
		return nil, err
	}
	var r models.StreamInfoResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) AddCustomTagToStreamInfo(manifestID string, streamInfoID string, a *models.CustomTag) (*models.CustomTagResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}
	path := HLSManifestEndpoint + "/" + manifestID + "/streams/" + streamInfoID + "/custom-tags"
	o, err := s.RestService.Create(path, b)
	if err != nil {
		return nil, err
	}
	var r models.CustomTagResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) AddIFramePlaylistToStreamInfo(manifestID string, streamInfoID string, a *models.IFramePlaylist) (*models.IFramePlaylistResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}
	path := HLSManifestEndpoint + "/" + manifestID + "/streams/" + streamInfoID + "/iframe"
	o, err := s.RestService.Create(path, b)
	if err != nil {
		return nil, err
	}
	var r models.IFramePlaylistResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) AddCustomTagToMediaInfo(manifestID string, mediaInfoID string, a *models.CustomTag) (*models.CustomTagResponse, error) {
	b, err := json.Marshal(*a)
	if err != nil {
		return nil, err
	}
	path := HLSManifestEndpoint + "/" + manifestID + "/media/" + mediaInfoID + "/custom-tags"
	o, err := s.RestService.Create(path, b)
	if err != nil {
		return nil, err
	}
	var r models.CustomTagResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) Start(manifestID string) (*models.StartStopResponse, error) {
	path := HLSManifestEndpoint + "/" + manifestID + "/start"
	o, err := s.RestService.Create(path, nil)
	if err != nil {
		return nil, err
	}
	var r models.StartStopResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *HLSManifestService) RetrieveStatus(manifestID string) (*models.StatusResponse, error) {
	path := HLSManifestEndpoint + "/" + manifestID + "/status"
	o, err := s.RestService.Retrieve(path)
	if err != nil {
		return nil, err
	}
	var r models.StatusResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
