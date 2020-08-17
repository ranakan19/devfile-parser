package data

import (
	"github.com/devfile/kubernetes-api/pkg/apis/workspaces/v1alpha1"
	"github.com/ranakan19/parser/pkg/devfile/parser/data/common"
)

// DevfileData is an interface that defines functions for Devfile data operations
type DevfileData interface {
	SetSchemaVersion(version string)

	// parent related methods
	GetParent() v1alpha1.Parent
	SetParent(parent v1alpha1.Parent)

	// event related methods
	GetEvents() v1alpha1.WorkspaceEvents
	AddEvents(events v1alpha1.WorkspaceEvents) error
	UpdateEvents(postStart, postStop, preStart, preStop []string)

	// component related methods
	GetComponents() []v1alpha1.Component
	AddComponents(components []v1alpha1.Component) error
	UpdateComponent(component v1alpha1.Component)

	// project related methods
	GetProjects() []v1alpha1.Project
	AddProjects(projects []v1alpha1.Project) error
	UpdateProject(project v1alpha1.Project)

	// command related methods
	GetCommands() []v1alpha1.Command
	AddCommands(commands []v1alpha1.Command) error
	UpdateCommand(command v1alpha1.Command)

	GetMetadata() common.DevfileMetadata
}
