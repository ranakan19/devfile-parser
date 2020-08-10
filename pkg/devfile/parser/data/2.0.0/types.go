package version200

import (
	"github.com/devfile/parser/pkg/devfile/parser/data/common"
	apiComp "github.com/ranakan19/kubernetes-api/pkg/apis/workspaces/v1alpha1"
)

// Devfile200 Devfile schema.
type Devfile200 struct {

	// Predefined, ready-to-use, workspace-related commands
	Commands []common.DevfileCommand `json:"commands,omitempty"`

	// List of the workspace components, such as editor and plugins, user-provided containers, or other types of components
	Components []common.DevfileComponent `json:"components,omitempty"`

	// Bindings of commands to events. Each command is referred-to by its name.
	Events apiComp.WorkspaceEvents `json:"events,omitempty"`

	// Optional metadata
	Metadata common.DevfileMetadata `json:"metadata,omitempty"`

	// Parent workspace template
	Parent apiComp.Parent `json:"parent,omitempty"`

	// Projects worked on in the workspace, containing names and sources locations
	Projects []apiComp.Project `json:"projects,omitempty"`

	// Devfile schema version
	SchemaVersion string `json:"schemaVersion"`
}
