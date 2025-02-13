// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiserver "k8s.io/apiserver/pkg/apis/apiserver"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Connection)(nil), (*apiserver.Connection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Connection_To_apiserver_Connection(a.(*Connection), b.(*apiserver.Connection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.Connection)(nil), (*Connection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_Connection_To_v1_Connection(a.(*apiserver.Connection), b.(*Connection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressSelection)(nil), (*apiserver.EgressSelection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressSelection_To_apiserver_EgressSelection(a.(*EgressSelection), b.(*apiserver.EgressSelection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.EgressSelection)(nil), (*EgressSelection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_EgressSelection_To_v1_EgressSelection(a.(*apiserver.EgressSelection), b.(*EgressSelection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressSelectorConfiguration)(nil), (*apiserver.EgressSelectorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(a.(*EgressSelectorConfiguration), b.(*apiserver.EgressSelectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.EgressSelectorConfiguration)(nil), (*EgressSelectorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_EgressSelectorConfiguration_To_v1_EgressSelectorConfiguration(a.(*apiserver.EgressSelectorConfiguration), b.(*EgressSelectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TCPTransport)(nil), (*apiserver.TCPTransport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TCPTransport_To_apiserver_TCPTransport(a.(*TCPTransport), b.(*apiserver.TCPTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.TCPTransport)(nil), (*TCPTransport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_TCPTransport_To_v1_TCPTransport(a.(*apiserver.TCPTransport), b.(*TCPTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TLSConfig)(nil), (*apiserver.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TLSConfig_To_apiserver_TLSConfig(a.(*TLSConfig), b.(*apiserver.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.TLSConfig)(nil), (*TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_TLSConfig_To_v1_TLSConfig(a.(*apiserver.TLSConfig), b.(*TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Transport)(nil), (*apiserver.Transport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Transport_To_apiserver_Transport(a.(*Transport), b.(*apiserver.Transport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.Transport)(nil), (*Transport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_Transport_To_v1_Transport(a.(*apiserver.Transport), b.(*Transport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UDSTransport)(nil), (*apiserver.UDSTransport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_UDSTransport_To_apiserver_UDSTransport(a.(*UDSTransport), b.(*apiserver.UDSTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.UDSTransport)(nil), (*UDSTransport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_UDSTransport_To_v1_UDSTransport(a.(*apiserver.UDSTransport), b.(*UDSTransport), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Connection_To_apiserver_Connection(in *Connection, out *apiserver.Connection, s conversion.Scope) error {
	out.ProxyProtocol = apiserver.ProtocolType(in.ProxyProtocol)
	out.Transport = (*apiserver.Transport)(unsafe.Pointer(in.Transport))
	return nil
}

// Convert_v1_Connection_To_apiserver_Connection is an autogenerated conversion function.
func Convert_v1_Connection_To_apiserver_Connection(in *Connection, out *apiserver.Connection, s conversion.Scope) error {
	return autoConvert_v1_Connection_To_apiserver_Connection(in, out, s)
}

func autoConvert_apiserver_Connection_To_v1_Connection(in *apiserver.Connection, out *Connection, s conversion.Scope) error {
	out.ProxyProtocol = ProtocolType(in.ProxyProtocol)
	out.Transport = (*Transport)(unsafe.Pointer(in.Transport))
	return nil
}

// Convert_apiserver_Connection_To_v1_Connection is an autogenerated conversion function.
func Convert_apiserver_Connection_To_v1_Connection(in *apiserver.Connection, out *Connection, s conversion.Scope) error {
	return autoConvert_apiserver_Connection_To_v1_Connection(in, out, s)
}

func autoConvert_v1_EgressSelection_To_apiserver_EgressSelection(in *EgressSelection, out *apiserver.EgressSelection, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1_Connection_To_apiserver_Connection(&in.Connection, &out.Connection, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_EgressSelection_To_apiserver_EgressSelection is an autogenerated conversion function.
func Convert_v1_EgressSelection_To_apiserver_EgressSelection(in *EgressSelection, out *apiserver.EgressSelection, s conversion.Scope) error {
	return autoConvert_v1_EgressSelection_To_apiserver_EgressSelection(in, out, s)
}

func autoConvert_apiserver_EgressSelection_To_v1_EgressSelection(in *apiserver.EgressSelection, out *EgressSelection, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_apiserver_Connection_To_v1_Connection(&in.Connection, &out.Connection, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiserver_EgressSelection_To_v1_EgressSelection is an autogenerated conversion function.
func Convert_apiserver_EgressSelection_To_v1_EgressSelection(in *apiserver.EgressSelection, out *EgressSelection, s conversion.Scope) error {
	return autoConvert_apiserver_EgressSelection_To_v1_EgressSelection(in, out, s)
}

func autoConvert_v1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in *EgressSelectorConfiguration, out *apiserver.EgressSelectorConfiguration, s conversion.Scope) error {
	out.EgressSelections = *(*[]apiserver.EgressSelection)(unsafe.Pointer(&in.EgressSelections))
	return nil
}

// Convert_v1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration is an autogenerated conversion function.
func Convert_v1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in *EgressSelectorConfiguration, out *apiserver.EgressSelectorConfiguration, s conversion.Scope) error {
	return autoConvert_v1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in, out, s)
}

func autoConvert_apiserver_EgressSelectorConfiguration_To_v1_EgressSelectorConfiguration(in *apiserver.EgressSelectorConfiguration, out *EgressSelectorConfiguration, s conversion.Scope) error {
	out.EgressSelections = *(*[]EgressSelection)(unsafe.Pointer(&in.EgressSelections))
	return nil
}

// Convert_apiserver_EgressSelectorConfiguration_To_v1_EgressSelectorConfiguration is an autogenerated conversion function.
func Convert_apiserver_EgressSelectorConfiguration_To_v1_EgressSelectorConfiguration(in *apiserver.EgressSelectorConfiguration, out *EgressSelectorConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_EgressSelectorConfiguration_To_v1_EgressSelectorConfiguration(in, out, s)
}

func autoConvert_v1_TCPTransport_To_apiserver_TCPTransport(in *TCPTransport, out *apiserver.TCPTransport, s conversion.Scope) error {
	out.URL = in.URL
	out.TLSConfig = (*apiserver.TLSConfig)(unsafe.Pointer(in.TLSConfig))
	return nil
}

// Convert_v1_TCPTransport_To_apiserver_TCPTransport is an autogenerated conversion function.
func Convert_v1_TCPTransport_To_apiserver_TCPTransport(in *TCPTransport, out *apiserver.TCPTransport, s conversion.Scope) error {
	return autoConvert_v1_TCPTransport_To_apiserver_TCPTransport(in, out, s)
}

func autoConvert_apiserver_TCPTransport_To_v1_TCPTransport(in *apiserver.TCPTransport, out *TCPTransport, s conversion.Scope) error {
	out.URL = in.URL
	out.TLSConfig = (*TLSConfig)(unsafe.Pointer(in.TLSConfig))
	return nil
}

// Convert_apiserver_TCPTransport_To_v1_TCPTransport is an autogenerated conversion function.
func Convert_apiserver_TCPTransport_To_v1_TCPTransport(in *apiserver.TCPTransport, out *TCPTransport, s conversion.Scope) error {
	return autoConvert_apiserver_TCPTransport_To_v1_TCPTransport(in, out, s)
}

func autoConvert_v1_TLSConfig_To_apiserver_TLSConfig(in *TLSConfig, out *apiserver.TLSConfig, s conversion.Scope) error {
	out.CABundle = in.CABundle
	out.ClientKey = in.ClientKey
	out.ClientCert = in.ClientCert
	return nil
}

// Convert_v1_TLSConfig_To_apiserver_TLSConfig is an autogenerated conversion function.
func Convert_v1_TLSConfig_To_apiserver_TLSConfig(in *TLSConfig, out *apiserver.TLSConfig, s conversion.Scope) error {
	return autoConvert_v1_TLSConfig_To_apiserver_TLSConfig(in, out, s)
}

func autoConvert_apiserver_TLSConfig_To_v1_TLSConfig(in *apiserver.TLSConfig, out *TLSConfig, s conversion.Scope) error {
	out.CABundle = in.CABundle
	out.ClientKey = in.ClientKey
	out.ClientCert = in.ClientCert
	return nil
}

// Convert_apiserver_TLSConfig_To_v1_TLSConfig is an autogenerated conversion function.
func Convert_apiserver_TLSConfig_To_v1_TLSConfig(in *apiserver.TLSConfig, out *TLSConfig, s conversion.Scope) error {
	return autoConvert_apiserver_TLSConfig_To_v1_TLSConfig(in, out, s)
}

func autoConvert_v1_Transport_To_apiserver_Transport(in *Transport, out *apiserver.Transport, s conversion.Scope) error {
	out.TCP = (*apiserver.TCPTransport)(unsafe.Pointer(in.TCP))
	out.UDS = (*apiserver.UDSTransport)(unsafe.Pointer(in.UDS))
	return nil
}

// Convert_v1_Transport_To_apiserver_Transport is an autogenerated conversion function.
func Convert_v1_Transport_To_apiserver_Transport(in *Transport, out *apiserver.Transport, s conversion.Scope) error {
	return autoConvert_v1_Transport_To_apiserver_Transport(in, out, s)
}

func autoConvert_apiserver_Transport_To_v1_Transport(in *apiserver.Transport, out *Transport, s conversion.Scope) error {
	out.TCP = (*TCPTransport)(unsafe.Pointer(in.TCP))
	out.UDS = (*UDSTransport)(unsafe.Pointer(in.UDS))
	return nil
}

// Convert_apiserver_Transport_To_v1_Transport is an autogenerated conversion function.
func Convert_apiserver_Transport_To_v1_Transport(in *apiserver.Transport, out *Transport, s conversion.Scope) error {
	return autoConvert_apiserver_Transport_To_v1_Transport(in, out, s)
}

func autoConvert_v1_UDSTransport_To_apiserver_UDSTransport(in *UDSTransport, out *apiserver.UDSTransport, s conversion.Scope) error {
	out.UDSName = in.UDSName
	return nil
}

// Convert_v1_UDSTransport_To_apiserver_UDSTransport is an autogenerated conversion function.
func Convert_v1_UDSTransport_To_apiserver_UDSTransport(in *UDSTransport, out *apiserver.UDSTransport, s conversion.Scope) error {
	return autoConvert_v1_UDSTransport_To_apiserver_UDSTransport(in, out, s)
}

func autoConvert_apiserver_UDSTransport_To_v1_UDSTransport(in *apiserver.UDSTransport, out *UDSTransport, s conversion.Scope) error {
	out.UDSName = in.UDSName
	return nil
}

// Convert_apiserver_UDSTransport_To_v1_UDSTransport is an autogenerated conversion function.
func Convert_apiserver_UDSTransport_To_v1_UDSTransport(in *apiserver.UDSTransport, out *UDSTransport, s conversion.Scope) error {
	return autoConvert_apiserver_UDSTransport_To_v1_UDSTransport(in, out, s)
}
