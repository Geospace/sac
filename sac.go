// sac is a simple key/value store that can work with
// a configuration file on your filesystem
// supports nested values (like.this.key for example)
// no automatic conversions (only some casts)
// no error handling by default, returns the 0 equivalent
// safe functions equivalents in case you need error handling
// case sensitive
package sac

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
)

// Sac represents a configuration file in memory
type Sac struct {
	fields map[string]interface{}
	// ConfigType is the configuration format (JSON, YAML), see sac.ConfigType
	ConfigType ConfigType
	// Path is the path the configuration will be written to
	Path string
	fs   afero.Fs
}

const (
	nestingChar     = "."
	writePermission = 0644
)

// New returns a new sac, ready to be used
func New() *Sac {
	fields := make(map[string]interface{})
	return &Sac{
		fields,
		UNKNOWN,
		"",
		afero.NewOsFs(),
	}
}

// ChangeFS changes the underlying filesystem that will be used for
// config read/write operations
func (s *Sac) ChangeFS(fs afero.Fs) {
	s.fs = fs
}

// Set writes a value into the store
func (s *Sac) Set(k string, v interface{}) {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	level[last] = v
}

// Get reads a value from the store, without making any
// assertion about its type
func (s *Sac) Get(k string) interface{} {
	split := strings.Split(k, nestingChar)
	last := split[len(split)-1]
	level := getNest(s.fields, split[0:len(split)-1])

	if v, ok := level[last]; ok {
		return v
	}

	return nil
}

// Delete removes a key from the map
func (s *Sac) Delete(k string) {
	delete(s.fields, k)
}

// ReadConfig reads the configuration file at path into the sac object
func (s *Sac) ReadConfig(path string) error {
	switch s.ConfigType {
	case YAML:
		return fmt.Errorf("not implemented")
	case JSON:
		return s.readConfigJSON(path)
	default:
		return fmt.Errorf("bad config type")
	}
}

// WriteConfig writes the configuration file at the path it has been read
// and it does not care if the file exist or not
func (s *Sac) WriteConfig() error {
	switch s.ConfigType {
	case YAML:
		return fmt.Errorf("not implemented")
	case JSON:
		return s.writeConfigJSON(s.Path)
	default:
		return fmt.Errorf("bad config type")
	}
}

// WriteConfigSafe writes the configuration file, only if the file does not
// exist
func (s *Sac) WriteConfigSafe() error {
	if _, err := s.fs.Stat(s.Path); !os.IsNotExist(err) {
		return nil
	}

	return s.WriteConfig()
}
