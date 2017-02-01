package types

import "gopkg.in/mgo.v2/bson"

type Encoder struct {
  Id  bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Hostname string `json:"hostname"`
  Channels []bson.ObjectId `json:"channels"`
  Online bool `json:"online"`
}

const (
  Video string = "video"
  Audio string = "audio"
  Teletext string = "teletext"
)

type Channel struct {
  Id  bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name"`
  IsRunning bool `json:"is_running"`
  Configuration Config `json:"config"`
  Encoder  bson.ObjectId `json:"encoder"  bson:"encoder,omitempty"`
}

type Config struct {
  Version int64 `json:"version"`
  Input InputObj `json:"input"`
}

type InputObj struct {
  SegmentDuration int `json:"segment_duration_in_ms"`
  Url string `json:"url"`
  PIDs map[string]Pid `json:"PIDs"`
}

type Pid struct {
  Type string `json:"type"`
  Deinterlace string `json:"deinterlace,omitempty"`
  Processes []Process `json:"processes"`
  Lang string `json:"lang,omitempty"`
  Name string `json:"name,omitempty"`
}

type Process struct {
  Pages []Page `json:"pages,omitempty"`
  Encodings []Encoding `json:"encodings"`
  Outputs []Output `json:"outputs"`
}

type Page struct {
  Page string `json:"page"`
  Name string `json:"name"`
}

type Encoding struct {
  Width int64 `json:"width,omitempty"`
  Height int64 `json:"height,omitempty"`
  Bitrate int64 `json:"bitrate,omitempty"`
  Samplerate int64 `json:"sample_rate,omitempty"`
  NumberOfChannels int64 `json:"num_channels,omitempty"`
  Type string `json:"type", omitempty`
}

type Output struct {
  Type string `json:"type"`
  Path string `json:"path"`
}
