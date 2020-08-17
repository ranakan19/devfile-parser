package version200

import (
	"strings"

	v1 "github.com/devfile/kubernetes-api/pkg/apis/workspaces/v1alpha1"
	"github.com/ranakan19/parser/pkg/devfile/parser/data/common"
)

// GetComponents returns the slice of DevfileComponent objects parsed from the Devfile
func (d *Devfile200) GetComponents() []v1.Component {
	return d.Components
}

// AddComponents adds the slice of DevfileComponent objects to the devfile's components
// if a component is already defined, error out
func (d *Devfile200) AddComponents(components []v1.Component) error {

	// different map for volume and container component as a volume and a container with same name
	// can exist in devfile
	containerMap := make(map[string]bool)
	volumeMap := make(map[string]bool)

	for _, component := range d.Components {
		if component.Volume != nil {
			volumeMap[component.Volume.Name] = true
		}
		if component.Container != nil {
			containerMap[component.Container.Name] = true
		}
	}

	for _, component := range components {

		if component.Volume != nil {
			if _, ok := volumeMap[component.Volume.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.AlreadyExistError{Name: component.Volume.Name, Field: "component"}
			}
		}

		if component.Container != nil {
			if _, ok := containerMap[component.Container.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.AlreadyExistError{Name: component.Container.Name, Field: "component"}
			}
		}
	}
	return nil
}

// UpdateComponent updates the component with the given name
func (d *Devfile200) UpdateComponent(component v1.Component) {
	for i := range d.Components {
		if d.Components[i].Container.Name == strings.ToLower(component.Container.Name) {
			d.Components[i] = component
		}
	}
}

// GetCommands returns the slice of DevfileCommand objects parsed from the Devfile
func (d *Devfile200) GetCommands() []v1.Command {
	var commands []v1.Command

	for _, command := range d.Commands {
		// we convert devfile command id to lowercase so that we can handle
		// cases efficiently without being error prone
		// we also convert the odo push commands from build-command and run-command flags
		command.Exec.Id = strings.ToLower(command.Exec.Id)
		commands = append(commands, command)
	}

	return commands
}

// AddCommands adds the slice of DevfileCommand objects to the Devfile's commands
// if a command is already defined, error out
func (d *Devfile200) AddCommands(commands []v1.Command) error {
	commandsMap := make(map[string]bool)
	for _, command := range d.Commands {
		commandsMap[command.Exec.Id] = true
	}

	for _, command := range commands {
		if _, ok := commandsMap[command.Exec.Id]; !ok {
			d.Commands = append(d.Commands, command)
		} else {
			return &common.AlreadyExistError{Name: command.Exec.Id, Field: "command"}
		}
	}
	return nil
}

// UpdateCommand updates the command with the given id
func (d *Devfile200) UpdateCommand(command v1.Command) {
	for i := range d.Commands {
		if d.Commands[i].Exec.Id == strings.ToLower(command.Exec.Id) {
			d.Commands[i] = command
		}
	}
}

// GetParent returns the  DevfileParent object parsed from devfile
func (d *Devfile200) GetParent() v1.Parent {
	return d.Parent
}

// SetParent sets the parent for the devfile
func (d *Devfile200) SetParent(parent v1.Parent) {
	d.Parent = parent
}

// GetProjects returns the DevfileProject Object parsed from devfile
func (d *Devfile200) GetProjects() []v1.Project {
	return d.Projects
}

// AddProjects adss the slice of Devfile projects to the Devfile's project list
// if a project is already defined, error out
func (d *Devfile200) AddProjects(projects []v1.Project) error {
	projectsMap := make(map[string]bool)
	for _, project := range d.Projects {
		projectsMap[project.Name] = true
	}

	for _, project := range projects {
		if _, ok := projectsMap[project.Name]; !ok {
			d.Projects = append(d.Projects, project)
		} else {
			return &common.AlreadyExistError{Name: project.Name, Field: "project"}
		}
	}
	return nil
}

// UpdateProject updates the slice of DevfileCommand projects parsed from the Devfile
func (d *Devfile200) UpdateProject(project v1.Project) {
	for i := range d.Projects {
		if d.Projects[i].Name == strings.ToLower(project.Name) {
			d.Projects[i] = project
		}
	}
}

// GetMetadata returns the DevfileMetadata Object parsed from devfile
func (d *Devfile200) GetMetadata() common.DevfileMetadata {
	return d.Metadata
}

// SetMetadata sets the metadata for devfile
func (d *Devfile200) SetMetadata(name, version string) {
	d.Metadata = common.DevfileMetadata{
		Name:    name,
		Version: version,
	}
}

// GetEvents returns the Events Object parsed from devfile
func (d *Devfile200) GetEvents() v1.WorkspaceEvents {
	return d.Events
}

// AddEvents adds the Events Object to the devfile's events
// if the event is already defined in the devfile, error out
func (d *Devfile200) AddEvents(events v1.WorkspaceEvents) error {
	if len(events.PreStop) > 0 {
		if len(d.Events.PreStop) > 0 {
			return &common.AlreadyExistError{Field: "pre stop"}
		}
		d.Events.PreStop = events.PreStop
	}

	if len(events.PreStart) > 0 {
		if len(d.Events.PreStart) > 0 {
			return &common.AlreadyExistError{Field: "pre start"}
		}
		d.Events.PreStart = events.PreStart
	}

	if len(events.PostStop) > 0 {
		if len(d.Events.PostStop) > 0 {
			return &common.AlreadyExistError{Field: "post stop"}
		}
		d.Events.PostStop = events.PostStop
	}

	if len(events.PostStart) > 0 {
		if len(d.Events.PostStart) > 0 {
			return &common.AlreadyExistError{Field: "post start"}
		}
		d.Events.PostStart = events.PostStart
	}

	return nil
}

// UpdateEvents updates the devfile's events
// it only updates the events passed to it
func (d *Devfile200) UpdateEvents(postStart, postStop, preStart, preStop []string) {
	if len(postStart) != 0 {
		d.Events.PostStart = postStart
	}
	if len(postStop) != 0 {
		d.Events.PostStop = postStop
	}
	if len(preStart) != 0 {
		d.Events.PreStart = preStart
	}
	if len(preStop) != 0 {
		d.Events.PreStop = preStop
	}
}

//SetSchemaVersion sets devfile schema version
func (d *Devfile200) SetSchemaVersion(version string) {
	d.SchemaVersion = version
}
