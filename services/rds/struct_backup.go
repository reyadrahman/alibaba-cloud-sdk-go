package rds

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

type Backup struct {
	BackupExtractionStatus    string `json:"BackupExtractionStatus" xml:"BackupExtractionStatus"`
	BackupMode                string `json:"BackupMode" xml:"BackupMode"`
	BackupStatus              string `json:"BackupStatus" xml:"BackupStatus"`
	DBInstanceId              string `json:"DBInstanceId" xml:"DBInstanceId"`
	StoreStatus               string `json:"StoreStatus" xml:"StoreStatus"`
	BackupDBNames             string `json:"BackupDBNames" xml:"BackupDBNames"`
	BackupStartTime           string `json:"BackupStartTime" xml:"BackupStartTime"`
	BackupMethod              string `json:"BackupMethod" xml:"BackupMethod"`
	BackupSize                int    `json:"BackupSize" xml:"BackupSize"`
	TotalBackupSize           int    `json:"TotalBackupSize" xml:"TotalBackupSize"`
	BackupIntranetDownloadURL string `json:"BackupIntranetDownloadURL" xml:"BackupIntranetDownloadURL"`
	BackupDownloadURL         string `json:"BackupDownloadURL" xml:"BackupDownloadURL"`
	HostInstanceID            string `json:"HostInstanceID" xml:"HostInstanceID"`
	BackupScale               string `json:"BackupScale" xml:"BackupScale"`
	BackupId                  string `json:"BackupId" xml:"BackupId"`
	BackupEndTime             string `json:"BackupEndTime" xml:"BackupEndTime"`
	BackupLocation            string `json:"BackupLocation" xml:"BackupLocation"`
	BackupType                string `json:"BackupType" xml:"BackupType"`
}
