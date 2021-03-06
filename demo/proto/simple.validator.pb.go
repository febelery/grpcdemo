// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: simple.proto

package rpcdemo

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_SimpleRequest_Data = regexp.MustCompile(`^[a-z]{2,20}$`)
var _regex_SimpleRequest_UserId = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *SimpleRequest) Validate() error {
	if !_regex_SimpleRequest_Data.MatchString(this.Data) {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,20}$"`, this.Data))
	}
	if !_regex_SimpleRequest_UserId.MatchString(this.UserId) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.UserId))
	}
	if _, ok := Action_name[int32(this.Do)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Do", fmt.Errorf(`value '%v' must be a valid Action field`, this.Do))
	}
	return nil
}
func (this *SimpleResponse) Validate() error {
	return nil
}
