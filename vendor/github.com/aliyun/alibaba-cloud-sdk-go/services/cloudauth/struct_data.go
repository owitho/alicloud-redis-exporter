package cloudauth

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

// Data is a nested struct in cloudauth response
type Data struct {
	Name                 string       `json:"Name" xml:"Name"`
	IdCardFrontPic       string       `json:"IdCardFrontPic" xml:"IdCardFrontPic"`
	TrustedScore         float64      `json:"TrustedScore" xml:"TrustedScore"`
	FacePic              string       `json:"FacePic" xml:"FacePic"`
	IdCardExpiry         string       `json:"IdCardExpiry" xml:"IdCardExpiry"`
	AuditConclusions     string       `json:"AuditConclusions" xml:"AuditConclusions"`
	IdCardType           string       `json:"IdCardType" xml:"IdCardType"`
	SimilarityScore      float64      `json:"SimilarityScore" xml:"SimilarityScore"`
	IdentificationNumber string       `json:"IdentificationNumber" xml:"IdentificationNumber"`
	Address              string       `json:"Address" xml:"Address"`
	Sex                  string       `json:"Sex" xml:"Sex"`
	ImgWidth             int          `json:"ImgWidth" xml:"ImgWidth"`
	IdCardBackPic        string       `json:"IdCardBackPic" xml:"IdCardBackPic"`
	CloudauthPageUrl     string       `json:"CloudauthPageUrl" xml:"CloudauthPageUrl"`
	ImgHeight            int          `json:"ImgHeight" xml:"ImgHeight"`
	StatusCode           int          `json:"StatusCode" xml:"StatusCode"`
	ConfidenceThresholds string       `json:"ConfidenceThresholds" xml:"ConfidenceThresholds"`
	StsToken             StsToken     `json:"StsToken" xml:"StsToken"`
	VerifyToken          VerifyToken  `json:"VerifyToken" xml:"VerifyToken"`
	VerifyStatus         VerifyStatus `json:"VerifyStatus" xml:"VerifyStatus"`
	FaceInfos            FaceInfos    `json:"FaceInfos" xml:"FaceInfos"`
}
