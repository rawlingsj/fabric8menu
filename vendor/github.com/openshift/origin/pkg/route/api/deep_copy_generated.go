// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_api_Route,
		DeepCopy_api_RouteIngress,
		DeepCopy_api_RouteIngressCondition,
		DeepCopy_api_RouteList,
		DeepCopy_api_RoutePort,
		DeepCopy_api_RouteSpec,
		DeepCopy_api_RouteStatus,
		DeepCopy_api_RouterShard,
		DeepCopy_api_TLSConfig,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_api_Route(in Route, out *Route, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_RouteSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_api_RouteStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_RouteIngress(in RouteIngress, out *RouteIngress, c *conversion.Cloner) error {
	out.Host = in.Host
	out.RouterName = in.RouterName
	if in.Conditions != nil {
		in, out := in.Conditions, &out.Conditions
		*out = make([]RouteIngressCondition, len(in))
		for i := range in {
			if err := DeepCopy_api_RouteIngressCondition(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func DeepCopy_api_RouteIngressCondition(in RouteIngressCondition, out *RouteIngressCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	out.Reason = in.Reason
	out.Message = in.Message
	if in.LastTransitionTime != nil {
		in, out := in.LastTransitionTime, &out.LastTransitionTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.LastTransitionTime = nil
	}
	return nil
}

func DeepCopy_api_RouteList(in RouteList, out *RouteList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Route, len(in))
		for i := range in {
			if err := DeepCopy_api_Route(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_api_RoutePort(in RoutePort, out *RoutePort, c *conversion.Cloner) error {
	if err := intstr.DeepCopy_intstr_IntOrString(in.TargetPort, &out.TargetPort, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_RouteSpec(in RouteSpec, out *RouteSpec, c *conversion.Cloner) error {
	out.Host = in.Host
	out.Path = in.Path
	if err := api.DeepCopy_api_ObjectReference(in.To, &out.To, c); err != nil {
		return err
	}
	if in.Port != nil {
		in, out := in.Port, &out.Port
		*out = new(RoutePort)
		if err := DeepCopy_api_RoutePort(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Port = nil
	}
	if in.TLS != nil {
		in, out := in.TLS, &out.TLS
		*out = new(TLSConfig)
		if err := DeepCopy_api_TLSConfig(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.TLS = nil
	}
	return nil
}

func DeepCopy_api_RouteStatus(in RouteStatus, out *RouteStatus, c *conversion.Cloner) error {
	if in.Ingress != nil {
		in, out := in.Ingress, &out.Ingress
		*out = make([]RouteIngress, len(in))
		for i := range in {
			if err := DeepCopy_api_RouteIngress(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func DeepCopy_api_RouterShard(in RouterShard, out *RouterShard, c *conversion.Cloner) error {
	out.ShardName = in.ShardName
	out.DNSSuffix = in.DNSSuffix
	return nil
}

func DeepCopy_api_TLSConfig(in TLSConfig, out *TLSConfig, c *conversion.Cloner) error {
	out.Termination = in.Termination
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = in.InsecureEdgeTerminationPolicy
	return nil
}
