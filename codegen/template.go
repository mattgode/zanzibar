// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package codegen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	tmpl "text/template"

	"github.com/pkg/errors"
)

var funcMap = tmpl.FuncMap{
	"title": strings.Title,
	"Title": strings.Title,
}

// Template generates code for edge gateway clients and edgegateway endpoints.
type Template struct {
	template *tmpl.Template
}

// NewTemplate creates a bundle of templates.
func NewTemplate(templatePattern string) (*Template, error) {
	t, err := tmpl.New("main").Funcs(funcMap).ParseGlob(templatePattern)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template files")
	}
	return &Template{
		template: t,
	}, nil
}

// GenerateClientFile generates Go http code for services defined in thrift file.
// It returns the path of generated file or an error.
func (t *Template) GenerateClientFile(thrift string, h *PackageHelper) (string, error) {
	m, err := NewModuleSpec(thrift, h)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse thrift file:")
	}
	if len(m.Services) == 0 {
		return "", errors.Errorf("no service is found in thrift file %s", thrift)
	}

	err = t.execTemplateAndFmt("http_client.tmpl", m.GoClientFilePath, m)
	if err != nil {
		return "", err
	}

	err = t.execTemplateAndFmt(
		"http_client_structs.tmpl", m.GoClientStructsFilePath, m,
	)
	if err != nil {
		return "", err
	}

	return m.GoClientFilePath, nil
}

// GenerateHandlerFile generates Go http code for endpoint.
func (t *Template) GenerateHandlerFile(
	thrift string, h *PackageHelper, methodName string,
) (string, error) {
	m, err := NewModuleSpec(thrift, h)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse thrift file.")
	}
	if len(m.Services) == 0 {
		return "", errors.Errorf("no service is found in thrift file %s", thrift)
	}

	if len(m.Services) != 1 {
		panic("TODO: Do not support multiple services in thrift file yet.")
	}

	service := m.Services[0]
	var method *MethodSpec
	for _, v := range service.Methods {
		if v.Name == methodName {
			method = v
			break
		}
	}

	if method == nil {
		return "", errors.Errorf(
			"could not find method name %s in thrift file %s",
			methodName, thrift,
		)
	}

	endpointName := strings.Split(method.EndpointName, ".")[0]
	handlerName := strings.Split(method.Name, ".")[0]
	dest, err := h.TargetEndpointPath(thrift, methodName)
	if err != nil {
		return "", errors.Wrap(err, "Could not generate endpoint path")
	}

	// TODO(sindelar): Use an endpoint to client map instead of proxy naming.
	downstreamService := endpointName
	downstreamMethod := handlerName

	vals := map[string]string{
		"MyHandler":         handlerName,
		"Package":           endpointName,
		"DownstreamService": downstreamService,
		"DownstreamMethod":  downstreamMethod,
	}

	err = t.execTemplateAndFmt("endpoint_template.tmpl", dest, vals)
	if err != nil {
		return "", err
	}

	return dest, nil
}

// GenerateHandlerTestFile generates Go http code for endpoint test.
func (t *Template) GenerateHandlerTestFile(
	thrift string, h *PackageHelper, methodName string,
) (string, error) {
	m, err := NewModuleSpec(thrift, h)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse thrift file.")
	}
	if len(m.Services) == 0 {
		return "", errors.Errorf("no service is found in thrift file %s", thrift)
	}

	if len(m.Services) != 1 {
		panic("TODO: Do not support multiple services in thrift file yet.")
	}

	service := m.Services[0]
	var method *MethodSpec
	for _, v := range service.Methods {
		if v.Name == methodName {
			method = v
			break
		}
	}

	if method == nil {
		return "", errors.Errorf(
			"could not find method name %s in thrift file %s",
			methodName, thrift,
		)
	}

	endpointName := strings.Split(method.EndpointName, ".")[0]
	handlerName := strings.Split(method.Name, ".")[0]
	dest, err := h.TargetEndpointTestPath(thrift, methodName)
	if err != nil {
		return "", errors.Wrap(err, "Could not generate endpoint path")
	}

	// TODO(sindelar): Use an endpoint to client map instead of proxy naming.
	downstreamService := endpointName
	downstreamMethod := handlerName

	// TODO(sindelar): Dummy data, read from golden file.
	var clientResponse = "{\\\"statusCode\\\":200}"
	var endpointPath = "/googlenow/add-credentials"
	var endpointHTTPMethod = "POST"
	var clientPath = "/add-credentials"
	var clientHTTPMethod = "POST"
	var endpointRequest = "{\\\"testrequest\\\"}"
	var clientName = "GoogleNow"

	vals := map[string]string{
		"MyHandler":          handlerName,
		"Package":            endpointName,
		"DownstreamService":  downstreamService,
		"DownstreamMethod":   downstreamMethod,
		"EndpointPath":       endpointPath,
		"EndpointHttpMethod": endpointHTTPMethod,
		"ClientPath":         clientPath,
		"ClientName":         clientName,
		"ClientHttpMethod":   clientHTTPMethod,
		"ClientResponse":     clientResponse,
		"EndpointRequest":    endpointRequest,
	}

	err = t.execTemplateAndFmt("endpoint_test_template.tmpl", dest, vals)
	if err != nil {
		return "", err
	}

	return dest, nil
}

// GenerateEndpointFile generates Go code for an zanzibar endpoint defined in
// thrift file. It returns the path of generated file or an error.
func (t *Template) GenerateEndpointFile(thrift string, h *PackageHelper) (string, error) {
	m, err := NewModuleSpec(thrift, h)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse thrift file:")
	}
	if len(m.Services) == 0 {
		return "", errors.Errorf("no service is found in thrift file %s", thrift)
	}

	// TODO: method name ??
	dest, err := h.TargetEndpointPath(thrift, "")
	if err != nil {
		return "", errors.Wrap(err, "Could not generate endpoint path")
	}

	err = t.execTemplateAndFmt("endpoint.tmpl", dest, m)
	if err != nil {
		return "", err
	}
	return dest, nil
}

func (t *Template) execTemplateAndFmt(templName string, filePath string, data interface{}) error {
	file, err := openFileOrCreate(filePath)
	if err != nil {
		return errors.Wrapf(err, "failed to open file: ", err)
	}
	if err := t.template.ExecuteTemplate(file, templName, data); err != nil {
		return errors.Wrapf(err, "failed to execute template files for file %s", file)
	}

	gofmtCmd := exec.Command("gofmt", "-s", "-w", "-e", filePath)
	gofmtCmd.Stdout = os.Stdout
	gofmtCmd.Stderr = os.Stderr

	if err := gofmtCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to gofmt file: %s", filePath)
	}

	goimportsCmd := exec.Command("goimports", "-w", "-e", filePath)
	goimportsCmd.Stdout = os.Stdout
	goimportsCmd.Stderr = os.Stderr

	if err := goimportsCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to goimports file: %s", filePath)
	}

	if err := file.Close(); err != nil {
		return errors.Wrap(err, "failed to close file")
	}

	return nil
}

func openFileOrCreate(file string) (*os.File, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			return nil, err
		}
	}
	return os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
}