package mts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

type Video struct {
	Crop         string     `json:"Crop" xml:"Crop"`
	Gop          string     `json:"Gop" xml:"Gop"`
	Bufsize      string     `json:"Bufsize" xml:"Bufsize"`
	Remove       string     `json:"Remove" xml:"Remove"`
	Height       string     `json:"Height" xml:"Height"`
	Degrain      string     `json:"Degrain" xml:"Degrain"`
	Crf          string     `json:"Crf" xml:"Crf"`
	Width        string     `json:"Width" xml:"Width"`
	MaxFps       string     `json:"MaxFps" xml:"MaxFps"`
	Pad          string     `json:"Pad" xml:"Pad"`
	Bitrate      string     `json:"Bitrate" xml:"Bitrate"`
	ResoPriority string     `json:"ResoPriority" xml:"ResoPriority"`
	Fps          string     `json:"Fps" xml:"Fps"`
	Preset       string     `json:"Preset" xml:"Preset"`
	ScanMode     string     `json:"ScanMode" xml:"ScanMode"`
	Qscale       string     `json:"Qscale" xml:"Qscale"`
	PixFmt       string     `json:"PixFmt" xml:"PixFmt"`
	Profile      string     `json:"Profile" xml:"Profile"`
	Codec        string     `json:"Codec" xml:"Codec"`
	Maxrate      string     `json:"Maxrate" xml:"Maxrate"`
	BitrateBnd   BitrateBnd `json:"BitrateBnd" xml:"BitrateBnd"`
}
