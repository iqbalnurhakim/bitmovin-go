package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iqbalnurhakim/bitmovin-go/bitmovin"
	"github.com/iqbalnurhakim/bitmovin-go/bitmovintypes"
	"github.com/iqbalnurhakim/bitmovin-go/models"
	"github.com/iqbalnurhakim/bitmovin-go/services"
)

type H265CodecConfigDefinition struct {
	width    *int64
	height   *int64
	crf      *float64
	streamId *string
	muxingId *string
}

const (
	apiKey             = "YOUR API KEY"
	inputHost          = "YOUR_HTTP_INPUT_HOST"
	inputFilePath      = "/path/to/your/input/file.mov"
	s3OutputAccessKey  = "YOUR_ACCESS_KEY"
	s3OutputSecretKey  = "YOUR_SECRET_KEY"
	s3OutputBucketName = "YOUR_BUCKET_NAME"
	outputBasePath     = "golang_hls_example"
)

func main() {
	// Creating Bitmovin object
	bitmovin := bitmovin.NewBitmovin(apiKey, "https://api.bitmovin.com/v1/", 5)

	videoEncodingProfiles := []*H265CodecConfigDefinition{
		{width: intToPtr(4096), height: intToPtr(2160), crf: floatToPtr(19.0)},
		{width: intToPtr(1920), height: intToPtr(1080), crf: floatToPtr(19.0)},
		{width: intToPtr(1280), height: intToPtr(720), crf: floatToPtr(19.0)},
	}

	// Creating the HTTP Input
	httpIS := services.NewHTTPInputService(bitmovin)
	httpInput := &models.HTTPInput{
		Host: stringToPtr(inputHost),
	}
	httpResp, err := httpIS.Create(httpInput)
	errorHandler(err)

	s3OS := services.NewS3OutputService(bitmovin)
	s3Output := &models.S3Output{
		AccessKey:  stringToPtr(s3OutputAccessKey),
		SecretKey:  stringToPtr(s3OutputSecretKey),
		BucketName: stringToPtr(s3OutputBucketName),
	}
	s3OutputResp, err := s3OS.Create(s3Output)
	errorHandler(err)

	t := time.Now()
	outputBasePath := outputBasePath + "/" + t.Format("2006-01-02-15-04-05")

	encodingS := services.NewEncodingService(bitmovin)
	encoding := &models.Encoding{
		Name:        stringToPtr("4K HDR10 example encoding"),
		CloudRegion: bitmovintypes.CloudRegionGoogleEuropeWest1,
	}
	encodingResp, err := encodingS.Create(encoding)
	errorHandler(err)

	videoInputStream := models.InputStream{
		InputID:       httpResp.Data.Result.ID,
		InputPath:     stringToPtr(inputFilePath),
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	audioInputStream := models.InputStream{
		InputID:       httpResp.Data.Result.ID,
		InputPath:     stringToPtr(inputFilePath),
		SelectionMode: bitmovintypes.SelectionModeAuto,
	}

	aclEntry := models.ACLItem{
		Permission: bitmovintypes.ACLPermissionPublicRead,
	}
	acl := []models.ACLItem{aclEntry}

	vis := []models.InputStream{videoInputStream}
	ais := []models.InputStream{audioInputStream}

	for _, codecConfigDefinition := range videoEncodingProfiles {
		h265S := services.NewH265CodecConfigurationService(bitmovin)

		colorConfig := models.ColorConfig{
			ColorTransfer:  bitmovintypes.ColorTransferSMPTE2084,
			ColorPrimaries: bitmovintypes.ColorPrimariesBT2020,
			ColorSpace:     bitmovintypes.ColorSpaceBT2020_NCL,
		}

		videoConfig := &models.H265CodecConfiguration{
			Name:         stringToPtr(fmt.Sprintf("HEVC_HDR10_%d", *codecConfigDefinition.height)),
			Description:  stringToPtr("HEVC_HDR_10"),
			FrameRate:    floatToPtr(25.0),
			CRF:          codecConfigDefinition.crf,
			Width:        codecConfigDefinition.width,
			Height:       codecConfigDefinition.height,
			Profile:      bitmovintypes.H265ProfileMain10,
			BAdapt:       bitmovintypes.BAdaptFull,
			RCLookahead:  intToPtr(40),
			RefFrames:    intToPtr(5),
			MotionSearch: bitmovintypes.MotionSearchStar,
			SubMe:        intToPtr(4),
			TUInterDepth: bitmovintypes.TUInterDepth3,
			TUIntraDepth: bitmovintypes.TUIntraDepth3,
			MaxCTUSize:   bitmovintypes.MaxCTUSize64,
			BFrames:      intToPtr(4),
			SAO:          boolToPtr(true),
			WeightPredictionOnPSlice:    boolToPtr(true),
			WeightPredictionOnBSlice:    boolToPtr(false),
			ColorConfig:                 colorConfig,
			MasterDisplay:               stringToPtr("G(8500,39850)B(6550,2300)R(35400,14600)WP(15635,16450)L(10000000,1)"),
			MaxContentLightLevel:        intToPtr(1000),
			MaxPictureAverageLightLevel: intToPtr(180),
			HDR:         boolToPtr(true),
			PixelFormat: bitmovintypes.PixelFormatYUV420P10LE,
		}

		videoConfigResp, err := h265S.Create(videoConfig)
		errorHandler(err)

		videoStream := &models.Stream{
			CodecConfigurationID: videoConfigResp.Data.Result.ID,
			InputStreams:         vis,
		}

		videoStreamResp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, videoStream)
		errorHandler(err)

		videoMuxingStream := models.StreamItem{
			StreamID: videoStreamResp.Data.Result.ID,
		}

		videoMuxingOutput := models.Output{
			OutputID:   s3OutputResp.Data.Result.ID,
			OutputPath: stringToPtr(fmt.Sprintf("%s/video/%d", outputBasePath, *codecConfigDefinition.height)),
			ACL:        acl,
		}

		videoMuxing := &models.FMP4Muxing{
			SegmentLength:   floatToPtr(4.0),
			SegmentNaming:   stringToPtr("seg_%number%.m4s"),
			InitSegmentName: stringToPtr("init.mp4"),
			Streams:         []models.StreamItem{videoMuxingStream},
			Outputs:         []models.Output{videoMuxingOutput},
		}

		videoMuxingResp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, videoMuxing)
		errorHandler(err)

		codecConfigDefinition.streamId = videoStreamResp.Data.Result.ID
		codecConfigDefinition.muxingId = videoMuxingResp.Data.Result.ID
	}

	aacS := services.NewAACCodecConfigurationService(bitmovin)
	aacConfig := &models.AACCodecConfiguration{
		Name:         stringToPtr("example_audio_codec_configuration"),
		Bitrate:      intToPtr(128000),
		SamplingRate: floatToPtr(48000.0),
	}
	aacResp, err := aacS.Create(aacConfig)
	errorHandler(err)

	audioStream := &models.Stream{
		CodecConfigurationID: aacResp.Data.Result.ID,
		InputStreams:         ais,
	}
	aacStreamResp, err := encodingS.AddStream(*encodingResp.Data.Result.ID, audioStream)
	errorHandler(err)

	audioMuxingStream := models.StreamItem{
		StreamID: aacStreamResp.Data.Result.ID,
	}

	audioMuxingOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(fmt.Sprintf("%s/audio", outputBasePath)),
		ACL:        acl,
	}

	audioMuxing := &models.FMP4Muxing{
		SegmentLength:   floatToPtr(4.0),
		SegmentNaming:   stringToPtr("seg_%number%.m4s"),
		InitSegmentName: stringToPtr("init.mp4"),
		Streams:         []models.StreamItem{audioMuxingStream},
		Outputs:         []models.Output{audioMuxingOutput},
	}

	audioMuxingResp, err := encodingS.AddFMP4Muxing(*encodingResp.Data.Result.ID, audioMuxing)
	errorHandler(err)

	_, err = encodingS.Start(*encodingResp.Data.Result.ID)
	errorHandler(err)

	var status string
	status = ""
	for status != "FINISHED" {
		time.Sleep(10 * time.Second)
		statusResp, err := encodingS.RetrieveStatus(*encodingResp.Data.Result.ID)
		if err != nil {
			fmt.Println("error in Encoding Status")
			fmt.Println(err)
			return
		}
		// Polling and Printing out the response
		fmt.Printf("%+v\n", statusResp)
		status = *statusResp.Data.Result.Status
		if status == "ERROR" {
			fmt.Println("error in Encoding Status")
			fmt.Printf("%+v\n", statusResp)
			return
		}
	}

	manifestOutput := models.Output{
		OutputID:   s3OutputResp.Data.Result.ID,
		OutputPath: stringToPtr(fmt.Sprintf("%s/manifest", outputBasePath)),
		ACL:        acl,
	}
	hlsManifest := &models.HLSManifest{
		ManifestName: stringToPtr("your_manifest_name.m3u8"),
		Outputs:      []models.Output{manifestOutput},
	}
	hlsService := services.NewHLSManifestService(bitmovin)
	hlsManifestResp, err := hlsService.Create(hlsManifest)
	errorHandler(err)

	audioMediaInfo := &models.MediaInfo{
		Type:            bitmovintypes.MediaTypeAudio,
		URI:             stringToPtr("audio.m3u8"),
		GroupID:         stringToPtr("audio_group"),
		Language:        stringToPtr("en"),
		Name:            stringToPtr("Rendition Description"),
		IsDefault:       boolToPtr(false),
		Autoselect:      boolToPtr(false),
		Forced:          boolToPtr(false),
		Characteristics: []string{"public.accessibility.describes-video"},
		SegmentPath:     stringToPtr("../audio"),
		EncodingID:      encodingResp.Data.Result.ID,
		StreamID:        aacStreamResp.Data.Result.ID,
		MuxingID:        audioMuxingResp.Data.Result.ID,
	}
	_, err = hlsService.AddMediaInfo(*hlsManifestResp.Data.Result.ID, audioMediaInfo)
	errorHandler(err)

	for _, codecConfigDefinition := range videoEncodingProfiles {
		videoStreamInfo := &models.StreamInfo{
			Audio:       stringToPtr("audio_group"),
			SegmentPath: stringToPtr(fmt.Sprintf("../video/%d", *codecConfigDefinition.height)),
			URI:         stringToPtr(fmt.Sprintf("video_%d.m3u8", *codecConfigDefinition.height)),
			EncodingID:  encodingResp.Data.Result.ID,
			StreamID:    codecConfigDefinition.streamId,
			MuxingID:    codecConfigDefinition.muxingId,
		}

		_, err := hlsService.AddStreamInfo(*hlsManifestResp.Data.Result.ID, videoStreamInfo)
		errorHandler(err)
	}

	_, err = hlsService.Start(*hlsManifestResp.Data.Result.ID)
	errorHandler(err)

	status = ""
	for status != "FINISHED" {
		time.Sleep(5 * time.Second)
		statusResp, err := hlsService.RetrieveStatus(*hlsManifestResp.Data.Result.ID)
		if err != nil {
			fmt.Println("error in Manifest Status")
			fmt.Println(err)
			return
		}
		// Polling and Printing out the response
		fmt.Printf("%+v\n", statusResp)
		status = *statusResp.Data.Result.Status
		if status == "ERROR" {
			fmt.Println("error in Manifest Status")
			fmt.Printf("%+v\n", statusResp)
			return
		}
	}

	// Delete Encoding
	_, err = encodingS.Delete(*encodingResp.Data.Result.ID)
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		switch err.(type) {
		case models.BitmovinError:
			fmt.Println("Bitmovin Error")
		default:
			fmt.Println("General Error")
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func stringToPtr(s string) *string {
	return &s
}

func intToPtr(i int64) *int64 {
	return &i
}

func boolToPtr(b bool) *bool {
	return &b
}

func floatToPtr(f float64) *float64 {
	return &f
}
