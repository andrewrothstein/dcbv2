package dcbv2

import ("fmt"; "strings"; "os"; "path")

type Setting interface {
	Name() string
	Get() (string, bool)
}

type LiteralSetting struct {
	name, val string
}

func (l *LiteralSetting) Get() (string, bool) { return l.val, true }
func (l *LiteralSetting) Name() string { return l.name }
func CreateLiteralSetting(name string, val string) LiteralSetting {
	return LiteralSetting{name: name, val: val}
}

type EnvSetting struct {
	name, envkey string
}

func (e *EnvSetting) Get() (string, bool) { return os.LookupEnv(e.envkey) }
func (e *EnvSetting) Name() string { return e.name }

func CreateEnvSetting(l []string, sep string, prefix string) EnvSetting {
	envkey := strings.Join(append([]string{prefix}, l...), sep)
	return EnvSetting{name: fmt.Sprintf("Env[%s]", envkey), envkey: envkey}
}

type SplitSetting struct {
	name, envkey, sep string
	idx int
}

func (ss *SplitSetting) Get() (string, bool) {
	ev, has := os.LookupEnv(ss.envkey)
	if has {
		s := strings.Split(ev, ss.sep)
		if ss.idx < len(s) {
			return s[ss.idx], true
		}
	}
	return "", false
}
func (ss *SplitSetting) Name() string { return ss.name }

func CreateSplitSetting(name string, envkey string, sep string, idx int) SplitSetting {
	return SplitSetting{name, envkey, sep, idx}
}

func CreateOwnerFromSlugSetting(envkey string) SplitSetting {
	name := fmt.Sprintf("OwnerFromSlugSetting[%s]", envkey)
	return CreateSplitSetting(name, envkey, "/", 0)
}

func CreateProjectFromSlugSetting(envkey string) SplitSetting {
	name := fmt.Sprintf("ProjectFromSlug[%s]", envkey)
	return CreateSplitSetting(name, envkey, "/", 1)
}

type CwdSetting struct {}

func (c *CwdSetting) Name() string { return "CwdSetting" }
func (c *CwdSetting) Get() (string, bool) {
	wd, error := os.Getwd()
	if error != nil {
		return "", false
	} else {
		return path.Base(wd), true
	}
}

type ParentCwdSetting struct {}

func (p *ParentCwdSetting) Name() { return "ParentCwdSetting" }
func (p *ParentCwdSetting) Get() (string, bool) {
	wd, error := os.Getwd()
	if error != nil {
		return "", false
	} else {
		return path.Base(path.Dir(wd)), true
	}
}

