package config

type BackupJob struct {
	Name     string      `yaml:"name"`
	UseZip   bool        `yaml:"useZip"`
	Path     string      `yaml:"path"`
	UploadTo ServiceName `yaml:"uploadTo"`
}
