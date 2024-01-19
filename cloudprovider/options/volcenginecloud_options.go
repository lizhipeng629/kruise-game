package options

type VolcengineCloudOptions struct {
	Enable     bool       `toml:"enable"`
	CLBOptions CLBOptions `toml:"clb"`
}

type CLBOptions struct {
	MaxPort int32 `toml:"max_port"`
	MinPort int32 `toml:"min_port"`
}

func (v VolcengineCloudOptions) Valid() bool {
	return v.Enable
}

func (v VolcengineCloudOptions) Enabled() bool {
	clbOptions := v.CLBOptions
	if clbOptions.MinPort < 1 {
		return false
	}
	if clbOptions.MaxPort > 65535 {
		return false
	}
	return true
}
