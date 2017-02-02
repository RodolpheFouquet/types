package types

import "gopkg.in/mgo.v2/bson"

type Encoder struct {
  Id  bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Hostname string `json:"hostname"`
  Channels []bson.ObjectId `json:"channels"`
  Online bool `json:"online"`
}

// PID type
const (
  Video string = "video"
  Audio string = "audio"
  Teletext string = "teletext"
)

// Encoding type
const (
  Nvenc string = "nvenc"
  Software string ="software"
  Quicksync string ="quicksync"
)

type Channel struct {
  Id  bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name string `json:"name"`
  IsRunning bool `json:"is_running"`
  Configuration Config `json:"config"` //fk-json.exe config
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
  Name string `json:"name,omitempty"`
}

type Page struct {
  Page string `json:"page"`
  Name string `json:"name"`
  Lang string `json:"lang"`
}

type Encoding struct {
  Width int64 `json:"width,omitempty"`
  Height int64 `json:"height,omitempty"`
  Bitrate int64 `json:"bitrate,omitempty"`
  Samplerate int64 `json:"sample_rate,omitempty"`
  NumberOfChannels int64 `json:"num_channels,omitempty"`
  Type string `json:"type,omitempty"`
  Custom string `json:"custom,omitempty"`
}

type Output struct {
  Type string `json:"type"`
  Path string `json:"path"`
}

type Error struct {
  Err string `json:"err"`
}

func NewError(message string) Error {
  return Error{message}
}
