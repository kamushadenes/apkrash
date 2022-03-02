package apk

func (m *Manifest) GetPermissions() []string {
	var permissions []string
	for _, p := range m.UsesPermission {
		permissions = append(permissions, p.Name)
	}
	return permissions
}

func (m *Manifest) GetFeatures() []string {
	var features []string
	for _, f := range m.UsesFeature {
		features = append(features, f.Name)
	}
	return features
}

func (m *Manifest) GetActivities() []string {
	var activities []string
	for _, a := range m.Application.Activity {
		activities = append(activities, a.Name)
	}
	return activities
}

func (m *Manifest) GetServices() []string {
	var services []string
	for _, s := range m.Application.Service {
		services = append(services, s.Name)
	}
	return services
}

func (m *Manifest) GetReceivers() []string {
	var receivers []string
	for _, r := range m.Application.Receiver {
		receivers = append(receivers, r.Name)
	}
	return receivers
}

func (m *Manifest) GetProviders() []string {
	var providers []string
	for _, p := range m.Application.Provider {
		providers = append(providers, p.Name)
	}
	return providers
}

func (m *Manifest) GetPackageName() string {
	return m.Package
}

func (m *Manifest) GetApplicationName() string {
	return m.Application.Name
}
