package config

type BackupJob struct {
	Name     string `yaml:"name"`
	UseZip   bool   `yaml:"useZip"`
	Path     string `yaml:"path"`
	UploadTo string `yaml:"uploadTo"`
	RunAt    string `yaml:"runAt"`
}
