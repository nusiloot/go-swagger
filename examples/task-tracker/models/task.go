package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Task a structure describing a complete task.
//
// A Task is the main entity in this application. Everything revolves around tasks and managing them.
//
// swagger:model Task
type Task struct {
	TaskCard

	// The attached files.
	//
	// An issue can have at most 20 files attached to it.
	//
	Attachments map[string]TaskAttachmentsAnon `json:"attachments,omitempty"`

	// The 5 most recent items for this issue.
	//
	// The detail view of an issue includes the 5 most recent comments.
	// This field is read only, comments are added through a separate process.
	//
	// Read Only: true
	Comments []*Comment `json:"comments,omitempty"`

	// The time at which this issue was last updated.
	//
	// This field is read only so it's only sent as part of the response.
	//
	// Read Only: true
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// last updated by
	LastUpdatedBy *UserCard `json:"lastUpdatedBy,omitempty"`

	// reported by
	ReportedBy *UserCard `json:"reportedBy,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.TaskCard.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for k := range m.Attachments {

		if swag.IsZero(m.Attachments[k]) { // not required
			continue
		}

		if val, ok := m.Attachments[k]; ok {

			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateComments(formats strfmt.Registry) error {

	for i := 0; i < len(m.Comments); i++ {

		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {

			if err := m.Comments[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateLastUpdatedBy(formats strfmt.Registry) error {

	if m.LastUpdatedBy != nil {

		if err := m.LastUpdatedBy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Task) validateReportedBy(formats strfmt.Registry) error {

	if m.ReportedBy != nil {

		if err := m.ReportedBy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// TaskAttachmentsAnontask attachments anon
// swagger:model TaskAttachmentsAnon
type TaskAttachmentsAnon struct {

	// The content type of the file.
	//
	// The content type of the file is inferred from the upload request.
	//
	// Read Only: true
	ContentType string `json:"contentType,omitempty"`

	// Extra information to attach to the file.
	//
	// This is a free form text field with support for github flavored markdown.
	//
	// Min Length: 3
	Description string `json:"description,omitempty"`

	// The name of the file.
	//
	// This name is inferred from the upload request.
	//
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The file size in bytes.
	//
	// This property was generated during the upload request of the file.
	// Read Only: true
	Size float64 `json:"size,omitempty"`

	// The url to download or view the file.
	//
	// This URL is generated on the server, based on where it was able to store the file when it was uploaded.
	//
	// Read Only: true
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this task attachments anon
func (m *TaskAttachmentsAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAttachmentsAnon) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 3); err != nil {
		return err
	}

	return nil
}
