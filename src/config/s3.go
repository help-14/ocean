package config

type S3Config struct {
	Url             string `yaml:"url"`
	Region          string `yaml:"region"`
	BucketName      string `yaml:"bucket"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}
