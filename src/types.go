package main

type HelmRelease struct {
	Name    string
	Version string
	Status  string
	Content string
}

type releaseOptions struct {
	CleanupOnFail bool
	Wait          bool
	DryRun        bool
	DisableHooks  bool
	Force         bool
	Recreate      bool
	MaxHistory    string
	Timeout       string
}
