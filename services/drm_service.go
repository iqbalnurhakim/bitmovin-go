package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iqbalnurhakim/bitmovin-go/bitmovin"
	"github.com/iqbalnurhakim/bitmovin-go/models"
	"strings"
)

type DrmService struct {
	RestService *RestService
}

const (
	Fmp4DrmEndpoint          string = "encoding/encodings/{encoding_id}/muxings/fmp4/{fmp4_id}/drm"
	TsDrmEndpoint            string = "encoding/encodings/{encoding_id}/muxings/ts/{ts_id}/drm"
	ProgressiveTsDrmEndpoint string = "encoding/encodings/{encoding_id}/muxings/progressive-ts/{progressive_ts_id}/drm"
	Mp4DrmEndpoint           string = "encoding/encodings/{encoding_id}/muxings/mp4/{mp4_id}/drm"
)

func NewDrmService(bitmovin *bitmovin.Bitmovin) *DrmService {
	return &DrmService{RestService: NewRestService(bitmovin)}
}

func (s *DrmService) CreateFmp4Drm(encodingId string, fmp4MuxingId string, drm interface{}) (interface{}, error) {

	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{fmp4_id}", fmp4MuxingId,
	)
	requestUrl := replacer.Replace(Fmp4DrmEndpoint)

	switch v := drm.(type) {
	case models.WidevineDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		enpointUrl := requestUrl + "/widevine"
		response, err := s.RestService.Create(enpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.WidevineDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	case models.PlayReadyDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		endpointUrl := requestUrl + "/playready"
		response, err := s.RestService.Create(endpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.PlayReadyDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	case models.FairPlayDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		endpointUrl := requestUrl + "/fairplay"
		response, err := s.RestService.Create(endpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.FairPlayDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	case models.CencDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		endpointUrl := requestUrl + "/cenc"
		response, err := s.RestService.Create(endpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.CencDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	default:
		err := fmt.Sprintf("FMP4 DRM type %T is not supported!\n", v)
		return nil, errors.New(err)
	}
}

func (s *DrmService) CreateTsDrm(encodingId string, tsMuxingId string, drm interface{}) (interface{}, error) {

	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{ts_id}", tsMuxingId,
	)
	requestUrl := replacer.Replace(TsDrmEndpoint)

	switch v := drm.(type) {
	case models.FairPlayDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		enpointUrl := requestUrl + "/fairplay"
		response, err := s.RestService.Create(enpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.FairPlayDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	default:
		err := fmt.Sprintf("TS DRM type %T is not supported!\n", v)
		return nil, errors.New(err)
	}
}

func (s *DrmService) CreateProgressiveTsDrm(encodingId string, progressiveTsMuxingId string, drm interface{}) (interface{}, error) {

	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{progressive_ts_id}", progressiveTsMuxingId,
	)
	requestUrl := replacer.Replace(ProgressiveTsDrmEndpoint)

	switch v := drm.(type) {
	case models.FairPlayDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		enpointUrl := requestUrl + "/fairplay"
		response, err := s.RestService.Create(enpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.FairPlayDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	default:
		err := fmt.Sprintf("Progressive TS DRM type %T is not supported!\n", v)
		return nil, errors.New(err)
	}
}

func (s *DrmService) CreateMp4Drm(encodingId string, mp4MuxingId string, drm interface{}) (interface{}, error) {

	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{mp4_id}", mp4MuxingId,
	)
	requestUrl := replacer.Replace(Mp4DrmEndpoint)

	switch v := drm.(type) {
	case models.WidevineDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		enpointUrl := requestUrl + "/widevine"
		response, err := s.RestService.Create(enpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.WidevineDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	case models.PlayReadyDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		endpointUrl := requestUrl + "/playready"
		response, err := s.RestService.Create(endpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.PlayReadyDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	case models.CencDrm:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		endpointUrl := requestUrl + "/cenc"
		response, err := s.RestService.Create(endpointUrl, b)
		if err != nil {
			return nil, err
		}

		var result models.CencDrmResponse
		err = json.Unmarshal(response, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil

	default:
		err := fmt.Sprintf("FMP4 DRM type %T is not supported!\n", v)
		return nil, errors.New(err)
	}
}

func (s *DrmService) ListMp4CencDrm(encodingId string, mp4MuxingId string, offset int64, limit int64) (*models.CencDrmListResponse, error) {
	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{mp4_id}", mp4MuxingId,
	)
	requestUrl := replacer.Replace(Mp4DrmEndpoint)
	endpointUrl := requestUrl + "/cenc"
	o, err := s.RestService.List(endpointUrl, offset, limit)
	if err != nil {
		return nil, err
	}
	var r models.CencDrmListResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DrmService) ListFmp4CencDrm(encodingId string, fmp4MuxingId string, offset int64, limit int64) (*models.CencDrmListResponse, error) {
	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{fmp4_id}", fmp4MuxingId,
	)
	requestUrl := replacer.Replace(Fmp4DrmEndpoint)
	endpointUrl := requestUrl + "/cenc"
	o, err := s.RestService.List(endpointUrl, offset, limit)
	if err != nil {
		return nil, err
	}
	var r models.CencDrmListResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DrmService) ListTsFairPlayDrm(encodingId string, tsMuxingId string, offset int64, limit int64) (*models.FairPlayDrmListResponse, error) {
	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{ts_id}", tsMuxingId,
	)
	requestUrl := replacer.Replace(TsDrmEndpoint)
	endpointUrl := requestUrl + "/fairplay"
	o, err := s.RestService.List(endpointUrl, offset, limit)
	if err != nil {
		return nil, err
	}
	var r models.FairPlayDrmListResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *DrmService) ListProgressiveTsFairPlayDrm(encodingId string, progressiveTsMuxingId string, offset int64, limit int64) (*models.FairPlayDrmListResponse, error) {
	replacer := strings.NewReplacer(
		"{encoding_id}", encodingId,
		"{progressive_ts_id}", progressiveTsMuxingId,
	)
	requestUrl := replacer.Replace(ProgressiveTsDrmEndpoint)
	endpointUrl := requestUrl + "/fairplay"
	o, err := s.RestService.List(endpointUrl, offset, limit)
	if err != nil {
		return nil, err
	}
	var r models.FairPlayDrmListResponse
	err = json.Unmarshal(o, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
