/*
ParallelCluster

ParallelCluster API

API version: 3.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the LogStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStream{}

// LogStream struct for LogStream
type LogStream struct {
	// Name of the log stream.
	LogStreamName string `json:"logStreamName"`
	// The creation time of the stream.
	CreationTime time.Time `json:"creationTime"`
	// The time of the first event of the stream.
	FirstEventTimestamp time.Time `json:"firstEventTimestamp"`
	// The time of the last event of the stream. The lastEventTime value updates on an eventual consistency basis. It typically updates in less than an hour from ingestion, but in rare situations might take longer.
	LastEventTimestamp time.Time `json:"lastEventTimestamp"`
	// The last ingestion time.
	LastIngestionTime time.Time `json:"lastIngestionTime"`
	// The sequence token.
	UploadSequenceToken string `json:"uploadSequenceToken"`
	// The Amazon Resource Name (ARN) of the log stream.
	LogStreamArn string `json:"logStreamArn"`
}

type _LogStream LogStream

// NewLogStream instantiates a new LogStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStream(logStreamName string, creationTime time.Time, firstEventTimestamp time.Time, lastEventTimestamp time.Time, lastIngestionTime time.Time, uploadSequenceToken string, logStreamArn string) *LogStream {
	this := LogStream{}
	this.LogStreamName = logStreamName
	this.CreationTime = creationTime
	this.FirstEventTimestamp = firstEventTimestamp
	this.LastEventTimestamp = lastEventTimestamp
	this.LastIngestionTime = lastIngestionTime
	this.UploadSequenceToken = uploadSequenceToken
	this.LogStreamArn = logStreamArn
	return &this
}

// NewLogStreamWithDefaults instantiates a new LogStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamWithDefaults() *LogStream {
	this := LogStream{}
	return &this
}

// GetLogStreamName returns the LogStreamName field value
func (o *LogStream) GetLogStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogStreamName
}

// GetLogStreamNameOk returns a tuple with the LogStreamName field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLogStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogStreamName, true
}

// SetLogStreamName sets field value
func (o *LogStream) SetLogStreamName(v string) {
	o.LogStreamName = v
}

// GetCreationTime returns the CreationTime field value
func (o *LogStream) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *LogStream) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetFirstEventTimestamp returns the FirstEventTimestamp field value
func (o *LogStream) GetFirstEventTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FirstEventTimestamp
}

// GetFirstEventTimestampOk returns a tuple with the FirstEventTimestamp field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetFirstEventTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstEventTimestamp, true
}

// SetFirstEventTimestamp sets field value
func (o *LogStream) SetFirstEventTimestamp(v time.Time) {
	o.FirstEventTimestamp = v
}

// GetLastEventTimestamp returns the LastEventTimestamp field value
func (o *LogStream) GetLastEventTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastEventTimestamp
}

// GetLastEventTimestampOk returns a tuple with the LastEventTimestamp field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLastEventTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEventTimestamp, true
}

// SetLastEventTimestamp sets field value
func (o *LogStream) SetLastEventTimestamp(v time.Time) {
	o.LastEventTimestamp = v
}

// GetLastIngestionTime returns the LastIngestionTime field value
func (o *LogStream) GetLastIngestionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastIngestionTime
}

// GetLastIngestionTimeOk returns a tuple with the LastIngestionTime field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLastIngestionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastIngestionTime, true
}

// SetLastIngestionTime sets field value
func (o *LogStream) SetLastIngestionTime(v time.Time) {
	o.LastIngestionTime = v
}

// GetUploadSequenceToken returns the UploadSequenceToken field value
func (o *LogStream) GetUploadSequenceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadSequenceToken
}

// GetUploadSequenceTokenOk returns a tuple with the UploadSequenceToken field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetUploadSequenceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadSequenceToken, true
}

// SetUploadSequenceToken sets field value
func (o *LogStream) SetUploadSequenceToken(v string) {
	o.UploadSequenceToken = v
}

// GetLogStreamArn returns the LogStreamArn field value
func (o *LogStream) GetLogStreamArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogStreamArn
}

// GetLogStreamArnOk returns a tuple with the LogStreamArn field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLogStreamArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogStreamArn, true
}

// SetLogStreamArn sets field value
func (o *LogStream) SetLogStreamArn(v string) {
	o.LogStreamArn = v
}

func (o LogStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logStreamName"] = o.LogStreamName
	toSerialize["creationTime"] = o.CreationTime
	toSerialize["firstEventTimestamp"] = o.FirstEventTimestamp
	toSerialize["lastEventTimestamp"] = o.LastEventTimestamp
	toSerialize["lastIngestionTime"] = o.LastIngestionTime
	toSerialize["uploadSequenceToken"] = o.UploadSequenceToken
	toSerialize["logStreamArn"] = o.LogStreamArn
	return toSerialize, nil
}

func (o *LogStream) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logStreamName",
		"creationTime",
		"firstEventTimestamp",
		"lastEventTimestamp",
		"lastIngestionTime",
		"uploadSequenceToken",
		"logStreamArn",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLogStream := _LogStream{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogStream)

	if err != nil {
		return err
	}

	*o = LogStream(varLogStream)

	return err
}

type NullableLogStream struct {
	value *LogStream
	isSet bool
}

func (v NullableLogStream) Get() *LogStream {
	return v.value
}

func (v *NullableLogStream) Set(val *LogStream) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStream) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStream(val *LogStream) *NullableLogStream {
	return &NullableLogStream{value: val, isSet: true}
}

func (v NullableLogStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
