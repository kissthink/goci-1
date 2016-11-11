package domain

import (
	"errors"
	"log"
	"reflect"
)

var (
	pluginRegistry = make(map[string]reflect.Type)
)

func PluginManagerLoadPlugins() {
	PluginManagerAddPlugin(PLUGIN_CLI_NAME, reflect.TypeOf(PluginCLI{}))
	PluginManagerAddPlugin(PLUGIN_JS_NAME, reflect.TypeOf(PluginJS{}))
	log.Println("Plugins : OK")
}

func PluginManagerAddPlugin(name string, plugin reflect.Type) {
	pluginRegistry[name] = plugin
}

func PluginManagerInit(job *Job, step *ProjectTaskStep, stepIndex int) (*IPlugin, error) {
	if job == nil {
		return nil, errors.New("Job is invalid")
	}

	if step == nil {
		return nil, errors.New("Step is invalid")
	}

	for pluginName, pluginType := range pluginRegistry {
		if pluginName == step.Plugin {
			plugin := reflect.New(pluginType).Interface().(IPlugin)
			plugin.Init(job, step, stepIndex)
			return &plugin, nil
		}
	}

	return nil, errors.New("Plugin not found")
}
